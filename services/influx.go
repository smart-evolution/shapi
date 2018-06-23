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
)

func InitInfluxService() {
    InfluxClient, err = client.NewHTTPClient(client.HTTPConfig{
        Addr: os.Getenv("INFLUXADDR"),
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
