package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
)

type ServerUploadMeterPacket struct {
	Tid          uint64
	MeterReading uint32
	Power        uint16
	Status       uint8
	Va           uint16
	Vb           uint16
	Vc           uint16
	Ia           uint16
	Ib           uint16
	Ic           uint16
}

func (p *ServerUploadMeterPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_CHARGING_DATA_UPLOAD, p.Tid)
	base.WriteDWord(&writer, p.MeterReading)
	base.WriteWord(&writer, p.Power)
	writer.WriteByte(p.Status)
	base.WriteWord(&writer, p.Va)
	base.WriteWord(&writer, p.Vb)
	base.WriteWord(&writer, p.Vc)
	base.WriteWord(&writer, p.Ia)
	base.WriteWord(&writer, p.Ib)
	base.WriteWord(&writer, p.Ic)
	base.WriteBcdTime(&writer)

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
