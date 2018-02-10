package ImPduBase

import (
	"log"
	"testing"
)

func Test_NewPdu(t *testing.T) {

	log.Printf("xxx")
	pdu := ImPdu{}
	pdu.ServiceId = 1
	pdu.CommandId = 2
	pdu.SetMsgData([]byte{1, 2, 3, 3})

	log.Println(Encode(pdu))
	out := Encode(pdu)
	var pdu2 ImPdu
	err := Decode(out, &pdu2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pdu2)

}
