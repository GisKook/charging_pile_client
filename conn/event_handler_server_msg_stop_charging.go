package conn

import (
	"github.com/giskook/charging_pile_client/base"
	"github.com/giskook/charging_pile_client/pkg"
	"log"
)

func event_handler_server_msg_stop_charging(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_stop_charging")
	c.Charging_Pile_Status <- base.TOBE_STOP_CHARGING
	//login_pkg := p.(*protocol.ServerStopChargingPacket)
	//c.Send(p.Serialize())
}
