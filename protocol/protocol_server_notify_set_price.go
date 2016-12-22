package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
)

type ServerNotifyPricePacket struct {
	Tid    uint64
	Serial uint32
	Result uint8
}

func (p *ServerNotifyPricePacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_NOTIFY_SET_PRICE, p.Tid)
	base.WriteDWord(&writer, p.Serial)
	writer.WriteByte(p.Result)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerNotifyPrice(buffer []byte) *ServerNotifyPricePacket {
	reader, _, _, tid := ParseHeader(buffer)
	serial := base.ReadDWord(reader)

	return &ServerNotifyPricePacket{
		Tid:    tid,
		Serial: serial,
		Result: 0,
	}

}
