package api

import (
    "os"
    "log"
    "regexp"
    "strings"
    "net/http"
    "encoding/json"
    "github.com/tarm/serial"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
)

var (
    config  *serial.Config
    port    *serial.Port
    err     error
)

func getPackageData(stream string) string {
    pkgRegExp, _ := regexp.Compile("<[0-9]+\\.[0-9]+\\|[0-1]>")
    dataPackage := pkgRegExp.FindString(stream)

    return strings.Split(strings.Replace(dataPackage, "<", "", -1), ">")[0]
}

func getTemperature(data string) string {
    return strings.Split(data, "|")[0]
}

func getPresence(data string) string {
    return strings.Split(data, "|")[1]
}

func InitCtrlHome() {
    config = &serial.Config{Name: os.Getenv("SERIAL_PORT"), Baud: 9600}
    port, err = serial.OpenPort(config)

    if err != nil {
        log.Fatal(err)
    }
}

func CtrHome(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	buf := make([]byte, 128)
	bufLen, pErr := port.Read(buf)

	if pErr != nil {
		log.Fatal(pErr)
	}

    dataStream := string(buf[:bufLen])

    unwrappedData := getPackageData(dataStream)
    temperature := getTemperature(unwrappedData)
    presence := getPresence(unwrappedData)

	data := struct {
		Temperature string  `json:"temperature"`
        Presence    string  `json:"presence"`
	} {
        temperature,
        presence,
	}

	json.NewEncoder(w).Encode(data)
}

