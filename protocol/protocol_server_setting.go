package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
	"time"
)

type ServerSettingPacket struct {
	Tid uint64
}

func (p *ServerSettingPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_SETTING, p.Tid)
	writer.WriteByte(byte(time.Now().Unix() % 3))
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerSetting(buffer []byte) *ServerSettingPacket {
	_, _, _, tid := ParseHeader(buffer)

	return &ServerSettingPacket{
		Tid: tid,
	}

}
