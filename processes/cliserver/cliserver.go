package cliserver

import (
    "os"
    "net"
    "github.com/smart-evolution/smarthome/utils"
)

const (
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func RunService() {
    utils.Log("RunService")

    l, err := net.Listen(CONN_TYPE, ":" + CONN_PORT)

    if err != nil {
        os.Exit(1)
    }

    defer l.Close()

    for {
        conn, err := l.Accept()

        if err != nil {
            os.Exit(1)
        }

        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    utils.Log("Connection received!!!")

    data := make([]byte, 512)
    n, err := conn.Read(data)
    if err != nil { panic(err)  }
    s := string(data[:n])

    utils.Log(s)

    conn.Close()
}
