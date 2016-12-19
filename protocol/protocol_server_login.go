package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
)

var PROTOCOL_SERVER_LOGIN_LEN uint16 = PROTOCOL_COMMON_LEN + 1

type ServerLoginPacket struct {
	Tid    uint64
	Result uint8
}

func (p *ServerLoginPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, PROTOCOL_SERVER_LOGIN_LEN,
		PROTOCOL_REP_LOGIN, p.Tid)
	writer.WriteByte(p.Result)
	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerLogin(buffer []byte) *ServerLoginPacket {
	reader, _, _, tid := ParseHeader(buffer)
	result, _ := reader.ReadByte()

	return &ServerLoginPacket{
		Tid:    tid,
		Result: result,
	}
}
