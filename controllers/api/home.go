package api

import (
    "os"
    "log"
    "time"
    "regexp"
    "errors"
    "strings"
    "net/http"
    "encoding/json"
    "github.com/tarm/serial"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
    "github.com/influxdata/influxdb/client/v2"
)

var (
    config          *serial.Config
    port            *serial.Port
    err             error
    isConnected     bool
    influxClient    client.Client
    influxBp        client.BatchPoints
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

func influxDBClient() {
    influxClient, err = client.NewHTTPClient(client.HTTPConfig{
        Addr:     "http://localhost:8086",
        Username: "",
        Password: "",
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }
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

    influxDBClient()

    influxBp, err = client.NewBatchPoints(client.BatchPointsConfig{
        Database:  "smarthome",
        Precision: "s",
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    if err != nil {
        log.Fatal(err)
    }

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

    pt, _ := client.NewPoint(
        "home",
        map[string]string{ "home": "home" },
        map[string]interface{}{
            "temperature": temperature,
            "presence": presence,
        },
        time.Now(),
    )
    influxBp.AddPoint(pt)

    err = influxClient.Write(influxBp)

	data := struct {
		Temperature string  `json:"temperature"`
        Presence    string  `json:"presence"`
	} {
        temperature,
        presence,
	}

	json.NewEncoder(w).Encode(data)
}

