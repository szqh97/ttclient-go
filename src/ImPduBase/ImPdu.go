// Package ImPdu provides ...
package ImPduBase

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	log "log"
)

type Header struct {
	Length    uint32
	Version   uint16
	Flag      uint16
	ServiceId uint16
	CommandId uint16
	SeqNum    uint16
	Reserved  uint16
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

var (
	HeaderLen int = binary.Size(Header{})
)

type ImPdu struct {
	Header
	MsgData []byte
}

func (p *ImPdu) Reset() {
	*p = ImPdu{}
}

func (p *ImPdu) GetLenth() uint32 {
	if p != nil {
		return p.Length
	}
	return 0
}

func (p *ImPdu) GetServiceId() int32 {
	if p != nil {
		return int32(p.ServiceId)
	}
	return 0
}

func (p *ImPdu) GetCommandId() int32 {
	if p != nil {
		return int32(p.CommandId)
	}
	return 0
}

func (p *ImPdu) GetMsgData() []byte {
	if p != nil {
		return p.MsgData
	}
	return nil
}
func (p *ImPdu) SetServiceId(sid interface{}) {
	if p != nil {
		p.ServiceId = uint16(sid.(int32))
	}
}
func (p *ImPdu) SetCommandId(cid interface{}) {

	if p != nil {
		p.CommandId = uint16(cid.(int32))
	}
}
func (p *ImPdu) SetMsgData(msg []byte) {
	p.Length = uint32(HeaderLen + len(msg))
	p.MsgData = msg
}

func Encode(p ImPdu) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, p.Length)
	binary.Write(buf, binary.BigEndian, p.Version)
	binary.Write(buf, binary.BigEndian, p.Flag)
	binary.Write(buf, binary.BigEndian, p.ServiceId)
	binary.Write(buf, binary.BigEndian, p.CommandId)
	binary.Write(buf, binary.BigEndian, p.SeqNum)
	binary.Write(buf, binary.BigEndian, p.Reserved)
	binary.Write(buf, binary.BigEndian, p.MsgData)
	return buf.Bytes()
}

func HeaderDecode(buf []byte, h *Header) error {
	if len(buf) != HeaderLen {
		return errors.New("buf len is not 16")
	}
	h.Length = binary.BigEndian.Uint32(buf[:4])
	h.Version = binary.BigEndian.Uint16(buf[4:6])
	h.Flag = binary.BigEndian.Uint16(buf[6:8])
	h.ServiceId = binary.BigEndian.Uint16(buf[8:10])
	h.CommandId = binary.BigEndian.Uint16(buf[10:12])
	h.SeqNum = binary.BigEndian.Uint16(buf[12:14])
	h.Reserved = binary.BigEndian.Uint16(buf[14:16])
	return nil
}

func Decode(buf []byte, p *ImPdu) error {
	if len(buf) < HeaderLen {
		log.Println("buf Length is invalid")
		return errors.New("buf Length is invalid")
	}
	var h Header
	err := HeaderDecode(buf[:16], &h)
	if err != nil {
		log.Println("dencode header failed")
		return err
	}
	if len(buf) < int(h.Length) {
		log.Println("msg is truncated!")
		return errors.New("msg body is truncated!")
	}

	p.Header = h
	p.MsgData = buf[16:h.Length]
	/*
		p = &ImPdu{
			Header:  h,
			MsgData: buf[16:h.Length],
		}
	*/

	return nil
}

func ImPduReader(reader io.Reader, pdu *ImPdu) error {
	fullBuf := bytes.NewBuffer([]byte{})
	for {
		data := make([]byte, 1024)
		cnt, err := reader.Read(data)
		if err != nil {
			log.Println(err)
			return err
		}
		fullBuf.Write(data[:cnt])
		if cnt < 1024 {
			break
		}
	}

	err := Decode(fullBuf.Bytes(), pdu)
	if err != nil {
		return err
	}
	return nil
}
