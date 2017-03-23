package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
	"time"
)

type LoginPacket struct {
	Tid               uint64
	ProtocolVersion   uint8
	HardwareVersion   uint8
	PinCode           string
	Status            uint8
	UserID            string
	TransactionID     string
	StartTime         uint32
	StartMeterReading uint32
}

func (p *LoginPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_LOGIN, p.Tid)
	writer.WriteByte(p.ProtocolVersion)
	writer.WriteByte(p.HardwareVersion)
	//	base.WriteString(&writer, p.PinCode)
	base.WriteString(&writer, "12")
	writer.WriteByte(p.Status)
	base.WriteString(&writer, "0000000000000000")
	base.WriteString(&writer, "000000000000000")
	//	base.WriteString(&writer, p.UserID)
	//	base.WriteString(&writer, p.TransactionID)
	_cur_t := time.Now()
	base.WriteDWord(&writer, uint32(_cur_t.Unix()))
	base.WriteDWord(&writer, p.StartMeterReading)
	_time := _cur_t.Format("060102150405")
	base.WriteBcdString(&writer, _time)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}
