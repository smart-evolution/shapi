package dataflux

import (
	"github.com/coda-it/goutils/logger"
	"github.com/influxdata/influxdb1-client/v2"
	"github.com/smart-evolution/shapi/utils"
)

// IDataFlux -  interface for datasource to keep IOT data
type IDataFlux interface {
	IsConnected() bool
	AddData(*client.Point) error
	GetData(client.Query) (*client.Response, error)
}

// DataFlux - datasource keeping IOT data
type DataFlux struct {
	Client      client.Client
	BatchPoints client.BatchPoints
	isConnected bool
}

// New - creates new entity of DataFlux
func New(dbURI string) *DataFlux {
	isConnected := false

	Client, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     dbURI,
		Username: "",
		Password: "",
	})

	if err != nil {
		logger.Log(err)
		return &DataFlux{}
	}

	BatchPoints, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "smarthome",
		Precision: "s",
	})

	if err != nil {
		logger.Log(err)
		return &DataFlux{}
	}

	isConnected = true

	return &DataFlux{
		Client,
		BatchPoints,
		isConnected,
	}
}

// IsConnected - checks is DataFlux instance connected
func (df *DataFlux) IsConnected() bool {
	return df.isConnected
}

// GetData - gets data from DataFlux instance based on passed query
func (df *DataFlux) GetData(q client.Query) (*client.Response, error) {
	res, err := df.Client.Query(q)
	return res, err
}

// AddData - adds data to instance of DataFlux
func (df *DataFlux) AddData(pt *client.Point) error {
	df.BatchPoints.AddPoint(pt)
	err := df.Client.Write(df.BatchPoints)
	return err
}
