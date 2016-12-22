package conn

import (
	"github.com/giskook/charging_pile_client/base"
	"github.com/giskook/charging_pile_client/pkg"
	"github.com/giskook/charging_pile_client/protocol"
	"log"
)

func event_handler_server_msg_stop_charging(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_stop_charging")
	stop_charging_pkg := p.(*protocol.ServerStopChargingPacket)
	c.Charging_Pile.UserID = stop_charging_pkg.UserID
	c.Charging_Pile.TransactionID = stop_charging_pkg.TransactionID
	c.Charging_Pile_Status <- base.FULL
	//login_pkg := p.(*protocol.ServerStopChargingPacket)
	//c.Send(p.Serialize())
}
