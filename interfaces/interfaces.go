package interfaces

import (
    "github.com/smart-evolution/smarthome/models/agent"
    "github.com/influxdata/influxdb/client/v2"
    "gopkg.in/mgo.v2"
)

type IState interface {
    SetIsAlerts(bool)
    IsAlerts() bool
    SetSendAlert(bool)
    SendAlert() bool
    Agents() []agent.Agent
    FindAgentByID(string) (agent.Agent, error)
}

type IDataFlux interface {
    IsConnected() bool
    AddData(*client.Point) error
    GetData(client.Query) (*client.Response, error)
}

type IPersistance interface {
    GetCollection(string) *mgo.Collection
}

type IMailer interface {
    AddRecipient(string)
    SendEmail(string, string)
    BulkEmail(string)
}

type IAgent interface {
    URL()   string
}

type IHomeBot interface {
    FindAgentByID(string) (agent.Agent, error)
    RunService()
}

