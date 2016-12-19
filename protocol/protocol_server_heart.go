package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
	"time"
)

type ServerHeartPacket struct {
	Tid    uint64
	Status uint8
}

func (p *ServerHeartPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_HEART, p.Tid)
	writer.WriteByte(p.Status)
	_time := time.Now().Format("20060102150405")
	base.WriteBcdString(&writer, _time)

	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

//func ParseServerHeart(buffer []byte) *ServerHeartPacket {
//	reader, _, _, tid := ParseHeader(buffer)
//	serial := base.ReadDWord(reader)
//	pincode := base.ReadWord(reader)
//	userid_len := reader.ReadByte()
//	userid := base.ReadString(reader, userid_len)
//	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
//	amount := base.ReadDWord(reader)
//
//	return &ServerHeartPacket{
//		Tid:           tid,
//		Serial:        serial,
//		PinCode:       pincode,
//		UserID:        userid,
//		TransactionID: transaction_id,
//		Amount:        amount,
//	}
//
//}
