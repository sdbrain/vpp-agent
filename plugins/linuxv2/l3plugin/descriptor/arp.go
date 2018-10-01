// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package descriptor

import (
	"net"
	"strings"

	"github.com/go-errors/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/vishvananda/netlink"

	"github.com/ligato/cn-infra/logging"
	scheduler "github.com/ligato/vpp-agent/plugins/kvscheduler/api"

	"github.com/ligato/vpp-agent/plugins/linuxv2/ifplugin"
	ifdescriptor "github.com/ligato/vpp-agent/plugins/linuxv2/ifplugin/descriptor"
	"github.com/ligato/vpp-agent/plugins/linuxv2/l3plugin/descriptor/adapter"
	l3linuxcalls "github.com/ligato/vpp-agent/plugins/linuxv2/l3plugin/linuxcalls"
	ifmodel "github.com/ligato/vpp-agent/plugins/linuxv2/model/interfaces"
	"github.com/ligato/vpp-agent/plugins/linuxv2/model/l3"
	"github.com/ligato/vpp-agent/plugins/linuxv2/nsplugin"
	nslinuxcalls "github.com/ligato/vpp-agent/plugins/linuxv2/nsplugin/linuxcalls"
)

const (
	// ARPDescriptorName is the name of the descriptor for Linux ARP entries.
	ARPDescriptorName = "linux-arp"

	// dependency labels
	arpInterfaceDep = "interface"
)

// A list of non-retriable errors:
var (
	// ErrARPWithoutInterface is returned when Linux ARP configuration is missing
	// interface reference.
	ErrARPWithoutInterface = errors.New("Linux ARP entry defined without interface reference")

	// ErrARPWithoutIP is returned when Linux ARP configuration is missing IP address.
	ErrARPWithoutIP = errors.New("Linux ARP entry defined without IP address")

	// ErrARPWithInvalidIP is returned when Linux ARP configuration contains IP address that cannot be parsed.
	ErrARPWithInvalidIP = errors.New("Linux ARP entry defined with invalid IP address")

	// ErrARPWithoutHwAddr is returned when Linux ARP configuration is missing
	// MAC address.
	ErrARPWithoutHwAddr = errors.New("Linux ARP entry defined without MAC address")

	// ErrARPWithInvalidHwAddr is returned when Linux ARP configuration contains MAC address that cannot be parsed.
	ErrARPWithInvalidHwAddr = errors.New("Linux ARP entry defined with invalid MAC address")
)

// ARPDescriptor teaches KVScheduler how to configure Linux ARP entries.
type ARPDescriptor struct {
	log       logging.Logger
	l3Handler l3linuxcalls.NetlinkAPI
	ifPlugin  ifplugin.API
	nsPlugin  nsplugin.API
	scheduler scheduler.KVScheduler
}

// NewARPDescriptor creates a new instance of the ARP descriptor.
func NewARPDescriptor(
	scheduler scheduler.KVScheduler, ifPlugin ifplugin.API, nsPlugin nsplugin.API,
	l3Handler l3linuxcalls.NetlinkAPI, log logging.PluginLogger) *ARPDescriptor {

	return &ARPDescriptor{
		scheduler: scheduler,
		l3Handler: l3Handler,
		ifPlugin:  ifPlugin,
		nsPlugin:  nsPlugin,
		log:       log.NewLogger("-arp-descriptor"),
	}
}

// GetDescriptor returns descriptor suitable for registration (via adapter) with
// the KVScheduler.
func (arpd *ARPDescriptor) GetDescriptor() *adapter.ARPDescriptor {
	return &adapter.ARPDescriptor{
		Name:               ARPDescriptorName,
		KeySelector:        arpd.IsARPKey,
		ValueTypeName:      proto.MessageName(&l3.LinuxStaticARPEntry{}),
		ValueComparator:    arpd.EquivalentARPs,
		NBKeyPrefix:        l3.StaticArpKeyPrefix,
		Add:                arpd.Add,
		Delete:             arpd.Delete,
		Modify:             arpd.Modify,
		IsRetriableFailure: arpd.IsRetriableFailure,
		Dependencies:       arpd.Dependencies,
		Dump:               arpd.Dump,
		DumpDependencies:   []string{ifdescriptor.InterfaceDescriptorName},
	}
}

