package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
	"time"
)

type ServerChargingPacket struct {
	Tid           uint64
	PinCode       string
	UserID        string
	TransactionID string
	Result        uint8
	Amount        uint32
}

func (p *ServerChargingPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_CHARGING, p.Tid)
	writer.WriteByte(p.Result)
	base.WriteString(&writer, p.PinCode)
	_time := time.Now().Format("060102150405")
	base.WriteBcdString(&writer, _time)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerCharging(buffer []byte, result uint8) *ServerChargingPacket {
	reader, _, _, tid := ParseHeader(buffer)
	userid := base.ReadString(reader, PROTOCOL_USERID_LEN)
	pincode := base.ReadString(reader, PROTOCOL_PINCODE_LEN)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	amount := base.ReadDWord(reader)

	return &ServerChargingPacket{
		Tid:           tid,
		PinCode:       pincode,
		UserID:        userid,
		TransactionID: transaction_id,
		Result:        result,
		Amount:        amount,
	}

}
