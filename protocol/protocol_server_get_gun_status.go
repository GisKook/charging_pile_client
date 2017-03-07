package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
	//"time"
)

type ServerGetGunStatusPacket struct {
	Tid    uint64
	Status uint32
}

func (p *ServerGetGunStatusPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0, PROTOCOL_REP_GUN_STATUS, p.Tid)
	//writer.WriteByte(byte(time.Now().Unix() % 4))
	writer.WriteByte(0)
	base.WriteBcdTime(&writer)

	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseGetGunStatus(buffer []byte) *ServerGetGunStatusPacket {
	_, _, _, tid := ParseHeader(buffer)

	return &ServerGetGunStatusPacket{
		Tid: tid,
	}
}
