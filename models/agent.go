package models

import (
    "log"
    "io/ioutil"
    "net/http"
    "time"
    "regexp"
    "errors"
    "strings"
    "strconv"
    "github.com/smart-evolution/smarthome/utils"
)

const (
    separator = "\\|";
    tmpPattern = "[0-9]+\\.[0-9]+"
    motionPattern = "-?[0-9]+"
    gasPattern = "[0-1]"
    soundPattern = "([0-9]+\\.[0-9]+)|(inf)"
    pkgPattern = "<" +
    tmpPattern + separator +
    motionPattern + separator +
    gasPattern + separator +
    soundPattern +
    "\\>"
)

// Agent - hardware entity
type Agent struct {
    ID          string
    Name        string
    URL         string
    AgentType   string
}

var (
    // Agents - hardware agents list
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

// FetchPackage - fetches data packages
func (a Agent) FetchPackage(alertNotifier func(string), persistData func(Agent, map[string]interface{})) {
    response, err := http.Get(a.URL)

    if err != nil {
        log.Println("services: agent '" + a.Name + "'", err)
        return
    }

    defer response.Body.Close()

    contents, err := ioutil.ReadAll(response.Body)

    if err != nil {
        log.Println("services: agent '" + a.Name + "'", err)
        return
    }

    unwrappedData, err := getPackageData(string(contents))

    if err != nil {
        log.Println("services: agent '" + a.Name + "'", err)
        return
    }

    temperature := getTemperature(unwrappedData)
    motion := getMotion(unwrappedData)
    gas := getGas(unwrappedData)
    sound := getSound(unwrappedData)

    if utils.IsAlerts == true {
        if t, err := strconv.ParseFloat(temperature, 32); err == nil {
            if t > 40 {
                now := time.Now()

                if now.Sub(tmpNotifyTime).Hours() >= 1 {
                    tmpNotifyTime = now
                    alertNotifier("[" + now.UTC().String() + "][" + a.Name + "] temperature = " + temperature)
                }
            }
        }

        if motion != "0" {
            now := time.Now()

            if now.Sub(motionNotifyTime).Hours() >= 1 {
                motionNotifyTime = now
                alertNotifier("[" + now.UTC().String() + "][" + a.Name + "] motion detected")
            }
        }

        if gas != "0" {
            now := time.Now()

            if now.Sub(gasNotifyTime).Hours() >= 1 {
                gasNotifyTime = now
                alertNotifier("[" + now.UTC().String() + "][" + a.Name + "] gas detected")
            }
        }
    }

    data := map[string]interface{}{
        "temperature": temperature,
        "presence": motion,
        "gas": gas,
        "sound": sound,
        "agent": a.Name,
    }

    persistData(a, data)
}

// FindAgentByID - find corresponding agent by ID
func FindAgentByID(id string) (Agent, error) {
    for _, a := range Agents {
        if a.ID == id {
            return a, nil
        }
    }

    return  Agent{}, errors.New("No matching agent")
}
