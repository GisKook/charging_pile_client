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

type Conn struct {
	conn          *net.TCPConn
	config        *conf.Configuration
	RecieveBuffer *bytes.Buffer
	sendChan      chan pkg.Packet
	ticker        *time.Ticker
	readflag      int64
	writeflag     int64
	closeChan     chan bool
	index         uint32
	ID            uint64
	Charging_Pile *base.Charging_Pile
	ReadMore      bool
	Status        uint8
}

func NewConn(tid uint64, config *conf.Configuration) *Conn {
	return &Conn{
		RecieveBuffer: bytes.NewBuffer([]byte{}),
		config:        config,
		readflag:      time.Now().Unix(),
		writeflag:     time.Now().Unix(),
		ticker:        time.NewTicker(time.Duration(config.Client.HeartInterval) * time.Second),
		closeChan:     make(chan bool),
		index:         0,
		ReadMore:      true,
		Charging_Pile: &base.Charging_Pile{
			ID:     tid,
			Status: 0,
		},
		Status: ConnUnauth,
	}
}

func (c *Conn) Close() {
	c.closeChan <- true
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
			event_handler_server_msg_comon(c)
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
	c.conn.Write(data)
}

func (c *Conn) heart() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case <-c.ticker.C:
			heart := &protocol.ServerHeartPacket{
				Tid:    c.Charging_Pile.ID,
				Status: c.Charging_Pile.Status,
			}
			c.conn.Write(heart.Serialize())
		case <-c.closeChan:
			log.Println("recv close")
			return
		}
	}
}
