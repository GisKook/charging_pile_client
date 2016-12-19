package conn

import (
	"github.com/giskook/charging_pile_client/pkg"
	"github.com/giskook/charging_pile_client/protocol"
	"log"
)

func event_handler_server_msg_max_current(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_max_current")
	packet := p.(*protocol.ServerMaxCurrentPacket)

	c.Charging_Pile.MaxCurrent = packet.MaxCurrent
}
