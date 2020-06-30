package netlink

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/vishvananda/netlink/nl"
)

type TCPInfo struct {
	State                uint8
	CaState              uint8
	Retransmits          uint8
	Probes               uint8
	Backoff              uint8
	Options              uint8
	Scale                uint8
	RateLimitAndFastOpen uint8
	Rto                  uint32
	Ato                  uint32
	SndMss               uint32
	RcvMss               uint32
	Unacked              uint32
	Sacked               uint32
	Lost                 uint32
	Retrans              uint32
	Fackets              uint32
	LastDataSent         uint32
	LastAckSent          uint32
	LastDataRecv         uint32
	LastAckRecv          uint32
	Pmtu                 uint32
	RcvSsThresh          uint32
	Rtt                  uint32
	Rttvar               uint32
	SndSsthresh          uint32
	SndCwnd              uint32
	Advmss               uint32
	Reordering           uint32
	RcvRtt               uint32
	RcvSpace             uint32
	TotalRetrans         uint32
	PacingRate           uint64
	MaxPacingRate        uint64
	BytesAcked           uint64 /* RFC4898 tcpEStatsAppHCThruOctetsAcked */
	BytesReceived        uint64 /* RFC4898 tcpEStatsAppHCThruOctetsReceived */
	SegsOut              uint32 /* RFC4898 tcpEStatsPerfSegsOut */
	SegsIn               uint32 /* RFC4898 tcpEStatsPerfSegsIn */
	NotsentBytes         uint32
	MinTtt               uint32
	DataSegsIn           uint32 /* RFC4898 tcpEStatsDataSegsIn */
	DataSegsOut          uint32 /* RFC4898 tcpEStatsDataSegsOut */
	DeliveryRate         uint64
	BusyTime             uint64 /* Time (usec) busy sending data */
	RwndLimited          uint64 /* Time (usec) limited by receive window */
	SndbufLimited        uint64 /* Time (usec) limited by send buffer */
	Delivered            uint32
	DeliveredCe          uint32
	BytesSent            uint64 /* RFC4898 tcpEStatsPerfHCDataOctetsOut */
	BytesRetrans         uint64 /* RFC4898 tcpEStatsPerfOctetsRetrans */
	DsackDups            uint32 /* RFC4898 tcpEStatsStackDSACKDups */
	ReordSeen            uint32 /* reordering events seen */
	RcvOoopack           uint32 /* Out-of-order packets received */
	SndWnd               uint32 /* peer's advertised receive window after * scaling (bytes) */
}

func (t *TCPInfo) SndWscale() uint8 {
	return t.Scale >> 4 // first 4 bits
}

func (t *TCPInfo) RcvWscale() uint8 {
	return t.Scale & 0xf // last 4 bits
}

func (t *TCPInfo) DeliveryRateAppLimited() uint8 {
	return t.RateLimitAndFastOpen >> 7 // get first bit
}

func (t *TCPInfo) FastopenClientFail() uint8 {
	return t.RateLimitAndFastOpen >> 5 & 3 // get next two bits
}

func (t *TCPInfo) deserialize(b []byte) error {
	rb := bytes.NewBuffer(b)

	err := binary.Read(rb, nl.NativeEndian(), t)
	if err == io.EOF {
		return nil
	} else if err != nil {
		return err
	}

	return nil
}
