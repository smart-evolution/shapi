package sapi

import (
	"fmt"
	aca "github.com/smart-evolution/agents-cmd-api"
	"golang.org/x/net/websocket"
	"log"
	"net"
)

type message struct {
	ID   string `json:"id"`
	Left int    `json:"left"`
	Top  int    `json:"top"`
}

var conn net.Conn
var err error
var devType string
var device = "192.168.2.166:81"

func connect() {
	if conn == nil {
		conn, err = net.Dial("tcp", device)

		if err != nil {
			fmt.Println("error connecting device " + device)
			return
		}

		_, err = conn.Write([]byte("CMDWHO"))

		if err != nil {
			fmt.Println("error getting device type")
			return
		}

		buff := make([]byte, 512)
		n, err := conn.Read(buff)

		if err != nil {
			fmt.Println("error retrieving device type")
			fmt.Println(err)
			return
		}

		devType = string(buff[:n])

		if _, ok := aca.ApiMap[devType]; !ok {
			fmt.Println("unknown device type '" + devType + "'")
			return
		}

		fmt.Println("connected to device type '" + devType + "'")
	}
}

var prevCmd string

func move(m message) {
	var cmd string

	if m.Top < 15 {
		cmd = "w"
	} else if m.Top > 40 {
		cmd = "x"
	} else if m.Left < 15 {
		cmd = "a"
	} else if m.Left > 40 {
		cmd = "d"
	} else {
		cmd = "s"
	}

	if cmd != prevCmd {
		prevCmd = cmd
		apiVersion := aca.ApiMap[devType]
		hardwareComms := aca.Comms[apiVersion][cmd]

		for _, c := range hardwareComms {
			_, err := conn.Write([]byte(c))
			if err != nil {
				fmt.Println("RES: sending command failed " + c)
				break
			}
		}
	}
}

// AgentStreaming - handle agent streaming websocket connection
func AgentStreaming(ws *websocket.Conn) {
	go connect()

	for {
		var m message

		if err := websocket.JSON.Receive(ws, &m); err != nil {
			log.Println(err)
			break
		}

		move(m)
	}
}
