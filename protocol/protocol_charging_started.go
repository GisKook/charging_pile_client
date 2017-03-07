package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
)

type ChargingStartedPacket struct {
	Tid               uint64
	StartMeterReading uint32
	UserID            string
	StartTime         uint32
	PinCode           string
	TransactionID     string
	Amount            uint32
	Timestamp         string
}

func (p *ChargingStartedPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_CHARGING_STARTED, p.Tid)
	base.WriteDWord(&writer, p.StartMeterReading)
	base.WriteString(&writer, p.UserID)
	base.WriteDWord(&writer, p.StartTime)
	base.WriteString(&writer, p.PinCode)
	base.WriteBcdString(&writer, p.TransactionID)
	base.WriteDWord(&writer, p.Amount)
	base.WriteBcdTime(&writer)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}
