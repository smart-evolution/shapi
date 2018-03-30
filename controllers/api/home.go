package api

import (
    "os"
    "log"
    "regexp"
    "errors"
    "strings"
    "net/http"
    "encoding/json"
    "github.com/tarm/serial"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
)

var (
    config      *serial.Config
    port        *serial.Port
    err         error
    isConnected bool
)

func getPackageData(stream string) (string, error) {
    pkgRegExp, _ := regexp.Compile("<[0-9]+\\.[0-9]+\\|[0-9]+>")
    dataPackage := pkgRegExp.FindString(stream)

    if dataPackage == "" {
        return "", errors.New("HomeCtrl: Data package not valid")
    }

    return strings.Split(strings.Replace(dataPackage, "<", "", -1), ">")[0], nil
}

func getTemperature(data string) string {
    return strings.Split(data, "|")[0]
}

func getPresence(data string) string {
    return strings.Split(data, "|")[1]
}

func InitCtrlHome() {
    isConnected = false;
    config = &serial.Config{Name: os.Getenv("SERIAL_PORT"), Baud: 9600}
    port, err = serial.OpenPort(config)

    if err != nil {
        log.Println(err)
        return
    }

    isConnected = true;
}

func CtrHome(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    if isConnected == false {
        InitCtrlHome()
        return
    }

	buf := make([]byte, 128)
	bufLen, err := port.Read(buf)

	if err != nil {
        isConnected = false
		log.Println(err)
        return
	}

    dataStream := string(buf[:bufLen])

    unwrappedData, err := getPackageData(dataStream)

    if err != nil {
        return
    }

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

