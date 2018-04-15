package services

import (
    "os"
    "log"
    "time"
    "regexp"
    "errors"
    "strings"
    "strconv"
    "github.com/tarm/serial"
    "github.com/influxdata/influxdb/client/v2"
    "github.com/oskarszura/smarthome/utils"
)

const (
    pkgPattern = "<[0-9]+\\.[0-9]+\\|-?[0-9]+\\|[0-1]>"
)

var (
    config              *serial.Config
    port                *serial.Port
    err                 error
    isConnected         bool
    tmpNotifyTime       time.Time
    motionNotifyTime    time.Time
    gasNotifyTime       time.Time
)

func getPackageData(stream string) (string, error) {
    pkgRegExp, _ := regexp.Compile(pkgPattern)
    dataPackage := pkgRegExp.FindString(stream)

    if dataPackage == "" {
        return "", errors.New("Data stream doesn't contain valid package (" + stream + ")")
    }

    return strings.Split(strings.Replace(dataPackage, "<", "", -1), ">")[0], nil
}

func getTemperature(data string) string {
    return strings.Split(data, "|")[0]
}

func getMotion(data string) string {
    return strings.Split(data, "|")[1]
}

func getGas(data string) string {
    return strings.Split(data, "|")[2]
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
    motion := getMotion(unwrappedData)
    gas := getGas(unwrappedData)

    if utils.IsAlerts == true {
        if t, err := strconv.ParseFloat(temperature, 32); err == nil {
            if t > 30 {
                now := time.Now()

                if now.Sub(tmpNotifyTime).Hours() >= 1 {
                    tmpNotifyTime = now
                    SendEmail("[" + now.UTC().String() + "] temperature = " + temperature)
                }
            }
        }

        if motion != "0" {
            now := time.Now()

            if now.Sub(motionNotifyTime).Hours() >= 1 {
                motionNotifyTime = now
                SendEmail("[" + now.UTC().String() + "] motion detected")
            }
        }

        if gas != "1" {
            now := time.Now()

            if now.Sub(gasNotifyTime).Hours() >= 1 {
                gasNotifyTime = now
                SendEmail("[" + now.UTC().String() + "] gas alert")
            }
        }
    }

    pt, _ := client.NewPoint(
        "home",
        map[string]string{ "home": "home" },
        map[string]interface{}{
            "temperature": temperature,
            "presence": motion,
            "gas": gas,
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


