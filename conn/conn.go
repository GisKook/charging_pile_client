package conn

import (
	"bytes"
	"github.com/giskook/charging_pile_client/base"
	"github.com/giskook/charging_pile_client/conf"
	"github.com/giskook/charging_pile_client/pkg"
	"github.com/giskook/charging_pile_client/protocol"
	"log"
	"net"
	"time"
)

var ConnSuccess uint8 = 0
var ConnUnauth uint8 = 1

const (
	METER_READING_INIT uint32 = 3000
	METER_READING_STEP uint32 = 1
)

type Conn struct {
	conn                 *net.TCPConn
	config               *conf.Configuration
	RecieveBuffer        *bytes.Buffer
	sendChan             chan pkg.Packet
	ticker               *time.Ticker
	readflag             int64
	writeflag            int64
	closeChan            chan bool
	ID                   uint64
	Charging_Pile        *base.Charging_Pile
	Charging_Pile_Status chan uint8
	ReadMore             bool
	Status               uint8
}

func NewConn(tid uint64, config *conf.Configuration) *Conn {
	return &Conn{
		RecieveBuffer: bytes.NewBuffer([]byte{}),
		config:        config,
		readflag:      time.Now().Unix(),
		writeflag:     time.Now().Unix(),
		ticker:        time.NewTicker(time.Duration(config.Client.HeartInterval) * time.Second),
		closeChan:     make(chan bool),
		ReadMore:      true,
		ID:            tid,
		Charging_Pile: &base.Charging_Pile{
			ID:               tid,
			Status:           base.IDLE,
			ChargingCapacity: 0,
			MeterReading:     METER_READING_INIT,
			RealtimeA:        0,
			RealtimeV:        0,
		},
		Charging_Pile_Status: make(chan uint8),
		Status:               ConnUnauth,
	}
}

func (c *Conn) Close() {
	c.closeChan <- true
	close(c.Charging_Pile_Status)
	c.ticker.Stop()
	c.RecieveBuffer.Reset()
	c.conn.Close()
	close(c.closeChan)
}

func (c *Conn) GetBuffer() *bytes.Buffer {
	return c.RecieveBuffer
}

func (c *Conn) Start() {
	defer func() {
		c.Close()
	}()
	tcpaddr, err := net.ResolveTCPAddr("tcp", c.config.Server.Addr)

	c.conn, err = net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		log.Println(err.Error())
		return
	}

	login := &protocol.LoginPacket{
		Tid: c.Charging_Pile.ID,
	}

	log.Println("send login")
	_, err = c.conn.Write(login.Serialize())
	if err != nil {
		log.Println(err.Error())
	}

	//log.Println("send req price")
	//req_price := &protocol.ServerPricePacket{
	//	Tid: c.Charging_Pile.ID,
	//}
	//c.Send(req_price.Serialize())

	//	log.Println("send setting")
	//	setting := &protocol.ServerSettingPacket{
	//		Tid: c.Charging_Pile.ID,
	//	}
	//
	//	_, err = c.conn.Write(setting.Serialize())
	//	if err != nil {
	//		log.Println(err.Error())
	//	}

	//	heart := &protocol.ServerHeartPacket{
	//		Tid:    c.Charging_Pile.ID,
	//		Status: 0,
	//	}
	//	_, err = c.conn.Write(heart.Serialize())
	//	if err != nil {
	//		log.Println(err.Error())
	//	}

	c.handle()
}

func (c *Conn) handle() {
	defer func() {
		c.conn.Close()
	}()

	for {
		buffer := make([]byte, 1024)
		buf_len, err := c.conn.Read(buffer)
		if err != nil {
			log.Println(err)
		}
		c.RecieveBuffer.Write(buffer[0:buf_len])
		if buf_len > 0 {
			log.Printf("<IN> %x\n", buffer[0:buf_len])
			c.ReadMore = true
			event_handler_server_msg_common(c)
		}
	}
}

func (c *Conn) UpdateReadflag() {
	c.readflag = time.Now().Unix()
}

func (c *Conn) UpdateWriteflag() {
	c.writeflag = time.Now().Unix()
}

func (c *Conn) Send(data []byte) {
	log.Printf("<OUT> %x\n", data)
	c.conn.Write(data)
}

func (c *Conn) heart() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case status := <-c.Charging_Pile_Status:
			c.Charging_Pile.Status = status
			if status == base.FULL {
				c.ProccessChargingPileStopChargingStatus()
				c.Charging_Pile.Status = base.IDLE
			} else if status == base.TOBECHARGING {
				c.SendChargingStarted()
				c.Charging_Pile.Status = base.CHARGING
			} else if status == base.TOBE_STOP_CHARGING {
				c.SendChargingStopped()
				c.Charging_Pile.Status = base.IDLE
			}
		case <-c.ticker.C:
			if c.Charging_Pile.Status == base.IDLE {
				c.ProccessChargingPileIDLEStatus()
			} else if c.Charging_Pile.Status == base.CHARGING {
				c.ProccessChargingPileChargingStatus()
			}

		case <-c.closeChan:
			log.Println("recv close")
			return
		}
	}
}

func (c *Conn) ProccessChargingPileIDLEStatus() {
	heart := &protocol.ServerHeartPacket{
		Tid:    c.Charging_Pile.ID,
		Status: c.Charging_Pile.Status,
	}
	c.Send(heart.Serialize())
}

func (c *Conn) ProccessChargingPileChargingStatus() {
	time.Sleep(2 * time.Second)
	upload_meter := &protocol.ServerUploadMeterPacket{
		Tid:          c.ID,
		MeterReading: c.Charging_Pile.MeterReading + METER_READING_STEP,
		Power:        200,
		Status:       0,
		Va:           uint16(time.Now().Unix()) % 380,
		Vb:           uint16(time.Now().Unix()) % 380,
		Vc:           uint16(time.Now().Unix()) % 380,
		Ia:           uint16(time.Now().Unix()) % 20,
		Ib:           uint16(time.Now().Unix()) % 20,
		Ic:           uint16(time.Now().Unix()) % 20,
	}
	// for test
	if c.Charging_Pile.MeterReading >= 3005 && c.Charging_Pile.MeterReading < 3007 {
		upload_meter = &protocol.ServerUploadMeterPacket{
			Tid:          c.ID,
			MeterReading: 0,
			Power:        200,
			Status:       0,
			Va:           uint16(time.Now().Unix()) % 380,
			Vb:           uint16(time.Now().Unix()) % 380,
			Vc:           uint16(time.Now().Unix()) % 380,
			Ia:           uint16(time.Now().Unix()) % 20,
			Ib:           uint16(time.Now().Unix()) % 20,
			Ic:           uint16(time.Now().Unix()) % 20,
		}
	}
	// for test
	c.Send(upload_meter.Serialize())
	c.Charging_Pile.MeterReading += METER_READING_STEP
}

func (c *Conn) ProccessChargingPileStopChargingStatus() {
	stop_charging := &protocol.ServerStopChargingPacket{
		Tid: c.ID,
	}
	c.Send(stop_charging.Serialize())
	c.Charging_Pile.MeterReading += METER_READING_STEP
}
