package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
	"time"
)

type ServerChargingPacket struct {
	Tid           uint64
	Serial        uint32
	PinCode       uint16
	UserID        string
	TransactionID string
	Amount        uint32
}

func (p *ServerChargingPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_CHARGING, p.Tid)
	base.WriteDWord(&writer, p.Serial)
	writer.WriteByte(byte(len(p.UserID)))
	base.WriteString(&writer, p.UserID)
	base.WriteBcdString(&writer, p.TransactionID)
	writer.WriteByte(byte(time.Now().Unix() % 4))
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseServerCharging(buffer []byte) *ServerChargingPacket {
	reader, _, _, tid := ParseHeader(buffer)
	serial := base.ReadDWord(reader)
	pincode := base.ReadWord(reader)
	userid_len, _ := reader.ReadByte()
	userid := base.ReadString(reader, userid_len)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	amount := base.ReadDWord(reader)

	return &ServerChargingPacket{
		Tid:           tid,
		Serial:        serial,
		PinCode:       pincode,
		UserID:        userid,
		TransactionID: transaction_id,
		Amount:        amount,
	}

}
