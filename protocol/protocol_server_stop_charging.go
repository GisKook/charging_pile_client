package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
)

type ServerStopChargingPacket struct {
	Tid uint64
}

func (p *ServerStopChargingPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_STOP_CHARGING, p.Tid)
	writer.WriteByte(0)
	base.WriteBcdTime(&writer)

	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerStopCharging(buffer []byte) *ServerStopChargingPacket {
	_, _, _, tid := ParseHeader(buffer)

	return &ServerStopChargingPacket{
		Tid: tid,
	}
}
