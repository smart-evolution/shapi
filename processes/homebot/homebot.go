package homebot

import (
	"fmt"
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/goutils/mailer"
	"github.com/influxdata/influxdb1-client/v2"
	"github.com/smart-evolution/shapi/data/dataflux"
	"github.com/smart-evolution/shapi/data/repositories/agentconfigs"
	"github.com/smart-evolution/shapi/data/repositories/state"
	"github.com/smart-evolution/shapi/domain/models/agent"
	"github.com/smart-evolution/shapi/domain/models/type1"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"sync"
	"time"
)

// HomeBot - struct for homebot administrator
type HomeBot struct {
	store        dataflux.IDataFlux
	state        state.Repository
	agentconfigs agentconfigs.Repository
	mailer       mailer.IMailer
}

// New - creates new instances of HomeBot
func New(
	store dataflux.IDataFlux,
	ac agentconfigs.Repository,
	mailer mailer.IMailer,
	st state.Repository,
) *HomeBot {
	return &HomeBot{
		store:        store,
		agentconfigs: ac,
		state:        st,
		mailer:       mailer,
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

	logger.Log("Temperature adjustment [" + tmpStr + " + " + agentConfig.TmpAdjust + " = " + tmpAdjustedStr + "]")
	return data
}

func persistDataFactory(
	store dataflux.IDataFlux,
	agentConfig agent.Config,
) func(agent.IAgent, map[string]interface{}) {
	return func(ia agent.IAgent, data map[string]interface{}) {
		a, ok := ia.(*type1.Type1)

		if !ok {
			logger.Log("assertion type error")
			return
		}

		logger.Log("Persisting data for agent [" + a.Name + "]")

		adjustedData := adjustValues(data, agentConfig)

		pt, _ := client.NewPoint(
			a.ID,
			map[string]string{"home": a.Name},
			adjustedData,
			time.Now(),
		)

		err := store.AddData(pt)

		if err != nil {
			logger.Log("failed adding agent to store")
		}
	}
}

func (hb *HomeBot) runCommunicationLoop() {
	for range time.Tick(time.Second * 10) {
		if hb.store.IsConnected() == false {
			logger.Log("cannot fetch packages, Influx is down")
			return
		}

		agents := hb.state.Agents()
		done := make(chan struct{})
		var wg sync.WaitGroup
		wg.Add(len(agents))

		for _, it1 := range agents {
			t1, ok := it1.(*type1.Type1)

			if !ok {
				continue
			}

			cnf, err := hb.agentconfigs.FindOneAgentConfig(bson.M{
				"agentId": t1.ID,
			})

			if err != nil {
				logger.Log("AgentConfig not found for agent [" + t1.Name + "]")
			}

			persistData := persistDataFactory(hb.store, cnf)

			go t1.FetchPackage(hb.mailer.BulkEmail, persistData, hb.state.IsAlerts(), &wg)
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
