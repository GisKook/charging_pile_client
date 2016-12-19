package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
	"time"
)

type ServerStopChargingPacket struct {
	Tid           uint64
	Serial        uint32
	UserID        string
	TransactionID string
	StopReason    uint8
}

func (p *ServerStopChargingPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_STOP_CHARGING, p.Tid)
	base.WriteDWord(&writer, p.Serial)
	writer.WriteByte(byte(len(p.UserID)))
	base.WriteString(&writer, p.UserID)
	base.WriteBcdString(&writer, p.TransactionID)
	writer.WriteByte(p.StopReason)
	base.WriteDWord(&writer, 1024)
	base.WriteDWord(&writer, 1025)
	base.WriteDWord(&writer, 50)
	base.WriteDWord(&writer, 500)
	_time := time.Now().Format("20060102150405")
	base.WriteBcdString(&writer, _time)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerStopCharging(buffer []byte) *ServerStopChargingPacket {
	reader, _, _, tid := ParseHeader(buffer)
	serial := base.ReadDWord(reader)
	userid_len, _ := reader.ReadByte()
	userid := base.ReadString(reader, userid_len)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)

	return &ServerStopChargingPacket{
		Tid:           tid,
		Serial:        serial,
		UserID:        userid,
		TransactionID: transaction_id,
		StopReason:    1,
	}
}
