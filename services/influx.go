package services

import (
    "os"
    "log"
    "github.com/influxdata/influxdb/client/v2"
)

var (
    err             error
    InfluxClient    client.Client
    InfluxBp        client.BatchPoints
    InfluxConnected bool
)

func InitInfluxService() {
    InfluxConnected = false

    InfluxClient, err = client.NewHTTPClient(client.HTTPConfig{
        Addr: os.Getenv("INFLUXADDR"),
        Username: "",
        Password: "",
    })

    if err != nil {
        log.Println("services: ", err)
        return
    }

    InfluxBp, err = client.NewBatchPoints(client.BatchPointsConfig{
        Database:  "smarthome",
        Precision: "s",
    })

    if err != nil {
        log.Println("services ", err)
        return
    }

    InfluxConnected = true
}
