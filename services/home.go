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
    pkgPattern = "<[0-9]+\\.[0-9]+\\|-?[0-9]+\\|[0-1]\\|[0-9]+\\.[0-9]+\\>"
)

type Agent struct {
    Name        string
    config      *serial.Config
    port        *serial.Port
    isConnected bool
}

var (
    Agents              []Agent
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

func getSound(data string) string {
    return strings.Split(data, "|")[3]
}

func writePackage(port *serial.Port) {
    _, err := port.Write([]byte("CMD001"))
    if err != nil {
        log.Println("services: ", err)
    }
}

func addAgent(name string, device string) {
    log.Println("services: adding home agent '" + name + "'")

    config := &serial.Config{Name: os.Getenv(device), Baud: 9600, ReadTimeout: time.Second * 5}
    port, err := serial.OpenPort(config)
    isConnected := false

    if err != nil {
        log.Println("services: agent '" + name + "'", err)
        return
    } else {
        isConnected = true
    }

    agent := Agent{
        Name: name,
        config: config,
        port: port,
        isConnected: isConnected,
    }

    Agents = append(Agents, agent)
}

func (a Agent) connect() {
    log.Println("services: connecting to home agent '" + a.Name + "'")

    var err error

    a.isConnected = false
    a.port, err = serial.OpenPort(a.config)

    if err != nil {
        log.Println("services:", err)
        return
    } else {
        a.isConnected = true
    }
}

func (a Agent) fetchPackage() {
    buf := make([]byte, 128)
    bufLen, err := a.port.Read(buf)

    if err != nil {
        a.isConnected = false
        log.Println("services: agent '" + a.Name + "'", err)
        return
    }

    dataStream := string(buf[:bufLen])

    unwrappedData, err := getPackageData(dataStream)

    if err != nil {
        log.Println("services:  agent '" + a.Name + "'", err)
        return
    }

    temperature := getTemperature(unwrappedData)
    motion := getMotion(unwrappedData)
    gas := getGas(unwrappedData)
    sound := getSound(unwrappedData)

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

        if gas != "0" {
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
            "sound": sound,
            "agent": a.Name,
        },
        time.Now(),
    )
    InfluxBp.AddPoint(pt)

    err = InfluxClient.Write(InfluxBp)
}

func RunHomeService() {
    addAgent("livingroom", "AGENTDEV1")
    addAgent("bedroom", "AGENTDEV2")

    for range time.Tick(time.Second * 10){
        for i := 0; i < len(Agents); i++ {
            a := Agents[i]

            if a.isConnected == false{
                a.connect()
            }

            a.fetchPackage()

            if utils.SendAlert {
                // writePackage(a.port)
                // utils.SendAlert = false
            }
        }
    }
}


