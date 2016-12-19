package conn

import (
	"github.com/giskook/charging_pile_client/pkg"
	"github.com/giskook/charging_pile_client/protocol"
	"log"
)

func event_handler_server_msg_mode(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_mode")
	packet := p.(*protocol.ServerModePacket)
	c.Charging_Pile.Mode = packet.Mode
}
