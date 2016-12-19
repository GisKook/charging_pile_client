package conn

import (
	"github.com/giskook/charging_pile_client/pkg"
	//	"github.com/giskook/charging_pile_client/protocol"
	"log"
)

func event_handler_server_msg_setting(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_setting")
	//packet := p.(*protocol.ServerSettingPacket)
	//log.Println(c)
	//log.Println(p)
}
