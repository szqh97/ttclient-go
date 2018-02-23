package TT

import (
	"../IM/IM_BaseDefine"
	"../IM/IM_Buddy"
	"../IM/IM_Customer"
	"../IM/IM_Login"
	"../IM/IM_Message"
	"../ImPduBase"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	log "log"
)

func (client *ClientConn) handlePdu(pdu ImPduBase.ImPdu) {

	log.Printf("entering handlePdu..., pdu is %v\n", pdu)
	//switch pdu.command_id {
	log.Println("CommandId id :", pdu.GetCommandId())
	switch pdu.GetCommandId() {

	case int32(IM_BaseDefine.LoginCmdID_CID_LOGIN_RES_USERLOGIN):
		msg := &IM_Login.IMLoginRes{}
		proto.Unmarshal(pdu.GetMsgData(), msg)
		log.Println(msg)
		client.UserId = msg.GetUserInfo().GetUserId()
		client.IsLogin = true
		client.loginChan <- true
		log.Println(msg.GetUserInfo().GetCustomerId())

	case int32(IM_BaseDefine.BuddyListCmdID_CID_BUDDY_LIST_RECENT_CONTACT_SESSION_RESPONSE):
		msg := &IM_Buddy.IMRecentContactSessionRsp{}
		proto.Unmarshal(pdu.GetMsgData(), msg)
		out, _ := json.Marshal(msg)
		log.Println(msg.String(), string(out))

	case int32(IM_BaseDefine.MessageCmdID_CID_MSG_DATA):
		msg := &IM_Message.IMMsgData{}
		proto.Unmarshal(pdu.GetMsgData(), msg)
		out, _ := json.Marshal(msg)
		log.Println("receive msg data")
		log.Println(string(out))

	case int32(IM_BaseDefine.MessageCmdID_CID_MSG_DATA_ACK):
		msg := &IM_Message.IMMsgDataAck{}
		proto.Unmarshal(pdu.GetMsgData(), msg)
		out, _ := json.Marshal(msg)
		log.Println("receive msg data ack")
		log.Println(string(out))

	case int32(IM_BaseDefine.BuddyListCmdID_CID_BUDDY_LIST_USER_INFO_RESPONSE):
		msg := &IM_Buddy.IMUsersInfoRsp{}
		proto.Unmarshal(pdu.GetMsgData(), msg)
		out, _ := json.Marshal(msg)
		log.Println("Get user list info")
		log.Println(string(out))

	case int32(IM_BaseDefine.MessageCmdID_CID_MSG_UNREAD_CNT_RESPONSE):
		msg := &IM_Message.IMUnreadMsgCntRsp{}
		proto.Unmarshal(pdu.GetMsgData(), msg)
		out, _ := json.Marshal(msg)
		log.Println(" get unread msg count")
		log.Println(string(out))

	case int32(IM_BaseDefine.MessageCmdID_CID_MSG_LIST_RESPONSE):
		msg := &IM_Message.IMGetMsgListRsp{}
		proto.Unmarshal(pdu.GetMsgData(), msg)
		out, _ := json.Marshal(msg)
		log.Println("get mst list response:")
		log.Println(string(out))

	case int32(IM_BaseDefine.CustomerCmdId_CID_CUSTOMER_CUSTOMER_INFO_RSP):
		log.Println("get cusomter info resp")
		msg := &IM_Customer.IMCustomerInfoRsp{}
		proto.Unmarshal(pdu.GetMsgData(), msg)
		out, _ := json.Marshal(msg)

		log.Println(string(out))

	case int32(IM_BaseDefine.CustomerCmdId_CID_CUSTOMER_GET_FORWORDING_LIST_RSP):
		log.Println("get forwarding list resp")
		msg := &IM_Customer.IMGetForwardingUserListRsp{}
		proto.Unmarshal(pdu.GetMsgData(), msg)
		out, _ := json.Marshal(msg)
		log.Println(string(out))

	case int32(IM_BaseDefine.CustomerCmdId_CID_CUSTOMER_FORWORDING_RSP):
		log.Println("forwading resp")
		msg := &IM_Customer.IMCustomerForwardingRsp{}
		proto.Unmarshal(pdu.GetMsgData(), msg)
		out, _ := json.Marshal(msg)
		log.Println(string(out))

	case int32(IM_BaseDefine.OtherCmdID_CID_OTHER_HEARTBEAT):
		return

	default:
		log.Println("Invalid commdd id ", pdu.GetCommandId())
	}
}
