package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
)

type ServerNotifyPricePacket struct {
	Tid    uint64
	Result uint8
}

func (p *ServerNotifyPricePacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_CHARGING, p.Tid)
	writer.WriteByte(p.Result)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerNotifyPrice(buffer []byte) *ServerNotifyPricePacket {
	_, _, _, tid := ParseHeader(buffer)

	return &ServerNotifyPricePacket{
		Tid:    tid,
		Result: 0,
	}

}
