package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
	"time"
)

type ServerUploadMeterPacket struct {
	Tid              uint64
	UserID           string
	TransactionID    string
	MeterReading     uint32
	ChargingDuration uint32
	ChargingCapacity uint32
	ChargingPrice    uint32
	RealtimeA        uint32
	RealtimeV        uint32
}

func (p *ServerUploadMeterPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_UPLOAD_METER, p.Tid)
	writer.WriteByte(byte(len(p.UserID)))
	base.WriteString(&writer, p.UserID)
	base.WriteBcdString(&writer, p.TransactionID)
	base.WriteDWord(&writer, p.ChargingDuration)
	base.WriteDWord(&writer, p.ChargingCapacity)
	base.WriteDWord(&writer, p.ChargingPrice)
	base.WriteDWord(&writer, p.MeterReading)
	base.WriteDWord(&writer, p.RealtimeA)
	base.WriteDWord(&writer, p.RealtimeV)
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
