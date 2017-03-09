package base

type Price struct {
	Start_hour      uint8
	Start_min       uint8
	End_hour        uint8
	End_min         uint8
	Elec_unit_price uint16
	Service_price   uint16
}

const (
	IDLE               uint8 = 0
	CHARGING           uint8 = 1
	TOBECHARGING       uint8 = 2
	FULL               uint8 = 3
	MAINTAINCE         uint8 = 4
	TOBE_STOP_CHARGING uint8 = 5
	CHARGING_STOPPED   uint8 = 6
)

type Charging_Pile struct {
	ID              uint64
	BoxVersion      byte
	ProtocolVersion byte
	MaxCurrent      uint8
	Mode            uint8
	Prices          []Price
	Status          uint8

	ConnectWay uint8
	Wifi       string
	Passwd     string
	Interface  uint8

	UserID            string
	TransactionID     string
	StartMeterReading uint32
	ChargingCapacity  uint32
	ChargingPrice     uint32
	MeterReading      uint32
	Power             uint16
	Va                uint16
	Vb                uint16
	Vc                uint16
	Ia                uint16
	Ib                uint16
	Ic                uint16
	RealtimeA         uint32
	RealtimeV         uint32

	Serial  uint32
	PinCode string
	Amount  uint32
}
