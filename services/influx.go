package services

import (
    "log"
    "github.com/influxdata/influxdb/client/v2"
)

var (
    InfluxClient    client.Client
    InfluxBp        client.BatchPoints
)

func influxDBClient() {
    InfluxClient, err = client.NewHTTPClient(client.HTTPConfig{
        Addr:     "http://localhost:8086",
        Username: "",
        Password: "",
    })

    if err != nil {
        log.Fatalln(err)
    }
}

func InitInfluxService() {
    influxDBClient()

    InfluxBp, err = client.NewBatchPoints(client.BatchPointsConfig{
        Database:  "smarthome",
        Precision: "s",
    })
}
