package services

import (
    "os"
    "log"
    "time"
    "regexp"
    "errors"
    "strings"
    "github.com/tarm/serial"
    "github.com/influxdata/influxdb/client/v2"
)

var (
    config          *serial.Config
    port            *serial.Port
    err             error
    isConnected     bool
)

func getPackageData(stream string) (string, error) {
    pkgRegExp, _ := regexp.Compile("<[0-9]+\\.[0-9]+\\|[0-9]+>")
    dataPackage := pkgRegExp.FindString(stream)

    if dataPackage == "" {
        return "", errors.New("Data stream not valid (" + stream + ")")
    }

    return strings.Split(strings.Replace(dataPackage, "<", "", -1), ">")[0], nil
}

func getTemperature(data string) string {
    return strings.Split(data, "|")[0]
}

func getPresence(data string) string {
    return strings.Split(data, "|")[1]
}

func fetchPackage() {
    if isConnected == false {
        InitHomeService()
    }

    buf := make([]byte, 128)
    bufLen, err := port.Read(buf)

    if err != nil {
        isConnected = false
        log.Println("services: ", err)
        return
    }

    dataStream := string(buf[:bufLen])

    unwrappedData, err := getPackageData(dataStream)

    if err != nil {
        log.Println("services: ", err)
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
    InfluxBp.AddPoint(pt)

    err = InfluxClient.Write(InfluxBp)
}

func InitHomeService() {
    isConnected = false;
    config = &serial.Config{Name: os.Getenv("SERIAL_PORT"), Baud: 9600}
    port, err = serial.OpenPort(config)

    if err != nil {
        log.Println(err)
        return
    }

    isConnected = true;
}

func RunHomeService() {
    for range time.Tick(time.Second * 1){
        fetchPackage()
    }
}


