package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
	"time"
)

type LoginPacket struct {
	Tid uint64
}

func (p *LoginPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_LOGIN, p.Tid)
	writer.WriteByte(1)
	writer.WriteByte(1)
	writer.WriteByte(0)
	_time := time.Now().Format("20060102150405")
	base.WriteBcdString(&writer, _time)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}
