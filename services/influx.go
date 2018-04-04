package services

import (
    "log"
    "github.com/influxdata/influxdb/client/v2"
)

var (
    InfluxClient    client.Client
    InfluxBp        client.BatchPoints
)

func InitInfluxService() {
    InfluxClient, err = client.NewHTTPClient(client.HTTPConfig{
        Addr:     "http://localhost:8086",
        Username: "",
        Password: "",
    })

    if err != nil {
        log.Println("services: ", err)
    }

    InfluxBp, err = client.NewBatchPoints(client.BatchPointsConfig{
        Database:  "smarthome",
        Precision: "s",
    })
}
