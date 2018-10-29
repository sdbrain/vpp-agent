// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: l3.proto

package l3

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StaticRoute_RouteType int32

const (
	StaticRoute_INTRA_VRF StaticRoute_RouteType = 0
	StaticRoute_INTER_VRF StaticRoute_RouteType = 1
	StaticRoute_DROP      StaticRoute_RouteType = 2
)

var StaticRoute_RouteType_name = map[int32]string{
	0: "INTRA_VRF",
	1: "INTER_VRF",
	2: "DROP",
}
var StaticRoute_RouteType_value = map[string]int32{
	"INTRA_VRF": 0,
	"INTER_VRF": 1,
	"DROP":      2,
}

func (x StaticRoute_RouteType) String() string {
	return proto.EnumName(StaticRoute_RouteType_name, int32(x))
}
func (StaticRoute_RouteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_l3_aa599b35d1453312, []int{0, 0}
}

type IPScanNeighbor_Mode int32

const (
	IPScanNeighbor_DISABLED IPScanNeighbor_Mode = 0
	IPScanNeighbor_IPv4     IPScanNeighbor_Mode = 1
	IPScanNeighbor_IPv6     IPScanNeighbor_Mode = 2
	IPScanNeighbor_BOTH     IPScanNeighbor_Mode = 3
)

var IPScanNeighbor_Mode_name = map[int32]string{
	0: "DISABLED",
	1: "IPv4",
	2: "IPv6",
	3: "BOTH",
}
var IPScanNeighbor_Mode_value = map[string]int32{
	"DISABLED": 0,
	"IPv4":     1,
	"IPv6":     2,
	"BOTH":     3,
}

func (x IPScanNeighbor_Mode) String() string {
	return proto.EnumName(IPScanNeighbor_Mode_name, int32(x))
}
func (IPScanNeighbor_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_l3_aa599b35d1453312, []int{4, 0}
}

// Static Routes
type StaticRoute struct {
	Type              StaticRoute_RouteType `protobuf:"varint,10,opt,name=type,proto3,enum=l3.StaticRoute_RouteType" json:"type,omitempty"`
	VrfId             uint32                `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	DstNetwork        string                `protobuf:"bytes,3,opt,name=dst_network,json=dstNetwork,proto3" json:"dst_network,omitempty"`
	NextHopAddr       string                `protobuf:"bytes,4,opt,name=next_hop_addr,json=nextHopAddr,proto3" json:"next_hop_addr,omitempty"`
	OutgoingInterface string                `protobuf:"bytes,5,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	Weight            uint32                `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	Preference        uint32                `protobuf:"varint,7,opt,name=preference,proto3" json:"preference,omitempty"`
	// (a poor man's primary and backup)
	ViaVrfId             uint32   `protobuf:"varint,8,opt,name=via_vrf_id,json=viaVrfId,proto3" json:"via_vrf_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticRoute) Reset()         { *m = StaticRoute{} }
func (m *StaticRoute) String() string { return proto.CompactTextString(m) }
func (*StaticRoute) ProtoMessage()    {}
func (*StaticRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_aa599b35d1453312, []int{0}
}
func (m *StaticRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticRoute.Unmarshal(m, b)
}
func (m *StaticRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticRoute.Marshal(b, m, deterministic)
}
func (dst *StaticRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticRoute.Merge(dst, src)
}
func (m *StaticRoute) XXX_Size() int {
	return xxx_messageInfo_StaticRoute.Size(m)
}
func (m *StaticRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticRoute.DiscardUnknown(m)
}

var xxx_messageInfo_StaticRoute proto.InternalMessageInfo

func (m *StaticRoute) GetType() StaticRoute_RouteType {
	if m != nil {
		return m.Type
	}
	return StaticRoute_INTRA_VRF
}

func (m *StaticRoute) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *StaticRoute) GetDstNetwork() string {
	if m != nil {
		return m.DstNetwork
	}
	return ""
}

func (m *StaticRoute) GetNextHopAddr() string {
	if m != nil {
		return m.NextHopAddr
	}
	return ""
}

func (m *StaticRoute) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *StaticRoute) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *StaticRoute) GetPreference() uint32 {
	if m != nil {
		return m.Preference
	}
	return 0
}

func (m *StaticRoute) GetViaVrfId() uint32 {
	if m != nil {
		return m.ViaVrfId
	}
	return 0
}

// IP ARP Entries
type ARPEntry struct {
	Interface            string   `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	IpAddress            string   `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	PhysAddress          string   `protobuf:"bytes,3,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	Static               bool     `protobuf:"varint,4,opt,name=static,proto3" json:"static,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ARPEntry) Reset()         { *m = ARPEntry{} }
