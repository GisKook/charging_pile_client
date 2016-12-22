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
		case protocol.PROTOCOL_REP_MODE:
			p := protocol.ParseServerMode(pkgbyte)
			event_handler_server_msg_mode(conn, p)
			conn.ReadMore = true
		case protocol.PROTOCOL_REP_MAX_CURRENT:
			p := protocol.ParseServerMaxCurrent(pkgbyte)
			event_handler_server_msg_max_current(conn, p)
			conn.ReadMore = true
		case protocol.PROTOCOL_REQ_CHARGING_PREPARE:
			p := protocol.ParseServerChargingPrepare(pkgbyte)
			event_handler_server_msg_charging_prepare(conn, p)
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

		case protocol.PROTOCOL_ILLEGAL:
			conn.ReadMore = false
		case protocol.PROTOCOL_HALF_PACK:
			conn.ReadMore = false
		}
	}
}
