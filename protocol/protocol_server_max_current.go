package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
)

type ServerMaxCurrentPacket struct {
	Tid        uint64
	MaxCurrent uint8
}

func (p *ServerMaxCurrentPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_MAX_CURRENT, p.Tid)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerMaxCurrent(buffer []byte) *ServerMaxCurrentPacket {
	reader, _, _, tid := ParseHeader(buffer)
	max_current, _ := reader.ReadByte()

	return &ServerMaxCurrentPacket{
		Tid:        tid,
		MaxCurrent: max_current,
	}

}
