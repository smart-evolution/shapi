package state

import (
    "sync"
    "github.com/smart-evolution/smarthome/models/agent"
    "github.com/influxdata/influxdb/client/v2"
    "gopkg.in/mgo.v2"
)

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
    RunService(sync.WaitGroup)
}

var (
    // IsAlerts - are alerts turned on
    IsAlerts    bool
    // SendAlert - should alerts be emailed
    SendAlert   bool
    // DataFlux - gathered data entity
    DataFlux    IDataFlux
    // Persistance - data persistance entity
    Persistance IPersistance
    // Mailer - mailer entity
    Mailer      IMailer
    // HomeBot - homebot administrator entity
    HomeBot     IHomeBot
)
