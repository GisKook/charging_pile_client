package conn

import (
	"github.com/giskook/charging_pile_client/pkg"
	//"github.com/giskook/charging_pile_client/protocol"
	"log"
)

func event_handler_server_msg_stop_charging(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_stop_charging")
	//login_pkg := p.(*protocol.ServerStopChargingPacket)
	//c.Send(p.Serialize())
}
