package Flight2BQ

import (
	"encoding/json"
	"log"
	"os"
	"context"
	"cloud.google.com/go/bigquery"
)

type position struct {
	Timestamp int64   `json:"timestamp"`
	Hexid     string  `json:"hexId"`
	Ident     string  `json:"ident"`
	Squawk    int64   `json:"squawk"`
	Alt       int64   `json:"alt"`
	Speed     int64   `json:"speed"`
	AirGround string  `json:"airground"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Heading   int64   `json:"heading"`
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func Flight2BQ(ctx context.Context, msg PubSubMessage) error {
	var bqClient *bigquery.Client
	bq, _ := bigquery.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	bqClient = bq
	uploader := bqClient.Dataset(os.Getenv("DATASET")).Table(os.Getenv("TABLE")).Uploader()
	var pos position

	if err := json.Unmarshal(msg.Data, &pos); err != nil {
		log.Printf("could not decode message data: %#v", msg)
		return err
	}
	items := []*position{
		&pos,
	}
	if err := uploader.Put(ctx, items); err != nil {
		log.Printf("Failed to insert row: %v", err)
		return err
	}
	log.Printf("Inserted")
	return nil
}