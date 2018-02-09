// Code generated by protoc-gen-go.
// source: IM.Control.proto
// DO NOT EDIT!

/*
Package IM_Control is a generated protocol buffer package.

It is generated from these files:
	IM.Control.proto

It has these top-level messages:
	IMUserCheckReq
	IMUserCheckRsp
*/
package IM_Control

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// service id is 0x0009;
type IMUserCheckReq struct {
	// cmd          0x0901
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionId        *uint32 `protobuf:"varint,2,req,name=session_id" json:"session_id,omitempty"`
	Type             *uint32 `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMUserCheckReq) Reset()                    { *m = IMUserCheckReq{} }
func (m *IMUserCheckReq) String() string            { return proto.CompactTextString(m) }
func (*IMUserCheckReq) ProtoMessage()               {}
func (*IMUserCheckReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IMUserCheckReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMUserCheckReq) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMUserCheckReq) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *IMUserCheckReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMUserCheckRsp struct {
	// cmd          0x0902
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	ResultCode       *uint32 `protobuf:"varint,2,req,name=result_code" json:"result_code,omitempty"`
	ResultMsg        []byte  `protobuf:"bytes,3,opt,name=result_msg" json:"result_msg,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMUserCheckRsp) Reset()                    { *m = IMUserCheckRsp{} }
func (m *IMUserCheckRsp) String() string            { return proto.CompactTextString(m) }
func (*IMUserCheckRsp) ProtoMessage()               {}
func (*IMUserCheckRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IMUserCheckRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMUserCheckRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMUserCheckRsp) GetResultMsg() []byte {
	if m != nil {
		return m.ResultMsg
	}
	return nil
}

func (m *IMUserCheckRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

func init() {
	proto.RegisterType((*IMUserCheckReq)(nil), "IM.Control.IMUserCheckReq")
	proto.RegisterType((*IMUserCheckRsp)(nil), "IM.Control.IMUserCheckRsp")
}

func init() { proto.RegisterFile("IM.Control.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0xcf, 0x41, 0x4b, 0x86, 0x40,
	0x10, 0xc6, 0x71, 0x7c, 0x0d, 0x82, 0x51, 0x2b, 0xd6, 0x20, 0xa1, 0x8b, 0x78, 0xf2, 0xb4, 0x1f,
	0x22, 0x2f, 0x79, 0xf0, 0x12, 0x04, 0xdd, 0x96, 0x6d, 0x9d, 0xd4, 0x52, 0xc7, 0x76, 0x66, 0x0f,
	0x7d, 0xfb, 0xc8, 0x84, 0x88, 0xba, 0xfe, 0x0e, 0xff, 0x87, 0x07, 0xae, 0xda, 0x4e, 0x37, 0xb4,
	0x8a, 0xa7, 0x59, 0x6f, 0x9e, 0x84, 0x14, 0xfc, 0x48, 0xf5, 0x04, 0x17, 0x6d, 0xf7, 0xc8, 0xe8,
	0x9b, 0x11, 0xdd, 0xdb, 0x03, 0xbe, 0xab, 0x4b, 0x38, 0x0f, 0x8c, 0xde, 0x4c, 0x7d, 0x11, 0x95,
	0xa7, 0x3a, 0x53, 0x0a, 0x80, 0x91, 0x79, 0xa2, 0xf5, 0xcb, 0x4e, 0xbb, 0xa5, 0x70, 0x26, 0x1f,
	0x1b, 0x16, 0x71, 0x19, 0xd5, 0x99, 0xca, 0x21, 0xb1, 0x22, 0xd6, 0x8d, 0xa6, 0xb7, 0x62, 0x8b,
	0xeb, 0x32, 0xaa, 0xd3, 0xca, 0xfc, 0x2e, 0xf3, 0xf6, 0xb7, 0x9c, 0x43, 0xe2, 0x91, 0xc3, 0x2c,
	0xc6, 0x51, 0x8f, 0x47, 0x5a, 0x01, 0x1c, 0xb8, 0xf0, 0xb0, 0x0f, 0xa4, 0xff, 0x0e, 0xdc, 0xdd,
	0xc2, 0x8d, 0xa3, 0x45, 0x2f, 0x34, 0x84, 0xd7, 0x09, 0xb5, 0xc8, 0xf7, 0xbd, 0xe7, 0xf0, 0x72,
	0x1f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xcd, 0xbc, 0x50, 0xf6, 0x00, 0x00, 0x00,
}