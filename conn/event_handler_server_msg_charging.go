package conn

import (
	"github.com/giskook/charging_pile_client/base"
	"github.com/giskook/charging_pile_client/pkg"
	"github.com/giskook/charging_pile_client/protocol"
	"log"
)

func event_handler_server_msg_charging(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_charging")
	c.Charging_Pile_Status <- base.CHARGING
	charging_pkg := p.(*protocol.ServerChargingPacket)
	c.Charging_Pile.Status = base.CHARGING
	c.Charging_Pile.UserID = charging_pkg.UserID
	c.Charging_Pile.TransactionID = charging_pkg.TransactionID
	c.Send(p.Serialize())
}
