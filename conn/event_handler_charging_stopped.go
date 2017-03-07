package conn

import (
	"github.com/giskook/charging_pile_client/protocol"
	"time"
)

func (c *Conn) SendChargingStopped() {
	charging_stopped := &protocol.ChargingStoppedPacket{
		Tid:              c.ID,
		StopReason:       0,
		EndMeterReadging: c.Charging_Pile.MeterReading + 5,
		UserID:           c.Charging_Pile.UserID,
		StopTime:         uint32(time.Now().Unix()),
		TransactionID:    c.Charging_Pile.TransactionID,
	}
	c.Send(charging_stopped.Serialize())
}
