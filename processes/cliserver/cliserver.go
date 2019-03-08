package cliserver

import (
	"encoding/json"
	"github.com/smart-evolution/smarthome/utils"
	"net"
)

func handleRequest(conn net.Conn) {
	utils.Log("received cli message")

	buff := make([]byte, 512)
	n, err := conn.Read(buff)

	if err != nil {
		utils.Log("error reading cli message")
		return
	}

	msg := make(map[string]interface{})
	json.Unmarshal(buff[:n], &msg)

	cmd := msg["cmd"]

	if cmd == "status" {
		conn.Write([]byte("application up and running"))
	} else {
		utils.Log("invalid cli command")
	}

	conn.Close()
}

// RunService - start CLIserver service
func RunService(port string) {
	l, err := net.Listen("tcp", ":"+port)

	if err != nil {
		utils.Log("failed to setup cli tcp server")
	}

	defer l.Close()

	for {
		conn, err := l.Accept()

		if err != nil {
			utils.Log("failed to accept cli tcp connection")
		}

		go handleRequest(conn)
	}
}
