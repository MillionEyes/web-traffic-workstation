package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/MillionEyes/web-traffic-workstation/backend/internal/app"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"

	"golang.org/x/net/websocket"
)

var (
	device       string = "en0"
	snapshot_len int32  = 1024
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = 30
	handle       pcap.Handle
)

func Echo(ws *websocket.Conn) {
	var err error

	handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	appClient, _ := app.NewClient(handle)

	//Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		reply := appClient.Run(packet)
		msg := ""
		if len(reply) != 0 {
			for _, r := range reply[1:] {
				key := `"` + strings.Join(r[:5], " ") + `"`
				val := `"` + strings.Join(r[5:], " ") + `"`

				currentMsg := "{" + `"key":` + key + `, "val": ` + val + "},"
				msg += currentMsg
			}
		}

		if msg != "" {
			msg = "[" + msg[:len(msg)-1] + "]"
		} else {
			msg = "[]"
		}
		// fmt.Println("Sending to client: ")

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			continue
		}
	}
}

func main() {
	http.Handle("/socket", websocket.Handler(Echo))
	log.Println("serving")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
