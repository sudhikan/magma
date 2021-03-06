// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feg/protos/basic_acct.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// user session descriptor
type AcctSession struct {
	// user identity
	// a union to enable future extension
	//
	// Types that are valid to be assigned to User:
	//	*AcctSession_IMSI
	//	*AcctSession_CertificateSerialNumber
	//	*AcctSession_HardwareAddr
	//	*AcctSession_Name
	User isAcctSession_User `protobuf_oneof:"user"`
	// unique session ID
	SessionId string `protobuf:"bytes,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// ID of the service provider gateway (AP, AGW, enodeb, etc.)
	ServingApn string `protobuf:"bytes,6,opt,name=serving_apn,json=servingApn,proto3" json:"serving_apn,omitempty"`
	// available QoS at the caller's site
	AvailableQos         *AcctQoS `protobuf:"bytes,7,opt,name=available_qos,json=availableQos,proto3" json:"available_qos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcctSession) Reset()         { *m = AcctSession{} }
func (m *AcctSession) String() string { return proto.CompactTextString(m) }
func (*AcctSession) ProtoMessage()    {}
func (*AcctSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_1026fc65f92257f8, []int{0}
}

func (m *AcctSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcctSession.Unmarshal(m, b)
}
func (m *AcctSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcctSession.Marshal(b, m, deterministic)
}
func (m *AcctSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcctSession.Merge(m, src)
}
func (m *AcctSession) XXX_Size() int {
	return xxx_messageInfo_AcctSession.Size(m)
}
func (m *AcctSession) XXX_DiscardUnknown() {
	xxx_messageInfo_AcctSession.DiscardUnknown(m)
}

var xxx_messageInfo_AcctSession proto.InternalMessageInfo

type isAcctSession_User interface {
	isAcctSession_User()
}

type AcctSession_IMSI struct {
	IMSI string `protobuf:"bytes,1,opt,name=IMSI,proto3,oneof"`
}

type AcctSession_CertificateSerialNumber struct {
	CertificateSerialNumber []byte `protobuf:"bytes,2,opt,name=certificate_serial_number,json=certificateSerialNumber,proto3,oneof"`
}

type AcctSession_HardwareAddr struct {
	HardwareAddr []byte `protobuf:"bytes,3,opt,name=hardware_addr,json=hardwareAddr,proto3,oneof"`
}

type AcctSession_Name struct {
	Name string `protobuf:"bytes,4,opt,name=name,proto3,oneof"`
}

func (*AcctSession_IMSI) isAcctSession_User() {}

func (*AcctSession_CertificateSerialNumber) isAcctSession_User() {}

func (*AcctSession_HardwareAddr) isAcctSession_User() {}

func (*AcctSession_Name) isAcctSession_User() {}

func (m *AcctSession) GetUser() isAcctSession_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *AcctSession) GetIMSI() string {
	if x, ok := m.GetUser().(*AcctSession_IMSI); ok {
		return x.IMSI
	}
	return ""
}

func (m *AcctSession) GetCertificateSerialNumber() []byte {
	if x, ok := m.GetUser().(*AcctSession_CertificateSerialNumber); ok {
		return x.CertificateSerialNumber
	}
	return nil
}

func (m *AcctSession) GetHardwareAddr() []byte {
	if x, ok := m.GetUser().(*AcctSession_HardwareAddr); ok {
		return x.HardwareAddr
	}
	return nil
}

func (m *AcctSession) GetName() string {
	if x, ok := m.GetUser().(*AcctSession_Name); ok {
		return x.Name
	}
	return ""
}

func (m *AcctSession) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *AcctSession) GetServingApn() string {
	if m != nil {
		return m.ServingApn
	}
	return ""
}

func (m *AcctSession) GetAvailableQos() *AcctQoS {
	if m != nil {
		return m.AvailableQos
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AcctSession) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AcctSession_IMSI)(nil),
		(*AcctSession_CertificateSerialNumber)(nil),
		(*AcctSession_HardwareAddr)(nil),
		(*AcctSession_Name)(nil),
	}
}

