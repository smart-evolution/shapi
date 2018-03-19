package api

import (
"log"
"strings"
"net/http"
"encoding/json"
"github.com/tarm/serial"
"github.com/oskarszura/gowebserver/router"
"github.com/oskarszura/gowebserver/session"
)


func CtrHome(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	config := &serial.Config{Name: "/dev/tty.usbmodem1411", Baud: 9600}
	port, oErr := serial.OpenPort(config)

	if oErr != nil {
		log.Fatal(oErr)
	}

	buf := make([]byte, 128)
	bufLen, pErr := port.Read(buf)

	if pErr != nil {
		log.Fatal(pErr)
	}

	data := struct {
		Temperature    string   	    `json:"temperature"`
	} {
		strings.Split(strings.Replace(string(buf[:bufLen]), "\n", "", -1), "\r")[0],
	}

	json.NewEncoder(w).Encode(data)
}

