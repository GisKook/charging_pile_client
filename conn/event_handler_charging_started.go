package conn

import (
	"github.com/giskook/charging_pile_client/protocol"
	"time"
)

func (c *Conn) SendChargingStarted() {
	charging_started := &protocol.ChargingStartedPacket{
		Tid:               c.ID,
		StartMeterReading: 0,
		UserID:            c.Charging_Pile.UserID,
		StartTime:         uint32(time.Now().Unix()),
		PinCode:           c.Charging_Pile.PinCode,
		TransactionID:     c.Charging_Pile.TransactionID,
		Amount:            c.Charging_Pile.Amount,
	}
	c.Send(charging_started.Serialize())
}
