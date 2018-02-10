package TT

import (
	"../ImPduBase"
	"bufio"
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"time"
)

type ClientConn struct {
	UserName string
	Passwd   string
	UserId   uint32
	IsLogin  bool

	conn           *net.TCPConn
	pduReceiveChan chan *ImPduBase.ImPdu
	pduSenderChan  chan *ImPduBase.ImPdu
	exitChan       chan bool
	loginChan      chan bool
	waitGroup      WaitGrapWrapper
}

func NewClientConn(MsgAddr, UserName string) *ClientConn {

	o1 := md5.Sum([]byte(UserName))
	o2 := md5.Sum([]byte(fmt.Sprintf("%x", o1)))
	passwd := fmt.Sprintf("%x", o2)

	addr, err := net.ResolveTCPAddr("tcp4", MsgAddr)
	checkErr(err)
	c, err := net.DialTCP("tcp", nil, addr)
	checkErr(err)
	client := ClientConn{
		UserName:       UserName,
		Passwd:         passwd,
		conn:           c,
		pduReceiveChan: make(chan *ImPduBase.ImPdu),
		pduSenderChan:  make(chan *ImPduBase.ImPdu),
		loginChan:      make(chan bool),
		exitChan:       make(chan bool),
	}

	client.waitGroup.Wrap(func() { client.readLoop() })
	client.waitGroup.Wrap(func() { client.writeLoop() })
	client.waitGroup.Wrap(func() { client.handleLoop() })

	return &client
}
func (client *ClientConn) CheckLoIn() {
	select {
	case <-client.loginChan:
		log.Printf("client %v logged in...")
	}

}

func (client *ClientConn) writeLoop() {
	for {
		select {
		case pdu := <-client.pduReceiveChan:

			cnt, err := client.conn.Write(ImPduBase.Encode(*pdu))
			if err != nil {
				fmt.Println("write file failed", cnt, err.Error())
			}
		default:
		}
	}
}

func (client *ClientConn) readLoop() {
	reader := bufio.NewReader(client.conn)
	log.Println("entering read loop...")
	for {
		var pdu ImPduBase.ImPdu
		err := ImPduBase.ImPduReader(reader, &pdu)
		if err != nil {
			return
		}
		client.pduSenderChan <- &pdu
	}

}
func (client *ClientConn) handleLoop() {
	for {
		select {
		case pdu := <-client.pduSenderChan:
			client.handlePdu(*pdu)
		case <-time.After(time.Second * 30):
			client.heartbeat()
		default:
			log.Fatal("unkonw loop...")
		}
	}
}
