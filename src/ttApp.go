package main

import (
	"./TT"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan)

	ip, port, _ := TT.GetMsgServerAddress()
	msgaddr := ip + ":" + port
	client := TT.NewClientConn(msgaddr, "dj352801")
	client.Login()
	client.CheckLoIn()

	//	client.GetRecentSession(0)
	client.GetUserInfoList([]uint32{12460, 2396, 54555})
	client.GetUnreadMsgCnt(client.UserId)
	client.GetMsgList(52866, 0, 0)

	select {
	case <-exitChan:
		fmt.Println("exiting ... ")
	}
}
