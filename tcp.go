package netlink

type TCPState uint32

// TCP States
const (
	TCPEstablished = iota + 0x01
	TCPSynSent
	TCPSynRecv
	TCPFinWait1
	TCPFinWait2
	TCPTimeWait
	TCPClose
	TCPCloseWait
	TCPLastAck
	TCPListen
	TCPClosing
	TCPNewSynRec
	TCPMaxStates
	TCPAll = 0xfff
)
