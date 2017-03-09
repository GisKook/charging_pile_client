package protocol

import (
	"github.com/giskook/charging_pile_client/base"
)

type ServerReqPinPacket struct {
	Tid    uint64
	Serial uint32
}

func (p *ServerReqPinPacket) Serialize() []byte {
	return nil
}

func ParseServerReqPin(buffer []byte) *ServerReqPinPacket {
	reader, _, _, tid := ParseHeader(buffer)
	serial := base.ReadDWord(reader)

	return &ServerReqPinPacket{
		Tid:    tid,
		Serial: serial,
	}
}
