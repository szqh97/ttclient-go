package TT

import (
	"../IM/IM_BaseDefine"
	"../IM/IM_Buddy"
	"../IM/IM_Customer"
	"../IM/IM_Login"
	"../IM/IM_Message"
	"../IM/IM_Other"
	"../ImPduBase"
	"github.com/golang/protobuf/proto"
	"log"
)

func (client *ClientConn) request(pdu *ImPduBase.ImPdu) {
	client.pduReceiveChan <- pdu
}

func (client *ClientConn) heartbeat() {
	msg := IM_Other.IMHeartBeat{}
	out, err := proto.Marshal(&msg)
	if err != nil {
		log.Println(err)
		return
	}
	var pdu ImPduBase.ImPdu
	pdu.Reset()
	pdu.SetServiceId(int32(IM_BaseDefine.OtherCmdID_CID_OTHER_HEARTBEAT))
	pdu.SetCommandId(int32(IM_BaseDefine.OtherCmdID_CID_OTHER_HEARTBEAT))
	pdu.SetMsgData(out)
	client.request(&pdu)
}

func (client *ClientConn) Login() {
	msg := IM_Login.IMLoginReq{}
	msg.UserName = &client.UserName
	msg.Password = &client.Passwd
	online := IM_BaseDefine.UserStatType_USER_STATUS_ONLINE
	w := IM_BaseDefine.ClientType_CLIENT_TYPE_ANDROID
	version := "v1.1.0"
	msg.OnlineStatus = &online
	msg.ClientType = &w
	msg.ClientVersion = &version
	out, err := proto.Marshal(&msg)
	if err != nil {
		log.Println(err)
	}
	var pdu ImPduBase.ImPdu
	pdu.Reset()
	pdu.SetServiceId(int32(IM_BaseDefine.ServiceID_SID_LOGIN))
	pdu.SetCommandId(int32(IM_BaseDefine.LoginCmdID_CID_LOGIN_REQ_USERLOGIN))
	pdu.SetMsgData(out)
	log.Println("login request...", pdu)

	client.pduReceiveChan <- &pdu
}

func (client *ClientConn) GetRecentSession(ts uint32) {
	msg := IM_Buddy.IMRecentContactSessionReq{
		UserId:           &client.UserId,
		LatestUpdateTime: &ts,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		log.Println(err)
		return
	}

	var pdu ImPduBase.ImPdu
	pdu.Reset()
	pdu.SetServiceId(int32(IM_BaseDefine.ServiceID_SID_BUDDY_LIST))
	pdu.SetCommandId(int32(IM_BaseDefine.BuddyListCmdID_CID_BUDDY_LIST_RECENT_CONTACT_SESSION_REQUEST))
	pdu.SetMsgData(out)

	log.Println("get recent session")
	client.pduReceiveChan <- &pdu
}

func (client *ClientConn) GetUserInfoList(user_id_list []uint32) {
	msg := IM_Buddy.IMUsersInfoReq{
		UserId:     &client.UserId,
		UserIdList: user_id_list,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		log.Println(err)
		return
	}

	var pdu ImPduBase.ImPdu
	pdu.Reset()
	pdu.SetServiceId(int32(IM_BaseDefine.ServiceID_SID_BUDDY_LIST))
	pdu.SetCommandId(int32(IM_BaseDefine.BuddyListCmdID_CID_BUDDY_LIST_USER_INFO_REQUEST))
	pdu.SetMsgData(out)

	log.Println("get user info list request")
	client.request(&pdu)
}

func (client *ClientConn) GetUnreadMsgCnt(nReqSubId uint32) {
	msg := IM_Message.IMUnreadMsgCntReq{
		UserId:   &client.UserId,
		ReqSubId: &nReqSubId,
	}
	out, err := proto.Marshal(&msg)
	checkErr(err)

	var pdu ImPduBase.ImPdu
	pdu.Reset()
	pdu.SetServiceId(int32(IM_BaseDefine.ServiceID_SID_MSG))
	pdu.SetCommandId(int32(IM_BaseDefine.MessageCmdID_CID_MSG_UNREAD_CNT_REQUEST))
	pdu.SetMsgData(out)

	log.Println("get unread msg count")
	client.request(&pdu)
}

