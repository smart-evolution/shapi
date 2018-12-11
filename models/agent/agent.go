package agent

import (
    "log"
    "io/ioutil"
    "net/http"
    "time"
    "regexp"
    "errors"
    "strings"
    "strconv"
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
    iD          string
    name        string
    uRL         string
    agentType   string
}

// New - creates new entity of Agent
func New(id string, name string, url string, agentType string) Agent {
    return Agent{
        iD: id,
        name: name,
        uRL: url,
        agentType: agentType,
    }
}

// ID - iD getter
func (a *Agent) ID() string {
    return a.iD
}

// SetID - iD setter
func (a *Agent) SetID(id string) {
    a.iD = id
}

// Name - name getter
func (a *Agent) Name() string {
    return a.name
}

// SetName - name setter
func (a *Agent) SetName(name string) {
    a.name = name
}

// URL - uRL getter
func (a *Agent) URL() string {
    return a.uRL
}

// SetURL - uRL setter
func (a *Agent) SetURL(URL string) {
    a.uRL = URL
}

// AgentType - agentType getter
func (a *Agent) AgentType() string {
    return a.agentType
}

// SetAgentType - agentType setter
func (a *Agent) SetAgentType(agentType string) {
    a.agentType = agentType
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
func (a Agent) FetchPackage(alertNotifier func(string), persistData func(Agent, map[string]interface{}), isAlerts bool) {
    response, err := http.Get(a.uRL)

    if err != nil {
        log.Println("services: agent '" + a.name + "'", err)
        return
    }

    defer response.Body.Close()

    contents, err := ioutil.ReadAll(response.Body)

    if err != nil {
        log.Println("services: agent '" + a.name + "'", err)
        return
    }

    unwrappedData, err := getPackageData(string(contents))

    if err != nil {
        log.Println("services: agent '" + a.name + "'", err)
        return
    }

    temperature := getTemperature(unwrappedData)
    motion := getMotion(unwrappedData)
    gas := getGas(unwrappedData)
    sound := getSound(unwrappedData)

    if isAlerts == true {
        if t, err := strconv.ParseFloat(temperature, 32); err == nil {
            if t > 40 {
                now := time.Now()

                if now.Sub(tmpNotifyTime).Hours() >= 1 {
                    tmpNotifyTime = now
                    alertNotifier("[" + now.UTC().String() + "][" + a.name + "] temperature = " + temperature)
                }
            }
        }

        if motion != "0" {
            now := time.Now()

            if now.Sub(motionNotifyTime).Hours() >= 1 {
                motionNotifyTime = now
                alertNotifier("[" + now.UTC().String() + "][" + a.name + "] motion detected")
            }
        }

        if gas != "0" {
            now := time.Now()

            if now.Sub(gasNotifyTime).Hours() >= 1 {
                gasNotifyTime = now
                alertNotifier("[" + now.UTC().String() + "][" + a.name + "] gas detected")
            }
        }
    }

    data := map[string]interface{}{
        "temperature": temperature,
        "presence": motion,
        "gas": gas,
        "sound": sound,
        "agent": a.name,
    }

    persistData(a, data)
}