// IsARPKey returns <true> if the key identifies a Linux ARP configuration.
func (arpd *ARPDescriptor) IsARPKey(key string) bool {
	return strings.HasPrefix(key, l3.StaticArpKeyPrefix)
}

// EquivalentARPs is case-insensitive comparison function for l3.LinuxStaticARPEntry.
func (arpd *ARPDescriptor) EquivalentARPs(key string, arp1, arp2 *l3.LinuxStaticARPEntry) bool {
	// interfaces compared as usually:
	if arp1.Interface != arp2.Interface {
		return false
	}

	// compare MAC addresses case-insensitively
	if strings.ToLower(arp1.HwAddress) != strings.ToLower(arp2.HwAddress) {
		return false
	}

	// compare IP addresses converted to net.IPNet
	return equalAddrs(arp1.IpAddress, arp2.IpAddress)
}

// IsRetriableFailure returns <false> for errors related to invalid configuration.
func (arpd *ARPDescriptor) IsRetriableFailure(err error) bool {
	nonRetriable := []error{
		ErrARPWithoutInterface,
		ErrARPWithoutIP,
		ErrARPWithInvalidIP,
		ErrARPWithoutHwAddr,
		ErrARPWithInvalidHwAddr,
	}
	for _, nonRetriableErr := range nonRetriable {
		if err == nonRetriableErr {
			return false
		}
	}
	return true
}

// Add creates ARP entry.
func (arpd *ARPDescriptor) Add(key string, arp *l3.LinuxStaticARPEntry) (metadata interface{}, err error) {
	err = arpd.updateARPEntry(arp, "add", arpd.l3Handler.SetARPEntry)
	return nil, err
}

// Delete removes ARP entry.
func (arpd *ARPDescriptor) Delete(key string, arp *l3.LinuxStaticARPEntry, metadata interface{}) error {
	return arpd.updateARPEntry(arp, "delete", arpd.l3Handler.DelARPEntry)
}

// Modify is able to change MAC address of the ARP entry.
func (arpd *ARPDescriptor) Modify(key string, oldARP, newARP *l3.LinuxStaticARPEntry, oldMetadata interface{}) (newMetadata interface{}, err error) {
	err = arpd.updateARPEntry(newARP, "modify", arpd.l3Handler.SetARPEntry)
	return nil, err
}

// updateARPEntry adds, modifies or deletes an ARP entry.
func (arpd *ARPDescriptor) updateARPEntry(arp *l3.LinuxStaticARPEntry, actionName string, actionClb func(arpEntry *netlink.Neigh) error) error {
	var err error

	// validate the configuration first
	if arp.Interface == "" {
		err = ErrARPWithoutInterface
		arpd.log.Error(err)
		return err
	}
	if arp.IpAddress == "" {
		err = ErrARPWithoutIP
		arpd.log.Error(err)
		return err
	}
	if arp.HwAddress == "" {
		err = ErrARPWithoutHwAddr
		arpd.log.Error(err)
		return err
	}

	// Prepare ARP entry object
	neigh := &netlink.Neigh{}

	// Get interface metadata
	ifMeta, found := arpd.ifPlugin.GetInterfaceIndex().LookupByName(arp.Interface)
	if !found || ifMeta == nil {
		err = errors.Errorf("failed to obtain metadata for interface %s", arp.Interface)
		arpd.log.Error(err)
		return err
	}

	// set link index
	neigh.LinkIndex = ifMeta.LinuxIfIndex

	// set IP address
	ipAddr := net.ParseIP(arp.IpAddress)
	if ipAddr == nil {
		err = ErrARPWithInvalidIP
		arpd.log.Error(err)
		return err
	}
	neigh.IP = ipAddr

	// set MAC address
	mac, err := net.ParseMAC(arp.HwAddress)
	if err != nil {
		err = ErrARPWithInvalidHwAddr
		arpd.log.Error(err)
		return err
	}
	neigh.HardwareAddr = mac

	// set ARP entry state (always permanent for static ARPs configured by the agent)
	neigh.State = netlink.NUD_PERMANENT

	// set ip family based on the IP address
	if neigh.IP.To4() != nil {
		neigh.Family = netlink.FAMILY_V4
	} else {
		neigh.Family = netlink.FAMILY_V6
	}

	// move to the namespace of the associated interface
	nsCtx := nslinuxcalls.NewNamespaceMgmtCtx()
	revertNs, err := arpd.nsPlugin.SwitchToNamespace(nsCtx, ifMeta.Namespace)
	if err != nil {
		err = errors.Errorf("failed to switch namespace: %v", err)
		arpd.log.Error(err)
		return err
	}
	defer revertNs()

	// update ARP entry in the interface namespace
	err = actionClb(neigh)
	if err != nil {
		err = errors.Errorf("failed to %s linux ARP entry: %v", actionName, err)
		arpd.log.Error(err)
		return err
	}

	return nil
}

