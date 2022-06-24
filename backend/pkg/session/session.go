package session

import (
	"errors"
	"time"

	"github.com/google/gopacket"
)

type Session struct {
	Client1IP   string
	Client2IP   string
	Client1Port string
	Client2Port string
	Protocol    string
}

type Metrics struct {
	TotalPackets int
	TotalData    int
	StartTime    time.Time
	EndTime      time.Time
}

func NewSession(packet gopacket.Packet) (*Session, error) {
	if packet.TransportLayer() == nil {
		return nil, errors.New("unknown protocol")
	}
	return &Session{
		Client1IP:   packet.NetworkLayer().NetworkFlow().Src().String(),
		Client2IP:   packet.NetworkLayer().NetworkFlow().Dst().String(),
		Client1Port: packet.TransportLayer().TransportFlow().Src().String(),
		Client2Port: packet.TransportLayer().TransportFlow().Dst().String(),
		Protocol:    packet.TransportLayer().LayerType().String(),
	}, nil
}

func GetSessionMetrics(sessionStream map[Session]*Metrics, packet gopacket.Packet) (*Session, *Metrics, error) {
	currentSession, err := NewSession(packet)
	if err != nil {
		return nil, nil, err
	}
	sessionMetrics, present := sessionStream[*currentSession]
	if present {
		sessionMetrics.EndTime = packet.Metadata().Timestamp
		sessionMetrics.TotalPackets += 1
		sessionMetrics.TotalData += packet.Metadata().Length
		return currentSession, sessionMetrics, nil
	}
	mateSession := &Session{
		Client1IP:   currentSession.Client2IP,
		Client2IP:   currentSession.Client1IP,
		Client1Port: currentSession.Client2Port,
		Client2Port: currentSession.Client1Port,
		Protocol:    currentSession.Protocol,
	}
	sessionMetrics, present = sessionStream[*mateSession]
	if present {
		sessionMetrics.EndTime = packet.Metadata().Timestamp
		sessionMetrics.TotalPackets += 1
		sessionMetrics.TotalData += packet.Metadata().Length
		return mateSession, sessionMetrics, nil
	} else {
		return currentSession, &Metrics{
			TotalPackets: 1,
			StartTime:    packet.Metadata().Timestamp,
			EndTime:      packet.Metadata().Timestamp,
			TotalData:    packet.Metadata().Length,
		}, nil
	}
}