func (m *ARPEntry) String() string { return proto.CompactTextString(m) }
func (*ARPEntry) ProtoMessage()    {}
func (*ARPEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_aa599b35d1453312, []int{1}
}
func (m *ARPEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ARPEntry.Unmarshal(m, b)
}
func (m *ARPEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ARPEntry.Marshal(b, m, deterministic)
}
func (dst *ARPEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ARPEntry.Merge(dst, src)
}
func (m *ARPEntry) XXX_Size() int {
	return xxx_messageInfo_ARPEntry.Size(m)
}
func (m *ARPEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ARPEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ARPEntry proto.InternalMessageInfo

func (m *ARPEntry) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *ARPEntry) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *ARPEntry) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *ARPEntry) GetStatic() bool {
	if m != nil {
		return m.Static
	}
	return false
}

type ProxyARP struct {
	Interfaces           []*ProxyARP_Interface `protobuf:"bytes,1,rep,name=interfaces" json:"interfaces,omitempty"`
	Ranges               []*ProxyARP_Range     `protobuf:"bytes,2,rep,name=ranges" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ProxyARP) Reset()         { *m = ProxyARP{} }
func (m *ProxyARP) String() string { return proto.CompactTextString(m) }
func (*ProxyARP) ProtoMessage()    {}
func (*ProxyARP) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_aa599b35d1453312, []int{2}
}
func (m *ProxyARP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyARP.Unmarshal(m, b)
}
func (m *ProxyARP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyARP.Marshal(b, m, deterministic)
}
func (dst *ProxyARP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyARP.Merge(dst, src)
}
func (m *ProxyARP) XXX_Size() int {
	return xxx_messageInfo_ProxyARP.Size(m)
}
func (m *ProxyARP) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyARP.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyARP proto.InternalMessageInfo

func (m *ProxyARP) GetInterfaces() []*ProxyARP_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *ProxyARP) GetRanges() []*ProxyARP_Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type ProxyARP_Interface struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyARP_Interface) Reset()         { *m = ProxyARP_Interface{} }
func (m *ProxyARP_Interface) String() string { return proto.CompactTextString(m) }
func (*ProxyARP_Interface) ProtoMessage()    {}
func (*ProxyARP_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_aa599b35d1453312, []int{2, 0}
}
func (m *ProxyARP_Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyARP_Interface.Unmarshal(m, b)
}
func (m *ProxyARP_Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyARP_Interface.Marshal(b, m, deterministic)
}
func (dst *ProxyARP_Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyARP_Interface.Merge(dst, src)
}
func (m *ProxyARP_Interface) XXX_Size() int {
	return xxx_messageInfo_ProxyARP_Interface.Size(m)
}
func (m *ProxyARP_Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyARP_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyARP_Interface proto.InternalMessageInfo

func (m *ProxyARP_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ProxyARP_Range struct {
	FirstIpAddr          string   `protobuf:"bytes,1,opt,name=first_ip_addr,json=firstIpAddr,proto3" json:"first_ip_addr,omitempty"`
	LastIpAddr           string   `protobuf:"bytes,2,opt,name=last_ip_addr,json=lastIpAddr,proto3" json:"last_ip_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyARP_Range) Reset()         { *m = ProxyARP_Range{} }
func (m *ProxyARP_Range) String() string { return proto.CompactTextString(m) }
func (*ProxyARP_Range) ProtoMessage()    {}
func (*ProxyARP_Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_aa599b35d1453312, []int{2, 1}
}
func (m *ProxyARP_Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyARP_Range.Unmarshal(m, b)
}
func (m *ProxyARP_Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyARP_Range.Marshal(b, m, deterministic)
}
func (dst *ProxyARP_Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyARP_Range.Merge(dst, src)
}
func (m *ProxyARP_Range) XXX_Size() int {
	return xxx_messageInfo_ProxyARP_Range.Size(m)
}
func (m *ProxyARP_Range) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyARP_Range.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyARP_Range proto.InternalMessageInfo

func (m *ProxyARP_Range) GetFirstIpAddr() string {
	if m != nil {
		return m.FirstIpAddr
	}
	return ""
}

func (m *ProxyARP_Range) GetLastIpAddr() string {
	if m != nil {
		return m.LastIpAddr
	}
	return ""
}

// STN (Steal The NIC) feature table
type STNEntry struct {
	IpAddress            string   `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Interface            string   `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *STNEntry) Reset()         { *m = STNEntry{} }
func (m *STNEntry) String() string { return proto.CompactTextString(m) }
func (*STNEntry) ProtoMessage()    {}
func (*STNEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_aa599b35d1453312, []int{3}
}
func (m *STNEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_STNEntry.Unmarshal(m, b)
}
func (m *STNEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_STNEntry.Marshal(b, m, deterministic)
}
func (dst *STNEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_STNEntry.Merge(dst, src)
}
func (m *STNEntry) XXX_Size() int {
	return xxx_messageInfo_STNEntry.Size(m)
}
func (m *STNEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_STNEntry.DiscardUnknown(m)
}

var xxx_messageInfo_STNEntry proto.InternalMessageInfo

func (m *STNEntry) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *STNEntry) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

