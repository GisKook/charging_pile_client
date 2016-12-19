package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
)

type ServerPricePacket struct {
	Tid uint64
}

func (p *ServerPricePacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_PRICE, p.Tid)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerPrice(buffer []byte) *ServerPricePacket {

	_, _, _, tid := ParseHeader(buffer)
	return &ServerPricePacket{
		Tid: tid,
	}

}
