package type1

import (
	"errors"
	"github.com/smart-evolution/smarthome/models/agent"
	"github.com/smart-evolution/smarthome/utils"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	separator     = "\\|"
	tmpPattern    = "[0-9]+\\.[0-9]+"
	motionPattern = "-?[0-9]+"
	gasPattern    = "[0-1]"
	soundPattern  = "([0-9]+\\.[0-9]+)|(inf)"
	pkgPattern    = "<" +
		tmpPattern + separator +
		motionPattern + separator +
		gasPattern + separator +
		soundPattern +
		"\\>"
)

// IType1 - infetrace for IType1 hardware
type IType1 interface {
	FetchPackage(
		alertNotifier func(string),
		persistData func(agent.IAgent, map[string]interface{}),
		isAlerts bool,
		wg *sync.WaitGroup,
	)
}

// Type1 - hardware entity
type Type1 struct {
	agent.Agent
	tmpNotifyTime    time.Time
	motionNotifyTime time.Time
	gasNotifyTime    time.Time
}

// New - creates new entity of Agent
func New(id string, name string, ip string, agentType string) *Type1 {
	a := agent.New(id, name, ip, agentType)

	return &Type1{
		Agent: *a,
	}
}

func getPackageData(stream string) (string, error) {
	pkgRegExp, _ := regexp.Compile(pkgPattern)
	dataPackage := pkgRegExp.FindString(stream)

	if dataPackage == "" {
		return "", errors.New("agent/getPackageData: Data stream doesn't contain valid package (" + stream + ")")
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
func (a *Type1) FetchPackage(
	alertNotifier func(string),
	persistData func(agent.IAgent, map[string]interface{}),
	isAlerts bool,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	utils.Log("fetching data from agent [" + a.Name() + "]")
	apiURL := "http://" + a.IP() + "/api"
	response, err := http.Get(apiURL)

	if err != nil {
		a.SetIsOnline(false)
		utils.Log("data fetching request to agent [" + a.Name() + "] failed")
		return
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		a.SetIsOnline(false)
		utils.Log("agent '"+a.Name()+"'", err)
		return
	}

	unwrappedData, err := getPackageData(string(contents))

	if err != nil {
		a.SetIsOnline(false)
		utils.Log("agent '"+a.Name()+"'", err)
		return
	}

	a.SetIsOnline(true)

	temperature := getTemperature(unwrappedData)
	motion := getMotion(unwrappedData)
	gas := getGas(unwrappedData)
	sound := getSound(unwrappedData)

	if isAlerts == true {
		if t, err := strconv.ParseFloat(temperature, 32); err == nil {
			if t > 40 {
				now := time.Now()

				if now.Sub(a.tmpNotifyTime).Hours() >= 1 {
					a.tmpNotifyTime = now
					alertNotifier("[" + now.UTC().String() + "][" + a.Name() + "] temperature = " + temperature)
				}
			}
		}

		if motion != "0" {
			now := time.Now()

			if now.Sub(a.motionNotifyTime).Hours() >= 1 {
				a.motionNotifyTime = now
				alertNotifier("[" + now.UTC().String() + "][" + a.Name() + "] motion detected")
			}
		}

		if gas != "0" {
			now := time.Now()

			if now.Sub(a.gasNotifyTime).Hours() >= 1 {
				a.gasNotifyTime = now
				alertNotifier("[" + now.UTC().String() + "][" + a.Name() + "] gas detected")
			}
		}
	}

	data := map[string]interface{}{
		"temperature": temperature,
		"presence":    motion,
		"gas":         gas,
		"sound":       sound,
		"agent":       a.Name(),
	}

	persistData(a, data)
}
