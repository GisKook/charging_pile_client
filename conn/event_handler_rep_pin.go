package conn

import (
	"github.com/giskook/charging_pile_client/pkg"
	"github.com/giskook/charging_pile_client/protocol"
)

func event_handler_server_msg_rep_pin(c *Conn, p pkg.Packet) {
	req_pin := p.(*protocol.ServerReqPinPacket)
	c.Charging_Pile.Serial = req_pin.Serial

	rep_pin := &protocol.ServerRepPinPacket{
		Tid:     c.ID,
		Serial:  c.Charging_Pile.Serial,
		PinCode: c.Charging_Pile.PinCode,
	}
	c.Send(rep_pin.Serialize())
}