// Enables/disables IP neighbor scanning
type IPScanNeighbor struct {
	Mode                 IPScanNeighbor_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=l3.IPScanNeighbor_Mode" json:"mode,omitempty"`
	ScanInterval         uint32              `protobuf:"varint,2,opt,name=scan_interval,json=scanInterval,proto3" json:"scan_interval,omitempty"`
	MaxProcTime          uint32              `protobuf:"varint,3,opt,name=max_proc_time,json=maxProcTime,proto3" json:"max_proc_time,omitempty"`
	MaxUpdate            uint32              `protobuf:"varint,4,opt,name=max_update,json=maxUpdate,proto3" json:"max_update,omitempty"`
	ScanIntDelay         uint32              `protobuf:"varint,5,opt,name=scan_int_delay,json=scanIntDelay,proto3" json:"scan_int_delay,omitempty"`
	StaleThreshold       uint32              `protobuf:"varint,6,opt,name=stale_threshold,json=staleThreshold,proto3" json:"stale_threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IPScanNeighbor) Reset()         { *m = IPScanNeighbor{} }
func (m *IPScanNeighbor) String() string { return proto.CompactTextString(m) }
func (*IPScanNeighbor) ProtoMessage()    {}
func (*IPScanNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_aa599b35d1453312, []int{4}
}
func (m *IPScanNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPScanNeighbor.Unmarshal(m, b)
}
func (m *IPScanNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPScanNeighbor.Marshal(b, m, deterministic)
}
func (dst *IPScanNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPScanNeighbor.Merge(dst, src)
}
func (m *IPScanNeighbor) XXX_Size() int {
	return xxx_messageInfo_IPScanNeighbor.Size(m)
}
func (m *IPScanNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_IPScanNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_IPScanNeighbor proto.InternalMessageInfo

func (m *IPScanNeighbor) GetMode() IPScanNeighbor_Mode {
	if m != nil {
		return m.Mode
	}
	return IPScanNeighbor_DISABLED
}

func (m *IPScanNeighbor) GetScanInterval() uint32 {
	if m != nil {
		return m.ScanInterval
	}
	return 0
}

func (m *IPScanNeighbor) GetMaxProcTime() uint32 {
	if m != nil {
		return m.MaxProcTime
	}
	return 0
}

func (m *IPScanNeighbor) GetMaxUpdate() uint32 {
	if m != nil {
		return m.MaxUpdate
	}
	return 0
}

func (m *IPScanNeighbor) GetScanIntDelay() uint32 {
	if m != nil {
		return m.ScanIntDelay
	}
	return 0
}

func (m *IPScanNeighbor) GetStaleThreshold() uint32 {
	if m != nil {
		return m.StaleThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*StaticRoute)(nil), "l3.StaticRoute")
	proto.RegisterType((*ARPEntry)(nil), "l3.ARPEntry")
	proto.RegisterType((*ProxyARP)(nil), "l3.ProxyARP")
	proto.RegisterType((*ProxyARP_Interface)(nil), "l3.ProxyARP.Interface")
	proto.RegisterType((*ProxyARP_Range)(nil), "l3.ProxyARP.Range")
	proto.RegisterType((*STNEntry)(nil), "l3.STNEntry")
	proto.RegisterType((*IPScanNeighbor)(nil), "l3.IPScanNeighbor")
	proto.RegisterEnum("l3.StaticRoute_RouteType", StaticRoute_RouteType_name, StaticRoute_RouteType_value)
	proto.RegisterEnum("l3.IPScanNeighbor_Mode", IPScanNeighbor_Mode_name, IPScanNeighbor_Mode_value)
}

func init() { proto.RegisterFile("l3.proto", fileDescriptor_l3_aa599b35d1453312) }

var fileDescriptor_l3_aa599b35d1453312 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x54, 0x5d, 0x6f, 0xda, 0x48,
	0x14, 0x8d, 0x1d, 0xc2, 0x9a, 0x6b, 0xcc, 0xb2, 0x23, 0x6d, 0xd6, 0x1b, 0xa5, 0x0d, 0x75, 0x2b,
	0x15, 0xb5, 0x0a, 0x0f, 0x50, 0xe5, 0x9d, 0x88, 0xb4, 0xb1, 0xd4, 0x10, 0x6b, 0xa0, 0x79, 0xb5,
	0x26, 0xf6, 0x00, 0x56, 0x6d, 0x8f, 0x35, 0x9e, 0x38, 0xf0, 0xda, 0x5f, 0xd3, 0xdf, 0xd3, 0x1f,
	0xd0, 0xdf, 0x52, 0xcd, 0xd8, 0xe6, 0x23, 0x2f, 0xc8, 0x73, 0xee, 0x19, 0xdd, 0x7b, 0xce, 0xb9,
	0x03, 0x18, 0xf1, 0x68, 0x90, 0x71, 0x26, 0x18, 0xd2, 0xe3, 0x91, 0xf3, 0x4b, 0x07, 0x73, 0x26,
	0x88, 0x88, 0x02, 0xcc, 0x9e, 0x04, 0x45, 0x97, 0xd0, 0x10, 0x9b, 0x8c, 0xda, 0xd0, 0xd3, 0xfa,
	0x9d, 0xe1, 0xff, 0x83, 0x78, 0x34, 0xd8, 0x2b, 0x0f, 0xd4, 0xef, 0x7c, 0x93, 0x51, 0xac, 0x68,
	0xe8, 0x5f, 0x68, 0x16, 0x7c, 0xe1, 0x47, 0xa1, 0xad, 0xf5, 0xb4, 0xbe, 0x85, 0x4f, 0x0a, 0xbe,
	0x70, 0x43, 0x74, 0x01, 0x66, 0x98, 0x0b, 0x3f, 0xa5, 0xe2, 0x99, 0xf1, 0xef, 0xf6, 0x71, 0x4f,
	0xeb, 0xb7, 0x30, 0x84, 0xb9, 0x98, 0x96, 0x08, 0x72, 0xc0, 0x4a, 0xe9, 0x5a, 0xf8, 0x2b, 0x96,
	0xf9, 0x24, 0x0c, 0xb9, 0xdd, 0x50, 0x14, 0x53, 0x82, 0xb7, 0x2c, 0x1b, 0x87, 0x21, 0x47, 0x97,
	0x80, 0xd8, 0x93, 0x58, 0xb2, 0x28, 0x5d, 0xfa, 0x51, 0x2a, 0x28, 0x5f, 0x90, 0x80, 0xda, 0x27,
	0x8a, 0xf8, 0x4f, 0x5d, 0x71, 0xeb, 0x02, 0x3a, 0x85, 0xe6, 0x33, 0x8d, 0x96, 0x2b, 0x61, 0x37,
	0xd5, 0x28, 0xd5, 0x09, 0xbd, 0x06, 0xc8, 0x38, 0x5d, 0x50, 0x4e, 0xd3, 0x80, 0xda, 0x7f, 0xa9,
	0xda, 0x1e, 0x82, 0xce, 0x01, 0x8a, 0x88, 0xf8, 0x95, 0x0c, 0x43, 0xd5, 0x8d, 0x22, 0x22, 0x0f,
	0x52, 0x89, 0x33, 0x82, 0xd6, 0x56, 0x33, 0xb2, 0xa0, 0xe5, 0x4e, 0xe7, 0x78, 0xec, 0x3f, 0xe0,
	0xcf, 0xdd, 0xa3, 0xea, 0x78, 0x83, 0xd5, 0x51, 0x43, 0x06, 0x34, 0x26, 0xf8, 0xde, 0xeb, 0xea,
	0xce, 0x0f, 0x0d, 0x8c, 0x31, 0xf6, 0x6e, 0x52, 0xc1, 0x37, 0xe8, 0x1c, 0x5a, 0xbb, 0xe9, 0x35,
	0x35, 0xfd, 0x0e, 0x40, 0xaf, 0x00, 0xa2, 0xd2, 0x02, 0x9a, 0xe7, 0xb6, 0x5e, 0x95, 0x95, 0x01,
	0x34, 0xcf, 0xd1, 0x1b, 0x68, 0x67, 0xab, 0x4d, 0xbe, 0x25, 0x94, 0x4e, 0x9a, 0x12, 0xab, 0x29,
	0xa7, 0xd0, 0xcc, 0x55, 0x42, 0xca, 0x43, 0x03, 0x57, 0x27, 0xe7, 0xb7, 0x06, 0x86, 0xc7, 0xd9,
	0x7a, 0x33, 0xc6, 0x1e, 0xba, 0x02, 0xd8, 0xf6, 0xcc, 0x6d, 0xad, 0x77, 0xdc, 0x37, 0x87, 0xa7,
	0x32, 0xdc, 0x9a, 0x31, 0xd8, 0x1a, 0x89, 0xf7, 0x98, 0xe8, 0x03, 0x34, 0x39, 0x49, 0x97, 0x54,
	0x8e, 0x26, 0xef, 0xa0, 0x83, 0x3b, 0x58, 0x96, 0x70, 0xc5, 0x38, 0xbb, 0x80, 0xd6, 0x2e, 0x0d,
	0x04, 0x8d, 0x94, 0x24, 0xb5, 0x60, 0xf5, 0x7d, 0x76, 0x07, 0x27, 0xea, 0x86, 0x4c, 0x7f, 0x11,
	0xf1, 0x5c, 0xf8, 0x95, 0xf4, 0x8a, 0x65, 0x2a, 0xd0, 0x2d, 0xd3, 0xef, 0x41, 0x3b, 0x26, 0x7b,
	0x94, 0xd2, 0x1a, 0x90, 0x58, 0xc9, 0x70, 0xbe, 0x80, 0x31, 0x9b, 0x4f, 0x4b, 0x93, 0x0f, 0x6d,
	0xd4, 0x5e, 0xda, 0x78, 0x90, 0x81, 0xfe, 0x22, 0x03, 0xe7, 0xa7, 0x0e, 0x1d, 0xd7, 0x9b, 0x05,
	0x24, 0x9d, 0xca, 0x95, 0x79, 0x64, 0x1c, 0x7d, 0x84, 0x46, 0xc2, 0xc2, 0x72, 0xfc, 0xce, 0xf0,
	0x3f, 0xa9, 0xfa, 0x90, 0x31, 0xb8, 0x63, 0x21, 0xc5, 0x8a, 0x84, 0xde, 0x82, 0x95, 0x07, 0x24,
	0x2d, 0x97, 0xb4, 0x20, 0xb1, 0xea, 0x60, 0xe1, 0xb6, 0x04, 0xdd, 0x0a, 0x93, 0x9a, 0x13, 0xb2,
	0xf6, 0x33, 0xce, 0x02, 0x5f, 0x44, 0x09, 0x55, 0x51, 0x5a, 0xd8, 0x4c, 0xc8, 0xda, 0xe3, 0x2c,
	0x98, 0x47, 0x89, 0x5a, 0x06, 0xc9, 0x79, 0xca, 0x42, 0x22, 0xa8, 0x8a, 0xd3, 0xc2, 0xad, 0x84,
	0xac, 0xbf, 0x29, 0x00, 0xbd, 0x83, 0x4e, 0xdd, 0xc7, 0x0f, 0x69, 0x4c, 0x36, 0xea, 0x31, 0xec,
	0x1a, 0x4d, 0x24, 0x86, 0xde, 0xc3, 0xdf, 0xb9, 0x20, 0x31, 0xf5, 0xc5, 0x8a, 0xd3, 0x7c, 0xc5,
	0xe2, 0xb0, 0x7a, 0x10, 0x1d, 0x05, 0xcf, 0x6b, 0xd4, 0x19, 0x42, 0x43, 0x8a, 0x40, 0x6d, 0x30,
	0x26, 0xee, 0x6c, 0x7c, 0xfd, 0xf5, 0x66, 0xd2, 0x3d, 0x92, 0x5b, 0xec, 0x7a, 0xc5, 0xa7, 0x72,
	0x9f, 0x5d, 0xaf, 0xb8, 0xea, 0xea, 0xf2, 0xeb, 0xfa, 0x7e, 0x7e, 0xdb, 0x3d, 0x7e, 0x6c, 0xaa,
	0x7f, 0x8e, 0xd1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x99, 0xc5, 0xe0, 0x45, 0x04, 0x00,
	0x00,
}
