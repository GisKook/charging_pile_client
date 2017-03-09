package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
)

type ServerRepPinPacket struct {
	Tid     uint64
	Serial  uint32
	PinCode string
}

func (p *ServerRepPinPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_PIN, p.Tid)
	base.WriteDWord(&writer, p.Serial)
	base.WriteString(&writer, p.PinCode)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}
