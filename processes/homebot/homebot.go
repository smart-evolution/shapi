package homebot

import (
	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"github.com/smart-evolution/smarthome/datasources/dataflux"
	"github.com/smart-evolution/smarthome/datasources/persistence"
	"github.com/smart-evolution/smarthome/datasources/state"
	"github.com/smart-evolution/smarthome/models/agent"
	"github.com/smart-evolution/smarthome/models/type1"
	"github.com/smart-evolution/smarthome/services/email"
	"github.com/smart-evolution/smarthome/utils"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"sync"
	"time"
)

// HomeBot - struct for homebot administrator
type HomeBot struct {
	store       dataflux.IDataFlux
	state       state.IState
	persistence persistence.IPersistance
	mailer      email.IMailer
}

// New - creates new instances of HomeBot
func New(
	store dataflux.IDataFlux,
	p persistence.IPersistance,
	mailer email.IMailer,
	st state.IState,
) *HomeBot {
	return &HomeBot{
		store:       store,
		persistence: p,
		state:       st,
		mailer:      mailer,
	}
}

func adjustValues(
	data map[string]interface{},
	agentConfig agent.Config,
) map[string]interface{} {
	tmpObj := data["temperature"]

	tmpStr, _ := tmpObj.(string)
	tmpNumber, _ := strconv.ParseFloat(tmpStr, 32)

	tmpAdjustNumber, _ := strconv.ParseFloat(agentConfig.TmpAdjust, 32)
	tmpAdjustedNumber := tmpNumber + tmpAdjustNumber

	tmpAdjustedStr := fmt.Sprintf("%.2f", tmpAdjustedNumber)
	data["temperature"] = tmpAdjustedStr

	utils.Log("Temperature adjustment [" + tmpStr + " + " + agentConfig.TmpAdjust + " = " + tmpAdjustedStr + "]")
	return data
}

func persistDataFactory(
	store dataflux.IDataFlux,
	agentConfig agent.Config,
) func(agent.IAgent, map[string]interface{}) {
	return func(agent agent.IAgent, data map[string]interface{}) {
		utils.Log("Persisting data for agent [" + agent.Name() + "]")

		adjustedData := adjustValues(data, agentConfig)

		pt, _ := client.NewPoint(
			agent.ID(),
			map[string]string{"home": agent.Name()},
			adjustedData,
			time.Now(),
		)

		err := store.AddData(pt)

		if err != nil {
			utils.Log(err)
		}
	}
}

func (hb *HomeBot) runCommunicationLoop() {
	for range time.Tick(time.Second * 10) {
		if hb.store.IsConnected() == false {
			utils.Log("cannot fetch packages, Influx is down")
			return
		}

		var agentConfig agent.Config

		c := hb.persistence.GetCollection("agentConfigs")
		agents := hb.state.Agents()
		done := make(chan struct{})
		var wg sync.WaitGroup
		wg.Add(len(agents))

		for i := 0; i < len(agents); i++ {
			a := agents[i]

			err := c.Find(bson.M{
				"agentId": a.ID(),
			}).One(&agentConfig)

			if err != nil {
				utils.Log("AgentConfig not found for agent [" + a.Name() + "]")
			}

			persistData := persistDataFactory(hb.store, agentConfig)

			at1, ok := a.(type1.IAgentType1)
			if ok {
				go at1.FetchPackage(hb.mailer.BulkEmail, persistData, hb.state.IsAlerts(), &wg)
			}
		}

		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			continue
		case <-time.After(3 * time.Second):
			continue
		}
	}
}

// RunService - setup and run everything
func (hb *HomeBot) RunService() {
	hb.runCommunicationLoop()
}
