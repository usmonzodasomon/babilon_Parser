package models

type AppSettings struct {
	PostgresSettings
	Flags
}

type PostgresSettings struct {
	User     string
	Password string
	Server   string
	Port     string
	Database string
	SSLMode  string
}

type Flags struct {
	AccountID     int64
	Tclass        int64
	SourceIP      string
	DestinationIP string
}

type ParserData struct {
	DeviceId      uint8
	SourceIP      uint32
	DestinationIP uint32
	NexthopIP     uint32
	Iface         uint16
	Oface         uint16
	Packets       uint32
	Bytes         uint32
	StartTime     uint32
	EndTime       uint32
	Sport         uint16
	Dport         uint16
	TcpFlags      uint8
	Proto         uint8
	Tos           uint8
	SrcAS         uint32
	DstAS         uint32
	SrcMask       uint8
	DstMask       uint8
	SlinkID       uint32
	AccountID     uint32
	BillingIP     uint32
	Tclass        uint32
	DateTime      uint32
	NfSourceIP    uint32
}

type DBData struct {
	SourceIP      string
	DestinationIP string
	Packets       uint32
	Bytes         uint32
	Sport         uint16
	Dport         uint16
	Proto         uint8
	AccountID     uint32
	Tclass        uint32
	DateTime      uint32
	NfSourceIP    string
}
