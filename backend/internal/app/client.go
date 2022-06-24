package app

import (
	"github.com/MillionEyes/web-traffic-workstation/backend/pkg/session"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Client struct {
	packetSource *gopacket.PacketSource

	sessionStream map[session.Session]*session.Metrics
}

func NewClient(handle *pcap.Handle) (*Client, error) {
	return &Client{
		packetSource:  gopacket.NewPacketSource(handle, handle.LinkType()),
		sessionStream: map[session.Session]*session.Metrics{},
	}, nil
}