func (client *ClientConn) GetMsgList(s_id, msgid, cnt uint32) {
	stype := IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
	msg := IM_Message.IMGetMsgListReq{
		UserId:      &client.UserId,
		SessionType: &stype,
		SessionId:   &s_id,
		MsgIdBegin:  &msgid,
		MsgCnt:      &cnt,
	}
	out, err := proto.Marshal(&msg)
	checkErr(err)

	var pdu ImPduBase.ImPdu
	pdu.Reset()
	pdu.SetServiceId(int32(IM_BaseDefine.ServiceID_SID_MSG))
	pdu.SetCommandId(int32(IM_BaseDefine.MessageCmdID_CID_MSG_LIST_REQUEST))
	pdu.SetMsgData(out)

	log.Println("get msg list request")
	client.request(&pdu)
}
func (client *ClientConn) SendCustomerMsg(customerId, toCustomerId, toSubId uint32) {
	msgdata := []byte("0gwvBXpM2gamndcPpQ1MKnMrh7eB2LTQ+m7yLHXZ3MQ=")
	var msgid uint32 = 0
	var createtm uint32 = 0
	msgtype := IM_BaseDefine.MsgType_MSG_TYPE_SINGLE_TEXT
	var msgTypeEx uint32 = 0x21
	msg := IM_Message.IMMsgData{
		FromUserId:  &customerId,
		ToSessionId: &toCustomerId,
		MsgId:       &msgid,
		CreateTime:  &createtm,
		MsgType:     &msgtype,
		MsgData:     msgdata,
		MsgTypeEx:   &msgTypeEx,
		FromSubId:   &client.UserId,
		ToSubId:     &toSubId,
	}
	out, err := proto.Marshal(&msg)
	checkErr(err)

	var pdu ImPduBase.ImPdu
	pdu.Reset()
	pdu.SetServiceId(int32(IM_BaseDefine.ServiceID_SID_MSG))
	pdu.SetCommandId(int32(IM_BaseDefine.MessageCmdID_CID_MSG_DATA))
	pdu.SetMsgData(out)

	log.Println("send customer msg")
	client.request(&pdu)
}

func (client *ClientConn) GetCustomerInfoReq(queryuserId uint32) {
	msg := IM_Customer.IMCustomerInfoReq{
		ReqUserId: &client.UserId,
		UserId:    &queryuserId,
	}
	out, err := proto.Marshal(&msg)
	checkErr(err)
	var pdu ImPduBase.ImPdu
	pdu.Reset()
	pdu.SetServiceId(int32(IM_BaseDefine.ServiceID_SID_CUSTOMER))
	pdu.SetCommandId(int32(IM_BaseDefine.CustomerCmdId_CID_CUSTOMER_CUSTOMER_INFO_REQ))
	pdu.SetMsgData(out)
	log.Println("get customer info req")
	client.request(&pdu)
}

func (client *ClientConn) GetForwardingUserListReq(customerId uint32) {
	msg := IM_Customer.IMGetForwardingUserListReq{
		ReqUserId:  &client.UserId,
		CustomerId: &customerId,
	}

	out, err := proto.Marshal(&msg)
	checkErr(err)
	var pdu ImPduBase.ImPdu
	pdu.Reset()
	pdu.SetServiceId(int32(IM_BaseDefine.ServiceID_SID_CUSTOMER))
	pdu.SetCommandId(int32(IM_BaseDefine.CustomerCmdId_CID_CUSTOMER_GET_FORWORDING_LIST_REQ))
	pdu.SetMsgData(out)
	log.Println("get customer forwarding list req")
	client.request(&pdu)
}

func (client *ClientConn) CustomerFowardingReq(customerId, fromId, toSubId uint32) {
	msg := IM_Customer.IMCustomerForwardingReq{
		ReqUserId:  &client.UserId,
		CustomerId: &customerId,
		FromId:     &fromId,
		ToSubId:    &toSubId,
	}
	out, err := proto.Marshal(&msg)
	checkErr(err)
	var pdu ImPduBase.ImPdu
	pdu.Reset()
	pdu.SetServiceId(int32(IM_BaseDefine.ServiceID_SID_CUSTOMER))
	pdu.SetCommandId(int32(IM_BaseDefine.CustomerCmdId_CID_CUSTOMER_FORWORDING_REQ))
	pdu.SetMsgData(out)
	log.Println("forwarding msg req")
	client.request(&pdu)
}
