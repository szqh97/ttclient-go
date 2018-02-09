package TT

import (
	"../IM/IM_BaseDefine"
	"../IM/IM_Buddy"
	"../IM/IM_Login"
	"../IM/IM_Message"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"log"
)

func (client *ClientConn) handlePdu(pdu *ImPdu) {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Printf("entering handlePdu..., pdu is %v\n", pdu)
	switch pdu.command_id {
	case uint16(IM_BaseDefine.LoginCmdID_CID_LOGIN_RES_USERLOGIN):
		msg := &IM_Login.IMLoginRes{}
		proto.Unmarshal(pdu.msg, msg)
		log.Println(msg)
		client.UserId = msg.GetUserInfo().GetUserId()
		client.IsLogin = true
		log.Println(msg.GetUserInfo().GetCustomerId())

	case uint16(IM_BaseDefine.BuddyListCmdID_CID_BUDDY_LIST_RECENT_CONTACT_SESSION_RESPONSE):
		msg := &IM_Buddy.IMRecentContactSessionRsp{}
		proto.Unmarshal(pdu.msg, msg)
		out, _ := json.Marshal(msg)
		log.Println(msg.String(), string(out))

	case uint16(IM_BaseDefine.BuddyListCmdID_CID_BUDDY_LIST_USER_INFO_RESPONSE):
		msg := &IM_Buddy.IMUsersInfoRsp{}
		proto.Unmarshal(pdu.msg, msg)
		out, _ := json.Marshal(msg)
		log.Println("Get user list info")
		log.Println(string(out))
	case uint16(IM_BaseDefine.MessageCmdID_CID_MSG_UNREAD_CNT_RESPONSE):
		msg := &IM_Message.IMUnreadMsgCntRsp{}
		proto.Unmarshal(pdu.msg, msg)
		out, _ := json.Marshal(msg)
		log.Println(" get unread msg count")
		log.Println(string(out))

	case uint16(IM_BaseDefine.MessageCmdID_CID_MSG_LIST_RESPONSE):
		msg := &IM_Message.IMGetMsgListRsp{}
		proto.Unmarshal(pdu.msg, msg)
		out, _ := json.Marshal(msg)
		log.Println("get mst list response:")
		log.Println(string(out))

	default:
		log.Fatal("Invalid commdd id ", pdu.command_id)
	}
}
