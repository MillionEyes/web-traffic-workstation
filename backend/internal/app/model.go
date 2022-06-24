package app

import "github.com/MillionEyes/web-traffic-workstation/backend/pkg/session"

type SessionDetails struct {
	Session session.Session `json:"session"`
	Metrics session.Metrics `json:"metrics"`
}
