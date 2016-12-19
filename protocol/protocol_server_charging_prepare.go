package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
	"time"
)

const (
	PROTOCOL_REP_CHARGING_PREPARE_LEN uint16 = PROTOCOL_COMMON_LEN + 7
)

type ServerChargingPreparePacket struct {
	Tid     uint64
	Serial  uint32
	PinCode uint16
}

func (p *ServerChargingPreparePacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, PROTOCOL_REP_CHARGING_PREPARE_LEN, PROTOCOL_REP_CHARGING, p.Tid)
	base.WriteDWord(&writer, p.Serial)
	writer.WriteByte(byte(time.Now().Unix() % 4))

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerChargingPrepare(buffer []byte) *ServerChargingPreparePacket {
	reader, _, _, tid := ParseHeader(buffer)
	serial := base.ReadDWord(reader)
	pincode := base.ReadWord(reader)

	return &ServerChargingPreparePacket{
		Tid:     tid,
		Serial:  serial,
		PinCode: pincode,
	}

}
