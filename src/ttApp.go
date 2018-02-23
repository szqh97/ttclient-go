package main

import (
	"./TT"
	"flag"
	"log"
	"os"
	"os/signal"
)

var username = flag.String("username", "dj352801", "tt username")

func Main(user string) {
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan)

	ip, port, _ := TT.GetMsgServerAddress()
	msgaddr := ip + ":" + port
	client := TT.NewClientConn(msgaddr, user)
	client.Login()
	client.CheckLogIn()

	//	client.GetRecentSession(0)
	//client.GetUserInfoList([]uint32{12460, 2396, 54555})
	///client.GetUnreadMsgCnt(client.UserId)
	//client.GetMsgList(52866, 0, 0)

	/*
		func (client *ClientConn) SendCustomerMsg(customerId, toCustomerId, toSubId uint32)
	*/

	client.SendCustomerMsg(2396, 52866, 0)
	/*
		client.GetCustomerInfoReq(2396)
		client.GetForwardingUserListReq(2396)
	*/

	/*
		func (client *ClientConn) CustomerFowardingReq(customerId, fromId, toSubId uint32)
	*/

	//client.CustomerFowardingReq(2396, 52866, 12460)

	select {
	case <-exitChan:
		log.Println("exiting ... ")
	}
}

func main() {
	flag.Parse()

	Main(*username)

}
