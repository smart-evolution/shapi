package interfaces

import (
    "github.com/smart-evolution/smarthome/models/agent"
    "github.com/influxdata/influxdb/client/v2"
    "gopkg.in/mgo.v2"
)

// IState - interface for datasource kept in the memory
type IState interface {
    SetIsAlerts(bool)
    IsAlerts() bool
    SetSendAlert(bool)
    SendAlert() bool
    Agents() []*agent.Agent
    FindAgentByID(string) (*agent.Agent, error)
}

// IDataFlux -  interface for datasource to keep IOT data
type IDataFlux interface {
    IsConnected() bool
    AddData(*client.Point) error
    GetData(client.Query) (*client.Response, error)
}

// IPersistance - interface for user settings and general purpose storage
type IPersistance interface {
    GetCollection(string) *mgo.Collection
}

// IMailer - interface for mailer
type IMailer interface {
    AddRecipient(string)
    SendEmail(string, string)
    BulkEmail(string)
}

//IAgent - agent interface
type IAgent interface {
    URL()   string
}

// IHomeBot - homebot interface
type IHomeBot interface {
    FindAgentByID(string) (agent.Agent, error)
    RunService()
}

