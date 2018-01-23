package conn

import (
	"github.com/giskook/charging_pile_client/protocol"
	"log"
	"time"
)

func (c *Conn) SendChargingStarted() {

	charging_started := &protocol.ChargingStartedPacket{
		Tid:               c.ID,
		StartMeterReading: METER_READING_INIT,
		UserID:            c.Charging_Pile.UserID,
		StartTime:         uint32(time.Now().Unix()),
		PinCode:           c.Charging_Pile.PinCode,
		TransactionID:     c.Charging_Pile.TransactionID,
		Amount:            c.Charging_Pile.Amount,
	}
	log.Println("---")
	log.Println(charging_started.StartTime)
	c.Send(charging_started.Serialize())
	c.Charging_Pile.MeterReading = METER_READING_INIT
}
