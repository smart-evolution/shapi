package dataflux

import (
    "log"
    "github.com/influxdata/influxdb/client/v2"
)

type DataFlux struct {
    Client          client.Client
    BatchPoints     client.BatchPoints
    IsConnected     bool
}

func New(dbURI string) DataFlux {
    IsConnected := false

    Client, err := client.NewHTTPClient(client.HTTPConfig{
        Addr: dbURI,
        Username: "",
        Password: "",
    })

    if err != nil {
        log.Println("services: ", err)
        return DataFlux{}
    }

    BatchPoints, err := client.NewBatchPoints(client.BatchPointsConfig{
        Database:  "smarthome",
        Precision: "s",
    })

    if err != nil {
        log.Println("services ", err)
        return DataFlux{}
    }

    IsConnected = true

    return DataFlux{
        Client,
        BatchPoints,
        IsConnected,
    }
}

