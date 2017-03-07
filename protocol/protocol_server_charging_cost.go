package protocol

import (
	"github.com/giskook/charging_pile_client/base"
)

type ServerChargingCostPacket struct {
	Tid  uint64
	Cost uint32
}

func (p *ServerChargingCostPacket) Serialize() []byte {
	//	var writer bytes.Buffer
	//	WriteHeader(&writer, PROTOCOL_REP_CHARGING_COST,
	//		0, p.Tid)
	//		base.Write
	//	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	//	writer.WriteByte(PROTOCOL_END_FLAG)
	//
	//	return writer.Bytes()

	return nil
}

func ParseChargingCost(buffer []byte) *ServerChargingCostPacket {
	reader, _, _, tid := ParseHeader(buffer)
	cost := base.ReadDWord(reader)

	return &ServerChargingCostPacket{
		Tid:  tid,
		Cost: cost,
	}
}
