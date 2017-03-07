package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
)

type ChargingStoppedPacket struct {
	Tid              uint64
	StopReason       uint8
	EndMeterReadging uint32
	UserID           string
	StopTime         uint32
	TransactionID    string
}

func (p *ChargingStoppedPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_CHARGING_STOPPED, p.Tid)
	writer.WriteByte(p.StopReason)
	base.WriteDWord(&writer, p.EndMeterReadging)
	base.WriteString(&writer, p.UserID)
	base.WriteDWord(&writer, p.StopTime)
	base.WriteBcdString(&writer, p.TransactionID)
	base.WriteBcdTime(&writer)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}
