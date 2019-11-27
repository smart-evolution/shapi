package cliserver

import (
	"encoding/json"
	"fmt"
	"github.com/smart-evolution/shapi/utils"
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

	msg := make(map[string]string)
	json.Unmarshal(buff[:n], &msg)

	cmd := msg["cmd"]
	param := msg["param"]

	fmt.Println(msg)

	if cmd == "status" {
		conn.Write([]byte("application up and running"))
	} else if cmd == "proxy" {
		devAddr := param
		devConn, err := net.Dial("tcp", devAddr)

		if err != nil {
			fmt.Println("error connecting to device " + devAddr)
			return
		}
		_, err = devConn.Write([]byte("CMDWHO"))

		resBuff := make([]byte, 512)
		devResBuff := make([]byte, 512)
		n, err := devConn.Read(devResBuff)

		if err != nil {
			fmt.Println("error retrieving device type")
			return
		}

		_, err = conn.Write(devResBuff[:n])

		for {
			n, _ := conn.Read(resBuff)
			devConn.Write(resBuff[:n])

			if string(resBuff[:n]) == "CMDLOK" {
				n, err := devConn.Read(devResBuff)

				if err != nil {
					fmt.Println("RES: error reading message from device")
					break
				}

				conn.Write(devResBuff[:n])
			} else if string(resBuff[:n]) == "CMDDIS" {
				conn.Close()
			}
		}
	} else {
		utils.Log("invalid cli command '" + cmd + "'")
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
