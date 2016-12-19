package conn

import (
	"github.com/giskook/charging_pile_client/pkg"
	//"github.com/giskook/charging_pile_client/protocol"
)

func event_handler_server_msg_price(c *Conn, p pkg.Packet) {
	//packet := p.(*protocol.ServerPricePacket)
	//c.Charging_Pile.Prices = p.Prices
}