// Dependencies lists dependencies for a Linux ARP entry.
func (arpd *ARPDescriptor) Dependencies(key string, arp *l3.LinuxStaticARPEntry) []scheduler.Dependency {
	// the associated interface must exist and be UP
	if arp.Interface != "" {
		return []scheduler.Dependency{
			{
				Label: arpInterfaceDep,
				Key:   ifmodel.InterfaceStateKey(arp.Interface, true),
			},
		}
	}
	return nil
}

// Dump returns all ARP entries associated with interfaces managed by this agent.
func (arpd *ARPDescriptor) Dump(correlate []adapter.ARPKVWithMetadata) ([]adapter.ARPKVWithMetadata, error) {
	var err error
	var dump []adapter.ARPKVWithMetadata
	nsCtx := nslinuxcalls.NewNamespaceMgmtCtx()
	ifMetaIdx := arpd.ifPlugin.GetInterfaceIndex()

	// dump only ARP entries which are associated with interfaces managed by this agent.
	for _, ifName := range ifMetaIdx.ListAllInterfaces() {
		// get interface metadata
		ifMeta, found := ifMetaIdx.LookupByName(ifName)
		if !found || ifMeta == nil {
			err = errors.Errorf("failed to obtain metadata for interface %s", ifName)
			arpd.log.Error(err)
			return dump, err
		}

		// switch to the namespace of the interface
		revertNs, err := arpd.nsPlugin.SwitchToNamespace(nsCtx, ifMeta.Namespace)
		if err != nil {
			err = errors.Errorf("failed to switch namespace: %v", err)
			arpd.log.Error(err)
			return dump, err
		}

		// get ARPs assigned to this interface
		arps, err := arpd.l3Handler.GetARPEntries(ifMeta.LinuxIfIndex)
		revertNs()
		if err != nil {
			arpd.log.Error(err)
			return dump, err
		}

		// convert each ARP from Netlink representation to the NB representation
		for _, arp := range arps {
			if arp.IP.IsLinkLocalMulticast() {
				// skip link-local multi-cast ARPs until there is a requirement to support them as well
				continue
			}
			ipAddr := arp.IP.String()
			hwAddr := arp.HardwareAddr.String()

			dump = append(dump, adapter.ARPKVWithMetadata{
				Key: l3.StaticArpKey(ifName, ipAddr),
				Value: &l3.LinuxStaticARPEntry{
					Interface: ifName,
					IpAddress: ipAddr,
					HwAddress: hwAddr,
				},
				Origin: scheduler.UnknownOrigin, // let the scheduler to determine the origin
			})
		}
	}
	arpd.log.WithField("dump", dump).Debug("Dumping Linux ARPs")
	return dump, nil
}