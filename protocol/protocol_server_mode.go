package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
)

type ServerModePacket struct {
	Tid  uint64
	Mode uint8
}

func (p *ServerModePacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_THREE_PHASE_MODE, p.Tid)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerMode(buffer []byte) *ServerModePacket {
	reader, _, _, tid := ParseHeader(buffer)
	mode, _ := reader.ReadByte()

	return &ServerModePacket{
		Tid:  tid,
		Mode: mode,
	}

}
