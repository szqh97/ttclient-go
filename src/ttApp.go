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
	client.CheckLoIn()

	//	client.GetRecentSession(0)
	client.GetUserInfoList([]uint32{12460, 2396, 54555})
	client.GetUnreadMsgCnt(client.UserId)
	client.GetMsgList(52866, 0, 0)

	select {
	case <-exitChan:
		log.Println("exiting ... ")
	}
}

func main() {
	flag.Parse()

	Main(*username)

}
