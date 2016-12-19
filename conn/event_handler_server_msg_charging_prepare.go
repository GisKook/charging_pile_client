package conn

import (
	"github.com/giskook/charging_pile_client/pkg"
	"log"
)

func event_handler_server_msg_charging_prepare(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_charging")
	c.Send(p.Serialize())
}
