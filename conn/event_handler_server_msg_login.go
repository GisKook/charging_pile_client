package conn

import (
	"github.com/giskook/charging_pile_client/pkg"
	"github.com/giskook/charging_pile_client/protocol"
	"log"
)

func event_handler_server_msg_login(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_login")
	login_pkg := p.(*protocol.ServerLoginPacket)
	c.Status = login_pkg.Result
	if c.Status == 0 {
		go c.heart()

		//		log.Println("send req price")
		//		req_price := &protocol.ServerPricePacket{
		//			Tid: c.Charging_Pile.ID,
		//		}
		//		c.Send(req_price.Serialize())

		//		log.Println("send max current")
		//		max_current := &protocol.ServerMaxCurrentPacket{
		//			Tid: c.Charging_Pile.ID,
		//		}
		//		c.Send(max_current.Serialize())
	}
}
