package sapi

import (
	aca "github.com/smart-evolution/agents-cmd-api"
	"github.com/smart-evolution/shapi/constants"
	"github.com/smart-evolution/shapi/utils"
	"golang.org/x/net/websocket"
	"net"
)

type message struct {
	ID   string `json:"id"`
	Left int    `json:"left"`
	Top  int    `json:"top"`
	IP   string `json:"ip"`
	Flag string `json:"flag"`
}

var (
	conn    net.Conn
	err     error
	devType string
)

func connect(ws *websocket.Conn, device string) {
	utils.Log("connecting to device with ip '" + device + "'")
	conn, err = net.Dial("tcp", device+":"+constants.AgentTCPPort)

	if err != nil {
		utils.Log("error connecting device " + device)
		websocket.JSON.Send(ws, `{"type":"error","message":"Error connecting to device `+device+`"}`)
		return
	}

	_, err = conn.Write([]byte("CMDWHO"))

	if err != nil {
		utils.Log("error getting device type")
		websocket.JSON.Send(ws, `{"type":"error","message":"Error getting device type"}`)
		return
	}

	buff := make([]byte, 512)
	n, err := conn.Read(buff)

	if err != nil {
		utils.Log("error retrieving device type")
		websocket.JSON.Send(ws, `{"type":"error","message":"Error command failed"}`)
		return
	}

	devType = string(buff[:n])

	if _, ok := aca.APIMap[devType]; !ok {
		utils.Log("unknown device type '" + devType + "'")
		websocket.JSON.Send(ws, `{"type":"error","message":"Unknown device type '`+devType+`'"}`)
		return
	}

	utils.Log("connected to device type '" + devType + "'")

	websocket.JSON.Send(ws, `{"type":"connected","message":"Connected to the device '`+device+`'"}`)
}

var prevCmd string

func move(ws *websocket.Conn, m message) {
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
		apiVersion := aca.APIMap[devType]
		hardwareComms := aca.Comms[apiVersion][cmd]

		for _, c := range hardwareComms {
			_, err := conn.Write([]byte(c))
			if err != nil {
				websocket.JSON.Send(ws, "{\"type\":\"error\",\"message\":\"Sending command failed\"}")
				utils.Log("RES: sending command failed " + c)
				break
			}
		}
	}
}

// AgentStreaming - handle agent streaming websocket connection
func AgentStreaming(ws *websocket.Conn) {
	var m message

	for {
		err := websocket.JSON.Receive(ws, &m)

		if err == nil {
			go connect(ws, m.IP)
			break
		}
	}

	for {
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			utils.Log(err)
			ws.Close()
			conn.Close()
			return
		}

		if m.Flag == "disconnect" {
			_, err := conn.Write([]byte("CMDDIS"))
			if err != nil {
				websocket.JSON.Send(ws, "{\"type\":\"disconnect\",\"message\":\"Disconnecting failed\"}")
				utils.Log("RES: sending command failed CMDDIS")
				return
			}

			websocket.JSON.Send(ws, "{\"type\":\"disconnect\",\"message\":\"Device disconnected\"}")

			ws.Close()
			conn.Close()
		} else {
			move(ws, m)
		}
	}
}
