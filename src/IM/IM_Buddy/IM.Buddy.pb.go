// Code generated by protoc-gen-go.
// source: IM.Buddy.proto
// DO NOT EDIT!

/*
Package IM_Buddy is a generated protocol buffer package.

It is generated from these files:
	IM.Buddy.proto

It has these top-level messages:
	IMRecentContactSessionReq
	IMRecentContactSessionRsp
	IMUserStatNotify
	IMUsersInfoReq
	IMUsersInfoRsp
	IMRemoveSessionReq
	IMRemoveSessionRsp
	IMAllUserReq
	IMAllUserRsp
	IMUsersStatReq
	IMUsersStatRsp
	IMChangeAvatarReq
	IMChangeAvatarRsp
	IMPCLoginStatusNotify
	IMRemoveSessionNotify
	IMDepartmentReq
	IMDepartmentRsp
	IMAvatarChangedNotify
	IMChangeSignInfoReq
	IMChangeSignInfoRsp
	IMSignInfoChangedNotify
	IMUsersInfoByNameReq
	IMUsersInfoByNameRsp
*/
package IM_Buddy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import IM_BaseDefine "../IM_BaseDefine"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IMRecentContactSessionReq struct {
	// cmd id:		0x0201
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	LatestUpdateTime *uint32 `protobuf:"varint,2,req,name=latest_update_time" json:"latest_update_time,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMRecentContactSessionReq) Reset()                    { *m = IMRecentContactSessionReq{} }
func (m *IMRecentContactSessionReq) String() string            { return proto.CompactTextString(m) }
func (*IMRecentContactSessionReq) ProtoMessage()               {}
func (*IMRecentContactSessionReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IMRecentContactSessionReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMRecentContactSessionReq) GetLatestUpdateTime() uint32 {
	if m != nil && m.LatestUpdateTime != nil {
		return *m.LatestUpdateTime
	}
	return 0
}

func (m *IMRecentContactSessionReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMRecentContactSessionRsp struct {
	// cmd id:		0x0202
	UserId             *uint32                             `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	ContactSessionList []*IM_BaseDefine.ContactSessionInfo `protobuf:"bytes,2,rep,name=contact_session_list" json:"contact_session_list,omitempty"`
	AttachData         []byte                              `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized   []byte                              `json:"-"`
}

func (m *IMRecentContactSessionRsp) Reset()                    { *m = IMRecentContactSessionRsp{} }
func (m *IMRecentContactSessionRsp) String() string            { return proto.CompactTextString(m) }
func (*IMRecentContactSessionRsp) ProtoMessage()               {}
func (*IMRecentContactSessionRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IMRecentContactSessionRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMRecentContactSessionRsp) GetContactSessionList() []*IM_BaseDefine.ContactSessionInfo {
	if m != nil {
		return m.ContactSessionList
	}
	return nil
}

func (m *IMRecentContactSessionRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMUserStatNotify struct {
	// cmd id:		0x0203
	UserStat         *IM_BaseDefine.UserStat `protobuf:"bytes,1,req,name=user_stat" json:"user_stat,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *IMUserStatNotify) Reset()                    { *m = IMUserStatNotify{} }
func (m *IMUserStatNotify) String() string            { return proto.CompactTextString(m) }
func (*IMUserStatNotify) ProtoMessage()               {}
func (*IMUserStatNotify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IMUserStatNotify) GetUserStat() *IM_BaseDefine.UserStat {
	if m != nil {
		return m.UserStat
	}
	return nil
}

type IMUsersInfoReq struct {
	// cmd id:		0x0204
	UserId           *uint32  `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	UserIdList       []uint32 `protobuf:"varint,2,rep,name=user_id_list" json:"user_id_list,omitempty"`
	AttachData       []byte   `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *IMUsersInfoReq) Reset()                    { *m = IMUsersInfoReq{} }
func (m *IMUsersInfoReq) String() string            { return proto.CompactTextString(m) }
func (*IMUsersInfoReq) ProtoMessage()               {}
func (*IMUsersInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IMUsersInfoReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMUsersInfoReq) GetUserIdList() []uint32 {
	if m != nil {
		return m.UserIdList
	}
	return nil
}

func (m *IMUsersInfoReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMUsersInfoRsp struct {
	// cmd id:		0x0205
	UserId           *uint32                   `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	UserInfoList     []*IM_BaseDefine.UserInfo `protobuf:"bytes,2,rep,name=user_info_list" json:"user_info_list,omitempty"`
	AttachData       []byte                    `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *IMUsersInfoRsp) Reset()                    { *m = IMUsersInfoRsp{} }
func (m *IMUsersInfoRsp) String() string            { return proto.CompactTextString(m) }
func (*IMUsersInfoRsp) ProtoMessage()               {}
func (*IMUsersInfoRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IMUsersInfoRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMUsersInfoRsp) GetUserInfoList() []*IM_BaseDefine.UserInfo {
	if m != nil {
		return m.UserInfoList
	}
	return nil
}

func (m *IMUsersInfoRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMRemoveSessionReq struct {
	// cmd id:		0x0206
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,2,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,3,req,name=session_id" json:"session_id,omitempty"`
	AttachData       []byte                     `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMRemoveSessionReq) Reset()                    { *m = IMRemoveSessionReq{} }
func (m *IMRemoveSessionReq) String() string            { return proto.CompactTextString(m) }
func (*IMRemoveSessionReq) ProtoMessage()               {}
func (*IMRemoveSessionReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IMRemoveSessionReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMRemoveSessionReq) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

func (m *IMRemoveSessionReq) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMRemoveSessionReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMRemoveSessionRsp struct {
	// cmd id:		0x0207
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	ResultCode       *uint32                    `protobuf:"varint,2,req,name=result_code" json:"result_code,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,3,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,4,req,name=session_id" json:"session_id,omitempty"`
	AttachData       []byte                     `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMRemoveSessionRsp) Reset()                    { *m = IMRemoveSessionRsp{} }
func (m *IMRemoveSessionRsp) String() string            { return proto.CompactTextString(m) }
func (*IMRemoveSessionRsp) ProtoMessage()               {}
func (*IMRemoveSessionRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IMRemoveSessionRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMRemoveSessionRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMRemoveSessionRsp) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

func (m *IMRemoveSessionRsp) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMRemoveSessionRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMAllUserReq struct {
	// cmd id:		0x0208
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	LatestUpdateTime *uint32 `protobuf:"varint,2,req,name=latest_update_time" json:"latest_update_time,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMAllUserReq) Reset()                    { *m = IMAllUserReq{} }
func (m *IMAllUserReq) String() string            { return proto.CompactTextString(m) }
func (*IMAllUserReq) ProtoMessage()               {}
func (*IMAllUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IMAllUserReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMAllUserReq) GetLatestUpdateTime() uint32 {
	if m != nil && m.LatestUpdateTime != nil {
		return *m.LatestUpdateTime
	}
	return 0
}

func (m *IMAllUserReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMAllUserRsp struct {
	// cmd id:		0x0209
	UserId           *uint32                   `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	LatestUpdateTime *uint32                   `protobuf:"varint,2,req,name=latest_update_time" json:"latest_update_time,omitempty"`
	UserList         []*IM_BaseDefine.UserInfo `protobuf:"bytes,3,rep,name=user_list" json:"user_list,omitempty"`
	AttachData       []byte                    `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *IMAllUserRsp) Reset()                    { *m = IMAllUserRsp{} }
func (m *IMAllUserRsp) String() string            { return proto.CompactTextString(m) }
func (*IMAllUserRsp) ProtoMessage()               {}
func (*IMAllUserRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *IMAllUserRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMAllUserRsp) GetLatestUpdateTime() uint32 {
	if m != nil && m.LatestUpdateTime != nil {
		return *m.LatestUpdateTime
	}
	return 0
}

func (m *IMAllUserRsp) GetUserList() []*IM_BaseDefine.UserInfo {
	if m != nil {
		return m.UserList
	}
	return nil
}

func (m *IMAllUserRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMUsersStatReq struct {
	// cmd id:		0x020a
	UserId           *uint32  `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	UserIdList       []uint32 `protobuf:"varint,2,rep,name=user_id_list" json:"user_id_list,omitempty"`
	AttachData       []byte   `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *IMUsersStatReq) Reset()                    { *m = IMUsersStatReq{} }
func (m *IMUsersStatReq) String() string            { return proto.CompactTextString(m) }
func (*IMUsersStatReq) ProtoMessage()               {}
func (*IMUsersStatReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *IMUsersStatReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMUsersStatReq) GetUserIdList() []uint32 {
	if m != nil {
		return m.UserIdList
	}
	return nil
}

func (m *IMUsersStatReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMUsersStatRsp struct {
	// cmd id:		0x020b
	UserId           *uint32                   `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	UserStatList     []*IM_BaseDefine.UserStat `protobuf:"bytes,2,rep,name=user_stat_list" json:"user_stat_list,omitempty"`
	AttachData       []byte                    `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *IMUsersStatRsp) Reset()                    { *m = IMUsersStatRsp{} }
func (m *IMUsersStatRsp) String() string            { return proto.CompactTextString(m) }
func (*IMUsersStatRsp) ProtoMessage()               {}
func (*IMUsersStatRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IMUsersStatRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMUsersStatRsp) GetUserStatList() []*IM_BaseDefine.UserStat {
	if m != nil {
		return m.UserStatList
	}
	return nil
}

func (m *IMUsersStatRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMChangeAvatarReq struct {
	// cmd id:		0x020c
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	AvatarUrl        *string `protobuf:"bytes,2,req,name=avatar_url" json:"avatar_url,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMChangeAvatarReq) Reset()                    { *m = IMChangeAvatarReq{} }
func (m *IMChangeAvatarReq) String() string            { return proto.CompactTextString(m) }
func (*IMChangeAvatarReq) ProtoMessage()               {}
func (*IMChangeAvatarReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *IMChangeAvatarReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMChangeAvatarReq) GetAvatarUrl() string {
	if m != nil && m.AvatarUrl != nil {
		return *m.AvatarUrl
	}
	return ""
}

func (m *IMChangeAvatarReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMChangeAvatarRsp struct {
	// cmd id:		0x020d
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	ResultCode       *uint32 `protobuf:"varint,2,req,name=result_code" json:"result_code,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMChangeAvatarRsp) Reset()                    { *m = IMChangeAvatarRsp{} }
func (m *IMChangeAvatarRsp) String() string            { return proto.CompactTextString(m) }
func (*IMChangeAvatarRsp) ProtoMessage()               {}
func (*IMChangeAvatarRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *IMChangeAvatarRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMChangeAvatarRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMChangeAvatarRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

// 只给移动端通知
type IMPCLoginStatusNotify struct {
	// cmd id:		0x020e
	UserId           *uint32                     `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	LoginStat        *IM_BaseDefine.UserStatType `protobuf:"varint,2,req,name=login_stat,enum=IM.BaseDefine.UserStatType" json:"login_stat,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *IMPCLoginStatusNotify) Reset()                    { *m = IMPCLoginStatusNotify{} }
func (m *IMPCLoginStatusNotify) String() string            { return proto.CompactTextString(m) }
func (*IMPCLoginStatusNotify) ProtoMessage()               {}
func (*IMPCLoginStatusNotify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *IMPCLoginStatusNotify) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMPCLoginStatusNotify) GetLoginStat() IM_BaseDefine.UserStatType {
	if m != nil && m.LoginStat != nil {
		return *m.LoginStat
	}
	return IM_BaseDefine.UserStatType_USER_STATUS_ONLINE
}

type IMRemoveSessionNotify struct {
	// cmd id:		0x020f
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,2,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,3,req,name=session_id" json:"session_id,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMRemoveSessionNotify) Reset()                    { *m = IMRemoveSessionNotify{} }
func (m *IMRemoveSessionNotify) String() string            { return proto.CompactTextString(m) }
func (*IMRemoveSessionNotify) ProtoMessage()               {}
func (*IMRemoveSessionNotify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *IMRemoveSessionNotify) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMRemoveSessionNotify) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

func (m *IMRemoveSessionNotify) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

type IMDepartmentReq struct {
	// cmd id:		0x0210
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	LatestUpdateTime *uint32 `protobuf:"varint,2,req,name=latest_update_time" json:"latest_update_time,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMDepartmentReq) Reset()                    { *m = IMDepartmentReq{} }
func (m *IMDepartmentReq) String() string            { return proto.CompactTextString(m) }
func (*IMDepartmentReq) ProtoMessage()               {}
func (*IMDepartmentReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *IMDepartmentReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMDepartmentReq) GetLatestUpdateTime() uint32 {
	if m != nil && m.LatestUpdateTime != nil {
		return *m.LatestUpdateTime
	}
	return 0
}

func (m *IMDepartmentReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMDepartmentRsp struct {
	// cmd id:		0x0211
	UserId           *uint32                     `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	LatestUpdateTime *uint32                     `protobuf:"varint,2,req,name=latest_update_time" json:"latest_update_time,omitempty"`
	DeptList         []*IM_BaseDefine.DepartInfo `protobuf:"bytes,3,rep,name=dept_list" json:"dept_list,omitempty"`
	AttachData       []byte                      `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *IMDepartmentRsp) Reset()                    { *m = IMDepartmentRsp{} }
func (m *IMDepartmentRsp) String() string            { return proto.CompactTextString(m) }
func (*IMDepartmentRsp) ProtoMessage()               {}
func (*IMDepartmentRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *IMDepartmentRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMDepartmentRsp) GetLatestUpdateTime() uint32 {
	if m != nil && m.LatestUpdateTime != nil {
		return *m.LatestUpdateTime
	}
	return 0
}

func (m *IMDepartmentRsp) GetDeptList() []*IM_BaseDefine.DepartInfo {
	if m != nil {
		return m.DeptList
	}
	return nil
}

func (m *IMDepartmentRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMAvatarChangedNotify struct {
	// cmd id:		0x02012
	ChangedUserId    *uint32 `protobuf:"varint,1,req,name=changed_user_id" json:"changed_user_id,omitempty"`
	AvatarUrl        *string `protobuf:"bytes,2,req,name=avatar_url" json:"avatar_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMAvatarChangedNotify) Reset()                    { *m = IMAvatarChangedNotify{} }
func (m *IMAvatarChangedNotify) String() string            { return proto.CompactTextString(m) }
func (*IMAvatarChangedNotify) ProtoMessage()               {}
func (*IMAvatarChangedNotify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *IMAvatarChangedNotify) GetChangedUserId() uint32 {
	if m != nil && m.ChangedUserId != nil {
		return *m.ChangedUserId
	}
	return 0
}

func (m *IMAvatarChangedNotify) GetAvatarUrl() string {
	if m != nil && m.AvatarUrl != nil {
		return *m.AvatarUrl
	}
	return ""
}

type IMChangeSignInfoReq struct {
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SignInfo         *string `protobuf:"bytes,2,req,name=sign_info" json:"sign_info,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMChangeSignInfoReq) Reset()                    { *m = IMChangeSignInfoReq{} }
func (m *IMChangeSignInfoReq) String() string            { return proto.CompactTextString(m) }
func (*IMChangeSignInfoReq) ProtoMessage()               {}
func (*IMChangeSignInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *IMChangeSignInfoReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMChangeSignInfoReq) GetSignInfo() string {
	if m != nil && m.SignInfo != nil {
		return *m.SignInfo
	}
	return ""
}

func (m *IMChangeSignInfoReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMChangeSignInfoRsp struct {
	// cmd id:		0x0214
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	ResultCode       *uint32 `protobuf:"varint,2,req,name=result_code" json:"result_code,omitempty"`
	SignInfo         *string `protobuf:"bytes,3,opt,name=sign_info" json:"sign_info,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMChangeSignInfoRsp) Reset()                    { *m = IMChangeSignInfoRsp{} }
func (m *IMChangeSignInfoRsp) String() string            { return proto.CompactTextString(m) }
func (*IMChangeSignInfoRsp) ProtoMessage()               {}
func (*IMChangeSignInfoRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *IMChangeSignInfoRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMChangeSignInfoRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMChangeSignInfoRsp) GetSignInfo() string {
	if m != nil && m.SignInfo != nil {
		return *m.SignInfo
	}
	return ""
}

func (m *IMChangeSignInfoRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

// 个性签名修改通知（广播）
type IMSignInfoChangedNotify struct {
	// cmd id:		0x0215
	ChangedUserId    *uint32 `protobuf:"varint,1,req,name=changed_user_id" json:"changed_user_id,omitempty"`
	SignInfo         *string `protobuf:"bytes,2,req,name=sign_info" json:"sign_info,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMSignInfoChangedNotify) Reset()                    { *m = IMSignInfoChangedNotify{} }
func (m *IMSignInfoChangedNotify) String() string            { return proto.CompactTextString(m) }
func (*IMSignInfoChangedNotify) ProtoMessage()               {}
func (*IMSignInfoChangedNotify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *IMSignInfoChangedNotify) GetChangedUserId() uint32 {
	if m != nil && m.ChangedUserId != nil {
		return *m.ChangedUserId
	}
	return 0
}

func (m *IMSignInfoChangedNotify) GetSignInfo() string {
	if m != nil && m.SignInfo != nil {
		return *m.SignInfo
	}
	return ""
}

type IMUsersInfoByNameReq struct {
	// cmd id:		0x0216
	UserId           *uint32  `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	NameList         [][]byte `protobuf:"bytes,2,rep,name=name_list" json:"name_list,omitempty"`
	ClientData       []byte   `protobuf:"bytes,3,opt,name=client_data" json:"client_data,omitempty"`
	AttachData       []byte   `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *IMUsersInfoByNameReq) Reset()                    { *m = IMUsersInfoByNameReq{} }
func (m *IMUsersInfoByNameReq) String() string            { return proto.CompactTextString(m) }
func (*IMUsersInfoByNameReq) ProtoMessage()               {}
func (*IMUsersInfoByNameReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *IMUsersInfoByNameReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMUsersInfoByNameReq) GetNameList() [][]byte {
	if m != nil {
		return m.NameList
	}
	return nil
}

func (m *IMUsersInfoByNameReq) GetClientData() []byte {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *IMUsersInfoByNameReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMUsersInfoByNameRsp struct {
	// cmd id:		0x0217
	UserId           *uint32                   `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	UserInfoList     []*IM_BaseDefine.UserInfo `protobuf:"bytes,2,rep,name=user_info_list" json:"user_info_list,omitempty"`
	ClientData       []byte                    `protobuf:"bytes,3,opt,name=client_data" json:"client_data,omitempty"`
	AttachData       []byte                    `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *IMUsersInfoByNameRsp) Reset()                    { *m = IMUsersInfoByNameRsp{} }
func (m *IMUsersInfoByNameRsp) String() string            { return proto.CompactTextString(m) }
func (*IMUsersInfoByNameRsp) ProtoMessage()               {}
func (*IMUsersInfoByNameRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *IMUsersInfoByNameRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMUsersInfoByNameRsp) GetUserInfoList() []*IM_BaseDefine.UserInfo {
	if m != nil {
		return m.UserInfoList
	}
	return nil
}

func (m *IMUsersInfoByNameRsp) GetClientData() []byte {
	if m != nil {
		return m.ClientData
	}
	return nil
}

func (m *IMUsersInfoByNameRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

func init() {
	proto.RegisterType((*IMRecentContactSessionReq)(nil), "IM.Buddy.IMRecentContactSessionReq")
	proto.RegisterType((*IMRecentContactSessionRsp)(nil), "IM.Buddy.IMRecentContactSessionRsp")
	proto.RegisterType((*IMUserStatNotify)(nil), "IM.Buddy.IMUserStatNotify")
	proto.RegisterType((*IMUsersInfoReq)(nil), "IM.Buddy.IMUsersInfoReq")
	proto.RegisterType((*IMUsersInfoRsp)(nil), "IM.Buddy.IMUsersInfoRsp")
	proto.RegisterType((*IMRemoveSessionReq)(nil), "IM.Buddy.IMRemoveSessionReq")
	proto.RegisterType((*IMRemoveSessionRsp)(nil), "IM.Buddy.IMRemoveSessionRsp")
	proto.RegisterType((*IMAllUserReq)(nil), "IM.Buddy.IMAllUserReq")
	proto.RegisterType((*IMAllUserRsp)(nil), "IM.Buddy.IMAllUserRsp")
	proto.RegisterType((*IMUsersStatReq)(nil), "IM.Buddy.IMUsersStatReq")
	proto.RegisterType((*IMUsersStatRsp)(nil), "IM.Buddy.IMUsersStatRsp")
	proto.RegisterType((*IMChangeAvatarReq)(nil), "IM.Buddy.IMChangeAvatarReq")
	proto.RegisterType((*IMChangeAvatarRsp)(nil), "IM.Buddy.IMChangeAvatarRsp")
	proto.RegisterType((*IMPCLoginStatusNotify)(nil), "IM.Buddy.IMPCLoginStatusNotify")
	proto.RegisterType((*IMRemoveSessionNotify)(nil), "IM.Buddy.IMRemoveSessionNotify")
	proto.RegisterType((*IMDepartmentReq)(nil), "IM.Buddy.IMDepartmentReq")
	proto.RegisterType((*IMDepartmentRsp)(nil), "IM.Buddy.IMDepartmentRsp")
	proto.RegisterType((*IMAvatarChangedNotify)(nil), "IM.Buddy.IMAvatarChangedNotify")
	proto.RegisterType((*IMChangeSignInfoReq)(nil), "IM.Buddy.IMChangeSignInfoReq")
	proto.RegisterType((*IMChangeSignInfoRsp)(nil), "IM.Buddy.IMChangeSignInfoRsp")
	proto.RegisterType((*IMSignInfoChangedNotify)(nil), "IM.Buddy.IMSignInfoChangedNotify")
	proto.RegisterType((*IMUsersInfoByNameReq)(nil), "IM.Buddy.IMUsersInfoByNameReq")
	proto.RegisterType((*IMUsersInfoByNameRsp)(nil), "IM.Buddy.IMUsersInfoByNameRsp")
}

func init() { proto.RegisterFile("IM.Buddy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x55, 0x97, 0x9f, 0xf4, 0xa3, 0xb7, 0x61, 0x63, 0x69, 0x51, 0xbb, 0xed, 0xa5, 0xe4, 0x69,
	0x42, 0xa8, 0xa0, 0x7d, 0x00, 0xd0, 0xfe, 0x20, 0x11, 0x69, 0xa9, 0xa6, 0x15, 0x1e, 0x78, 0x40,
	0x91, 0x49, 0xdc, 0xcc, 0x28, 0xb1, 0x43, 0x7c, 0x33, 0xa9, 0x12, 0x0f, 0xc0, 0x77, 0xe0, 0xfb,
	0x22, 0xc7, 0x8d, 0xd6, 0x76, 0x76, 0xd7, 0xa1, 0xbe, 0xb5, 0xb6, 0x73, 0xee, 0x39, 0xe7, 0x9e,
	0xeb, 0x04, 0x76, 0x83, 0x70, 0x74, 0x56, 0x25, 0xc9, 0x6c, 0x54, 0x94, 0x02, 0x85, 0xf7, 0xa4,
	0xf9, 0x7f, 0xd8, 0x55, 0xbf, 0x88, 0xa4, 0x17, 0x74, 0xca, 0x38, 0xd5, 0xdb, 0xfe, 0x17, 0x38,
	0x08, 0xc2, 0x6b, 0x1a, 0x53, 0x8e, 0xe7, 0x82, 0x23, 0x89, 0x71, 0x42, 0xa5, 0x64, 0x82, 0x5f,
	0xd3, 0xef, 0xde, 0x1e, 0xfc, 0x5f, 0x49, 0x5a, 0x46, 0x2c, 0x19, 0xb4, 0x86, 0x3b, 0xc7, 0x4f,
	0xbd, 0x43, 0xf0, 0x32, 0x82, 0x54, 0x62, 0x54, 0x15, 0x09, 0x41, 0x1a, 0x21, 0xcb, 0xe9, 0x60,
	0xa7, 0xde, 0xeb, 0x42, 0x87, 0x20, 0x92, 0xf8, 0x26, 0x4a, 0x08, 0x92, 0x41, 0x6f, 0xd8, 0x3a,
	0x76, 0xfd, 0xdf, 0x2d, 0x2b, 0xbe, 0x2c, 0xee, 0xe3, 0xbf, 0x83, 0x5e, 0xac, 0x4f, 0x45, 0x52,
	0x1f, 0x8b, 0x32, 0x26, 0x71, 0xb0, 0x33, 0x74, 0x8e, 0x3b, 0x27, 0x2f, 0x46, 0xcb, 0x0a, 0x96,
	0x01, 0x03, 0x3e, 0x15, 0x66, 0x12, 0x6f, 0xe1, 0x59, 0x10, 0x7e, 0x92, 0xb4, 0x9c, 0x20, 0xc1,
	0xb1, 0x40, 0x36, 0x9d, 0x79, 0x2f, 0xa1, 0x5d, 0x97, 0x96, 0x48, 0xb0, 0x2e, 0xde, 0x39, 0xe9,
	0xaf, 0xc0, 0x37, 0x4f, 0xf8, 0x97, 0xca, 0x54, 0xf5, 0x4f, 0xaa, 0x1a, 0x46, 0x63, 0x7a, 0xe0,
	0xce, 0x17, 0xee, 0x08, 0x5b, 0x2c, 0x49, 0x97, 0xd1, 0x4c, 0x36, 0xbc, 0x86, 0x5d, 0xbd, 0xc0,
	0xa7, 0x62, 0xd1, 0x00, 0x13, 0x43, 0xbb, 0xec, 0x9f, 0x2d, 0xf0, 0x94, 0xf7, 0xb9, 0xb8, 0xa5,
	0xeb, 0x9a, 0xfa, 0x06, 0xdc, 0xc6, 0x6c, 0x9c, 0x15, 0xba, 0x9d, 0xbb, 0x27, 0x87, 0x2b, 0xb5,
	0xe6, 0x08, 0x1f, 0x67, 0x05, 0xf5, 0x3c, 0x80, 0xe6, 0x09, 0x96, 0x0c, 0x1c, 0x7b, 0xfb, 0xff,
	0x18, 0x28, 0x98, 0x04, 0x77, 0xa1, 0x53, 0x52, 0x59, 0x65, 0x18, 0xc5, 0x22, 0x69, 0x02, 0xb5,
	0xca, 0xcb, 0x79, 0x24, 0xaf, 0xff, 0xec, 0xbc, 0xae, 0xc0, 0x0d, 0xc2, 0xd3, 0x2c, 0x53, 0x06,
	0x6e, 0x27, 0xe8, 0x3f, 0x16, 0x11, 0x4d, 0x12, 0xd7, 0x21, 0x36, 0x61, 0xac, 0x5b, 0xed, 0xfc,
	0x43, 0xab, 0xef, 0x12, 0xaa, 0x02, 0xbb, 0xbd, 0x84, 0xd6, 0x68, 0xeb, 0x12, 0xaa, 0xc6, 0xe7,
	0xa1, 0x84, 0x2a, 0x10, 0x73, 0xa1, 0x10, 0xf6, 0x83, 0xf0, 0xfc, 0x86, 0xf0, 0x94, 0x9e, 0xde,
	0x12, 0x24, 0xe6, 0x5e, 0x78, 0x00, 0xa4, 0xde, 0x8d, 0xaa, 0x32, 0xab, 0x1d, 0x6b, 0x9b, 0xe1,
	0xc6, 0xf7, 0xe0, 0x36, 0xce, 0x9a, 0x11, 0xef, 0x33, 0x3c, 0x0f, 0xc2, 0xab, 0xf3, 0x4b, 0x91,
	0x32, 0xae, 0x44, 0x54, 0x72, 0x7e, 0x79, 0x18, 0xec, 0x80, 0x4c, 0x9d, 0xd2, 0xd7, 0x89, 0x1e,
	0xa0, 0x23, 0x8b, 0x15, 0x2a, 0xa9, 0x3e, 0x57, 0xd0, 0x4b, 0x73, 0x61, 0x83, 0xde, 0xca, 0x74,
	0xfa, 0x13, 0xd8, 0x0b, 0xc2, 0x0b, 0x5a, 0x90, 0x12, 0x73, 0xca, 0x71, 0x3b, 0x99, 0xff, 0xd5,
	0x5a, 0x41, 0x7d, 0x6c, 0xee, 0x5f, 0x41, 0x3b, 0xa1, 0x05, 0x2e, 0xe6, 0xfe, 0x60, 0x45, 0x98,
	0x46, 0xb7, 0x27, 0xff, 0x42, 0x19, 0xa9, 0xbb, 0xad, 0x3b, 0x9f, 0xcc, 0x8d, 0xec, 0xc3, 0x5e,
	0xac, 0x17, 0xa2, 0x07, 0xe3, 0xe4, 0x8f, 0xa1, 0xdb, 0x24, 0x67, 0xc2, 0x52, 0x6e, 0xbd, 0xe6,
	0xf7, 0xa1, 0x2d, 0x59, 0xca, 0xeb, 0x8b, 0x79, 0x5d, 0x12, 0x63, 0x03, 0xde, 0xc6, 0x59, 0x5c,
	0x2a, 0xe2, 0x0c, 0x5b, 0xb6, 0x22, 0xef, 0xa1, 0x1f, 0x84, 0x0d, 0xfc, 0x86, 0xe2, 0xef, 0x0b,
	0xf0, 0x13, 0xe8, 0x2d, 0xbc, 0x8f, 0xce, 0x66, 0x63, 0x92, 0x53, 0x9b, 0x78, 0x4e, 0x72, 0x7a,
	0x37, 0xee, 0xae, 0xe2, 0x15, 0x67, 0x8c, 0x72, 0xd4, 0xbc, 0x14, 0x59, 0xd7, 0xfa, 0x21, 0x60,
	0x28, 0xb3, 0xad, 0x97, 0xdf, 0x66, 0x24, 0xce, 0x8e, 0xa0, 0x1f, 0x8b, 0x7c, 0x94, 0x8b, 0xb4,
	0xfa, 0xc6, 0xe8, 0x08, 0x51, 0x7f, 0x04, 0x7d, 0xad, 0xa6, 0x1f, 0x9c, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0x2b, 0xa8, 0xa5, 0x39, 0x09, 0x00, 0x00,
}
