package conn

import (
	"github.com/giskook/charging_pile_client/protocol"
	"log"
)

func event_handler_server_msg_common(conn *Conn) {
	for conn.ReadMore {
		cmdid, pkglen := protocol.CheckProtocol(conn.RecieveBuffer)
		log.Printf("protocol id %d\n", cmdid)

		pkgbyte := make([]byte, pkglen)
		conn.RecieveBuffer.Read(pkgbyte)
		switch cmdid {
		case protocol.PROTOCOL_REP_LOGIN:
			p := protocol.ParseServerLogin(pkgbyte)
			event_handler_server_msg_login(conn, p)
			conn.ReadMore = true
		case protocol.PROTOCOL_REP_HEART:
			event_handler_server_msg_heart(conn)
			conn.ReadMore = true
		case protocol.PROTOCOL_REP_SETTING:
			p := protocol.ParseServerSetting(pkgbyte)
			event_handler_server_msg_setting(conn, p)
			conn.ReadMore = true
		case protocol.PROTOCOL_REP_PRICE:
			p := protocol.ParseServerPrice(pkgbyte)
			event_handler_server_msg_price(conn, p)
			conn.ReadMore = true
		case protocol.PROTOCOL_REP_THREE_PHASE_MODE:
			p := protocol.ParseServerMode(pkgbyte)
			event_handler_server_msg_mode(conn, p)
			conn.ReadMore = true
		case protocol.PROTOCOL_REQ_GUN_STATUS:
			p := protocol.ParseGetGunStatus(pkgbyte)
			event_handler_server_msg_get_gun_status(conn, p)
			conn.ReadMore = true
		case protocol.PROTOCOL_REQ_CHARGING:
			p := protocol.ParseServerCharging(pkgbyte)
			event_handler_server_msg_charging(conn, p)
			conn.ReadMore = true
		case protocol.PROTOCOL_REQ_STOP_CHARGING:
			p := protocol.ParseServerStopCharging(pkgbyte)
			event_handler_server_msg_stop_charging(conn, p)
			conn.ReadMore = true
		case protocol.PROTOCOL_REP_NOTIFY_SET_PRICE:
			p := protocol.ParseServerNotifyPrice(pkgbyte)
			event_handler_server_msg_notify_set_price(conn, p)
			conn.ReadMore = true
		case protocol.PROTOCOL_REP_CHARGING_COST:
			p := protocol.ParseChargingCost(pkgbyte)
			event_handler_server_msg_rep_charging_cost(conn, p)
		case protocol.PROTOCOL_REQ_PIN:
			p := protocol.ParseServerReqPin(pkgbyte)
			event_handler_server_msg_rep_pin(conn, p)

		case protocol.PROTOCOL_ILLEGAL:
			conn.ReadMore = false
		case protocol.PROTOCOL_HALF_PACK:
			conn.ReadMore = false
		}
	}
}
