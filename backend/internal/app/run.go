package app

import (
	"strconv"

	"github.com/MillionEyes/web-traffic-workstation/backend/pkg/session"
	"github.com/google/gopacket"
)

func (c *Client) Run(packet gopacket.Packet) [][]string {
	actualSession, metrics, err := session.GetSessionMetrics(c.sessionStream, packet)
	if err != nil {
		return nil
	}

	c.sessionStream[*actualSession] = metrics
	sessionDetails := getSessionDetails(c.sessionStream)
	return transformSessionDetailsToMatrix(sessionDetails)
}

func getSessionDetails(sessionStream map[session.Session]*session.Metrics) []SessionDetails {
	var sessions []SessionDetails
	for currentSession, metric := range sessionStream {
		currentSessionDetails := SessionDetails{
			Session: currentSession,
			Metrics: *metric,
		}
		sessions = append(sessions, currentSessionDetails)
	}
	return sessions
}

func transformSessionDetailsToMatrix(sessions []SessionDetails) [][]string {
	var matrix [][]string
	matrix = append(
		matrix,
		[]string{"Client 1 IP", "Client 2 IP",
			"Client 1 Port", "Client 2 Port",
			"Protocol", "Total Packets Transferred",
			"Start Time", "End Time",
			"Data Transferred (in Bytes)",
		},
	)
	for _, currentSessionDetails := range sessions {
		matrix = append(
			matrix,
			[]string{currentSessionDetails.Session.Client1IP, currentSessionDetails.Session.Client2IP,
				currentSessionDetails.Session.Client1Port, currentSessionDetails.Session.Client2Port,
				currentSessionDetails.Session.Protocol, strconv.Itoa(currentSessionDetails.Metrics.TotalPackets),
				currentSessionDetails.Metrics.StartTime.String(), currentSessionDetails.Metrics.EndTime.String(),
				strconv.Itoa(currentSessionDetails.Metrics.TotalData),
			},
		)
	}
	return matrix
}