// Start session response
type AcctSessionResp struct {
	ReportingAdvisory *AcctSessionResp_ReportLimits `protobuf:"bytes,1,opt,name=reporting_advisory,json=reportingAdvisory,proto3" json:"reporting_advisory,omitempty"`
	// minimal required QoS, the session has to be terminated if service provider's site
	// cannot guarantee the requested QoS. Currently this is an optional parameter which may be ignored by the client
	MinAcceptableQos     *AcctQoS `protobuf:"bytes,2,opt,name=min_acceptable_qos,json=minAcceptableQos,proto3" json:"min_acceptable_qos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcctSessionResp) Reset()         { *m = AcctSessionResp{} }
func (m *AcctSessionResp) String() string { return proto.CompactTextString(m) }
func (*AcctSessionResp) ProtoMessage()    {}
func (*AcctSessionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1026fc65f92257f8, []int{1}
}

func (m *AcctSessionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcctSessionResp.Unmarshal(m, b)
}
func (m *AcctSessionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcctSessionResp.Marshal(b, m, deterministic)
}
func (m *AcctSessionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcctSessionResp.Merge(m, src)
}
func (m *AcctSessionResp) XXX_Size() int {
	return xxx_messageInfo_AcctSessionResp.Size(m)
}
func (m *AcctSessionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AcctSessionResp.DiscardUnknown(m)
}

var xxx_messageInfo_AcctSessionResp proto.InternalMessageInfo

func (m *AcctSessionResp) GetReportingAdvisory() *AcctSessionResp_ReportLimits {
	if m != nil {
		return m.ReportingAdvisory
	}
	return nil
}

func (m *AcctSessionResp) GetMinAcceptableQos() *AcctQoS {
	if m != nil {
		return m.MinAcceptableQos
	}
	return nil
}

// optional update triggers
// user identity provider will use report_limits to express its update
// frequency preferences. Service provider is encouraged, but not required
// to comply with specified reporting preferences
type AcctSessionResp_ReportLimits struct {
	// octets_in - trigger update when RX bytes were consumed by the user from the last update event
	// default is 0, no RX trigger
	OctetsIn uint64 `protobuf:"varint,1,opt,name=octets_in,json=octetsIn,proto3" json:"octets_in,omitempty"`
	// octets_out - trigger update when TX bytes were consumed by the user from the last update event
	// default is 0, no TX trigger
	OctetsOut uint64 `protobuf:"varint,2,opt,name=octets_out,json=octetsOut,proto3" json:"octets_out,omitempty"`
	// elapsed_time_sec - trigger update when elapsed_time_sec seconds passed from the last update event
	ElapsedTimeSec       uint32   `protobuf:"varint,3,opt,name=elapsed_time_sec,json=elapsedTimeSec,proto3" json:"elapsed_time_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcctSessionResp_ReportLimits) Reset()         { *m = AcctSessionResp_ReportLimits{} }
func (m *AcctSessionResp_ReportLimits) String() string { return proto.CompactTextString(m) }
func (*AcctSessionResp_ReportLimits) ProtoMessage()    {}
func (*AcctSessionResp_ReportLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_1026fc65f92257f8, []int{1, 0}
}

func (m *AcctSessionResp_ReportLimits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcctSessionResp_ReportLimits.Unmarshal(m, b)
}
func (m *AcctSessionResp_ReportLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcctSessionResp_ReportLimits.Marshal(b, m, deterministic)
}
func (m *AcctSessionResp_ReportLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcctSessionResp_ReportLimits.Merge(m, src)
}
func (m *AcctSessionResp_ReportLimits) XXX_Size() int {
	return xxx_messageInfo_AcctSessionResp_ReportLimits.Size(m)
}
func (m *AcctSessionResp_ReportLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_AcctSessionResp_ReportLimits.DiscardUnknown(m)
}

var xxx_messageInfo_AcctSessionResp_ReportLimits proto.InternalMessageInfo

func (m *AcctSessionResp_ReportLimits) GetOctetsIn() uint64 {
	if m != nil {
		return m.OctetsIn
	}
	return 0
}

func (m *AcctSessionResp_ReportLimits) GetOctetsOut() uint64 {
	if m != nil {
		return m.OctetsOut
	}
	return 0
}

func (m *AcctSessionResp_ReportLimits) GetElapsedTimeSec() uint32 {
	if m != nil {
		return m.ElapsedTimeSec
	}
	return 0
}

// update_req is relying information on user's ongoing session consumption
type AcctUpdateReq struct {
	// ongoing session information
	Session *AcctSession `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	// octets_in indicates how many octets have been received by the user
	// from the service provider over the course of this session (accumulative)
	// The accumulative nature of this field should compensate for intermittent
	// losses of connectivity
	OctetsIn uint64 `protobuf:"varint,2,opt,name=octets_in,json=octetsIn,proto3" json:"octets_in,omitempty"`
	// octets_out indicates how many octets have been sent on behalf of the user
	// by the service provider over the course of this session (accumulative)
	// The accumulative nature of this field should compensate for intermittent
	// losses of connectivity
	OctetsOut uint64 `protobuf:"varint,3,opt,name=octets_out,json=octetsOut,proto3" json:"octets_out,omitempty"`
	// session_time indicates how many seconds the user/session has received service for
	SessionTime          uint32   `protobuf:"varint,4,opt,name=session_time,json=sessionTime,proto3" json:"session_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcctUpdateReq) Reset()         { *m = AcctUpdateReq{} }
func (m *AcctUpdateReq) String() string { return proto.CompactTextString(m) }
func (*AcctUpdateReq) ProtoMessage()    {}
func (*AcctUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1026fc65f92257f8, []int{2}
}

func (m *AcctUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcctUpdateReq.Unmarshal(m, b)
}
func (m *AcctUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcctUpdateReq.Marshal(b, m, deterministic)
}
func (m *AcctUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcctUpdateReq.Merge(m, src)
}
func (m *AcctUpdateReq) XXX_Size() int {
	return xxx_messageInfo_AcctUpdateReq.Size(m)
}
func (m *AcctUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AcctUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_AcctUpdateReq proto.InternalMessageInfo

func (m *AcctUpdateReq) GetSession() *AcctSession {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *AcctUpdateReq) GetOctetsIn() uint64 {
	if m != nil {
		return m.OctetsIn
	}
	return 0
}

func (m *AcctUpdateReq) GetOctetsOut() uint64 {
	if m != nil {
		return m.OctetsOut
	}
	return 0
}

func (m *AcctUpdateReq) GetSessionTime() uint32 {
	if m != nil {
		return m.SessionTime
	}
	return 0
}

type AcctStopResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcctStopResp) Reset()         { *m = AcctStopResp{} }
func (m *AcctStopResp) String() string { return proto.CompactTextString(m) }
func (*AcctStopResp) ProtoMessage()    {}
func (*AcctStopResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1026fc65f92257f8, []int{3}
}

func (m *AcctStopResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcctStopResp.Unmarshal(m, b)
}
func (m *AcctStopResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcctStopResp.Marshal(b, m, deterministic)
}
func (m *AcctStopResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcctStopResp.Merge(m, src)
}
func (m *AcctStopResp) XXX_Size() int {
	return xxx_messageInfo_AcctStopResp.Size(m)
}
func (m *AcctStopResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AcctStopResp.DiscardUnknown(m)
}

var xxx_messageInfo_AcctStopResp proto.InternalMessageInfo

// Quality Of Service data
type AcctQoS struct {
	DownloadMbps         float32  `protobuf:"fixed32,1,opt,name=download_mbps,json=downloadMbps,proto3" json:"download_mbps,omitempty"`
	UploadMbps           float32  `protobuf:"fixed32,2,opt,name=upload_mbps,json=uploadMbps,proto3" json:"upload_mbps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcctQoS) Reset()         { *m = AcctQoS{} }
func (m *AcctQoS) String() string { return proto.CompactTextString(m) }
func (*AcctQoS) ProtoMessage()    {}
func (*AcctQoS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1026fc65f92257f8, []int{4}
}

func (m *AcctQoS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcctQoS.Unmarshal(m, b)
}
func (m *AcctQoS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcctQoS.Marshal(b, m, deterministic)
}
func (m *AcctQoS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcctQoS.Merge(m, src)
}
func (m *AcctQoS) XXX_Size() int {
	return xxx_messageInfo_AcctQoS.Size(m)
}
func (m *AcctQoS) XXX_DiscardUnknown() {
	xxx_messageInfo_AcctQoS.DiscardUnknown(m)
}

var xxx_messageInfo_AcctQoS proto.InternalMessageInfo

func (m *AcctQoS) GetDownloadMbps() float32 {
	if m != nil {
		return m.DownloadMbps
	}
	return 0
}

func (m *AcctQoS) GetUploadMbps() float32 {
	if m != nil {
		return m.UploadMbps
	}
	return 0
}

func init() {
	proto.RegisterType((*AcctSession)(nil), "magma.feg.AcctSession")
	proto.RegisterType((*AcctSessionResp)(nil), "magma.feg.AcctSessionResp")
	proto.RegisterType((*AcctSessionResp_ReportLimits)(nil), "magma.feg.AcctSessionResp.ReportLimits")
	proto.RegisterType((*AcctUpdateReq)(nil), "magma.feg.AcctUpdateReq")
	proto.RegisterType((*AcctStopResp)(nil), "magma.feg.AcctStopResp")
	proto.RegisterType((*AcctQoS)(nil), "magma.feg.AcctQoS")
}

func init() { proto.RegisterFile("feg/protos/basic_acct.proto", fileDescriptor_1026fc65f92257f8) }

var fileDescriptor_1026fc65f92257f8 = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x6d, 0xb6, 0x69, 0x4a, 0x26, 0x9b, 0x52, 0x2c, 0x44, 0xb7, 0xad, 0x10, 0x25, 0x08, 0x91,
	0x53, 0x82, 0xca, 0x01, 0x21, 0x10, 0x22, 0x3d, 0x35, 0x12, 0xa5, 0xaa, 0x03, 0x1c, 0xb8, 0xac,
	0x1c, 0xdb, 0x0d, 0x96, 0x76, 0x6d, 0xd7, 0xf6, 0xb6, 0xe2, 0x6b, 0x80, 0x7f, 0xe1, 0xc3, 0x90,
	0x1d, 0x67, 0x09, 0x91, 0x92, 0xd3, 0x6a, 0xdf, 0xbc, 0xf1, 0x3c, 0xbf, 0x37, 0x32, 0x1c, 0x5f,
	0xf3, 0xd9, 0x50, 0x1b, 0xe5, 0x94, 0x1d, 0x4e, 0x89, 0x15, 0x34, 0x27, 0x94, 0xba, 0x41, 0x40,
	0x50, 0xbb, 0x24, 0xb3, 0x92, 0x0c, 0xae, 0xf9, 0xac, 0xf7, 0x33, 0x81, 0xce, 0x88, 0x52, 0x37,
	0xe1, 0xd6, 0x0a, 0x25, 0xd1, 0x43, 0x68, 0x8e, 0x2f, 0x26, 0xe3, 0xac, 0x71, 0xd2, 0xe8, 0xb7,
	0xcf, 0xb7, 0x70, 0xf8, 0x43, 0xef, 0xe0, 0x90, 0x72, 0xe3, 0xc4, 0xb5, 0xa0, 0xc4, 0xf1, 0xdc,
	0x72, 0x23, 0x48, 0x91, 0xcb, 0xaa, 0x9c, 0x72, 0x93, 0x25, 0x27, 0x8d, 0x7e, 0x7a, 0xbe, 0x85,
	0x0f, 0x96, 0x28, 0x93, 0xc0, 0xf8, 0x14, 0x08, 0xe8, 0x39, 0x74, 0xbf, 0x13, 0xc3, 0xee, 0x88,
	0xe1, 0x39, 0x61, 0xcc, 0x64, 0xdb, 0xb1, 0x23, 0x5d, 0xc0, 0x23, 0xc6, 0x8c, 0x1f, 0x2d, 0x49,
	0xc9, 0xb3, 0xe6, 0x62, 0xb4, 0xff, 0x43, 0x8f, 0x01, 0xec, 0x5c, 0x5b, 0x2e, 0x58, 0xb6, 0xe3,
	0x6b, 0xb8, 0x1d, 0x91, 0x31, 0x43, 0x4f, 0xa0, 0x63, 0xb9, 0xb9, 0x15, 0x72, 0x96, 0x13, 0x2d,
	0xb3, 0x56, 0xa8, 0x43, 0x84, 0x46, 0x5a, 0xa2, 0xd7, 0xd0, 0x25, 0xb7, 0x44, 0x14, 0x64, 0x5a,
	0xf0, 0xfc, 0x46, 0xd9, 0x6c, 0xf7, 0xa4, 0xd1, 0xef, 0x9c, 0xa2, 0x41, 0xed, 0xc1, 0xc0, 0xdf,
	0xff, 0x4a, 0x4d, 0x70, 0x5a, 0x13, 0xaf, 0x94, 0x3d, 0x6b, 0x41, 0xb3, 0xb2, 0xdc, 0xf4, 0x7e,
	0x27, 0x70, 0x7f, 0xc9, 0x21, 0xcc, 0xad, 0x46, 0x5f, 0x01, 0x19, 0xae, 0x95, 0x71, 0x61, 0x2e,
	0xbb, 0x15, 0x56, 0x99, 0x1f, 0xc1, 0xb3, 0xce, 0xe9, 0x8b, 0x95, 0x93, 0x97, 0xfa, 0x06, 0x38,
	0x34, 0x7d, 0x14, 0xa5, 0x70, 0x16, 0x3f, 0xa8, 0x8f, 0x18, 0xc5, 0x13, 0xd0, 0x07, 0x40, 0xa5,
	0x90, 0x3e, 0x2a, 0xae, 0x5d, 0xad, 0x38, 0x59, 0xab, 0x78, 0xbf, 0x14, 0x72, 0x54, 0x93, 0xaf,
	0x94, 0x3d, 0x72, 0x90, 0x2e, 0x0f, 0x41, 0xc7, 0xd0, 0x56, 0xd4, 0x71, 0x67, 0x73, 0x21, 0x83,
	0xc0, 0x26, 0xbe, 0x37, 0x07, 0xc6, 0xd2, 0x7b, 0x1b, 0x8b, 0xaa, 0x72, 0x61, 0x4c, 0x13, 0x47,
	0xfa, 0x65, 0xe5, 0x50, 0x1f, 0xf6, 0x79, 0x41, 0xb4, 0xe5, 0x2c, 0x77, 0xa2, 0xf4, 0xb1, 0xd3,
	0x10, 0x5d, 0x17, 0xef, 0x45, 0xfc, 0xb3, 0x28, 0xf9, 0x84, 0xd3, 0xde, 0xaf, 0x06, 0x74, 0xbd,
	0xa6, 0x2f, 0x9a, 0x11, 0xc7, 0x31, 0xbf, 0x41, 0x2f, 0x61, 0x37, 0x86, 0x14, 0x6d, 0x79, 0xb4,
	0xc6, 0x96, 0x05, 0xed, 0x7f, 0xa5, 0xc9, 0x46, 0xa5, 0xdb, 0xab, 0x4a, 0x9f, 0x42, 0xba, 0x58,
	0x12, 0xaf, 0x34, 0xac, 0x50, 0x17, 0x77, 0x22, 0xe6, 0x55, 0xf6, 0xf6, 0x20, 0x0d, 0x63, 0x9d,
	0xd2, 0x3e, 0x8a, 0xde, 0x25, 0xec, 0x46, 0x17, 0xd1, 0x33, 0xe8, 0x32, 0x75, 0x27, 0x0b, 0x45,
	0x58, 0x5e, 0x4e, 0xb5, 0x0d, 0x8a, 0x13, 0x9c, 0x2e, 0xc0, 0x8b, 0xa9, 0xb6, 0x7e, 0xd1, 0x2a,
	0xfd, 0x8f, 0x92, 0x04, 0x0a, 0xcc, 0x21, 0x4f, 0x38, 0xfd, 0xd3, 0x00, 0x18, 0x51, 0xaa, 0x2a,
	0xe9, 0x23, 0x45, 0x6f, 0x61, 0x67, 0xe2, 0x88, 0x71, 0x68, 0xcd, 0xc5, 0x8f, 0x8e, 0xd6, 0xef,
	0x09, 0x7a, 0x0f, 0xad, 0xb9, 0x95, 0x28, 0x5b, 0x61, 0xd5, 0x0e, 0x6f, 0xec, 0x7f, 0x03, 0x4d,
	0x7f, 0xd1, 0x0d, 0xdd, 0x07, 0xab, 0xdd, 0xd1, 0x97, 0xb3, 0xe3, 0x6f, 0x87, 0xa1, 0x32, 0xf4,
	0x0f, 0x08, 0x2d, 0x54, 0xc5, 0x86, 0x33, 0x15, 0x5f, 0x92, 0x69, 0x2b, 0x7c, 0x5f, 0xfd, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x39, 0x39, 0x9c, 0x3a, 0x5e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccountingClient is the client API for Accounting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountingClient interface {
	// Start will be called at the end of every new user session creation
	// start is responsible for verification & initiation of an accounting contract
	// between the user identity provider/MNO and service provider (ISP/WISP/PLTE)
	// A non-error return will indicate successful contract establishment and will
	// result in the beginning of service for the user
	Start(ctx context.Context, in *AcctSession, opts ...grpc.CallOption) (*AcctSessionResp, error)
	// Update should be continuously called for every ongoing service session to update
	// the user bandwidth usage as well as current quality of provided service.
	// If update returns error the session should be terminated and the user disconnected,
	// In the case of unsuccessful update completion, service provider is suppose to follow up
	// with final Stop call
	Update(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctSessionResp, error)
	// Stop is a notification call to communicate to identity provider
	// user/network  initiated service termination.
	// stop will provide final used bandwidth count. stop call is issued
	// after the user session was terminated.
	Stop(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctStopResp, error)
}

type accountingClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountingClient(cc grpc.ClientConnInterface) AccountingClient {
	return &accountingClient{cc}
}

func (c *accountingClient) Start(ctx context.Context, in *AcctSession, opts ...grpc.CallOption) (*AcctSessionResp, error) {
	out := new(AcctSessionResp)
	err := c.cc.Invoke(ctx, "/magma.feg.Accounting/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) Update(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctSessionResp, error) {
	out := new(AcctSessionResp)
	err := c.cc.Invoke(ctx, "/magma.feg.Accounting/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) Stop(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctStopResp, error) {
	out := new(AcctStopResp)
	err := c.cc.Invoke(ctx, "/magma.feg.Accounting/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountingServer is the server API for Accounting service.
type AccountingServer interface {
	// Start will be called at the end of every new user session creation
	// start is responsible for verification & initiation of an accounting contract
	// between the user identity provider/MNO and service provider (ISP/WISP/PLTE)
	// A non-error return will indicate successful contract establishment and will
	// result in the beginning of service for the user
	Start(context.Context, *AcctSession) (*AcctSessionResp, error)
	// Update should be continuously called for every ongoing service session to update
	// the user bandwidth usage as well as current quality of provided service.
	// If update returns error the session should be terminated and the user disconnected,
	// In the case of unsuccessful update completion, service provider is suppose to follow up
	// with final Stop call
	Update(context.Context, *AcctUpdateReq) (*AcctSessionResp, error)
	// Stop is a notification call to communicate to identity provider
	// user/network  initiated service termination.
	// stop will provide final used bandwidth count. stop call is issued
	// after the user session was terminated.
	Stop(context.Context, *AcctUpdateReq) (*AcctStopResp, error)
}

// UnimplementedAccountingServer can be embedded to have forward compatible implementations.
type UnimplementedAccountingServer struct {
}

func (*UnimplementedAccountingServer) Start(ctx context.Context, req *AcctSession) (*AcctSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedAccountingServer) Update(ctx context.Context, req *AcctUpdateReq) (*AcctSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAccountingServer) Stop(ctx context.Context, req *AcctUpdateReq) (*AcctStopResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterAccountingServer(s *grpc.Server, srv AccountingServer) {
	s.RegisterService(&_Accounting_serviceDesc, srv)
}

func _Accounting_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcctSession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.Accounting/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Start(ctx, req.(*AcctSession))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcctUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.Accounting/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Update(ctx, req.(*AcctUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcctUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.Accounting/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Stop(ctx, req.(*AcctUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.Accounting",
	HandlerType: (*AccountingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Accounting_Start_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Accounting_Update_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Accounting_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/basic_acct.proto",
}
