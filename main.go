package main

import (
	"encoding/json"
	"log"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/pubsub"
	"github.com/namsral/flag"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var projectPtr string
var datasetPtr string
var tablePtr string
var subscriptionPtr string
var keyfilePtr string

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

func main() {
	flag.StringVar(&projectPtr, "project", "alex-olivier", "GCP Project")
	flag.StringVar(&datasetPtr, "dataset", "flighttracker_dev", "BigQuery Dataset")
	flag.StringVar(&tablePtr, "table", "aircraft_stream", "BigQuery Table")
	flag.StringVar(&subscriptionPtr, "subscription", "flight-data-prod-dev", "Pub/Sub Topic Name")
	flag.StringVar(&keyfilePtr, "keyfile", "default", "Path to keyfile")
	flag.Parse()
	log.Printf("Project: %s", projectPtr)
	log.Printf("Dataset: %s", datasetPtr)
	log.Printf("Table: %s", tablePtr)
	log.Printf("Subscription: %s", subscriptionPtr)
	log.Printf("Keyfile: %s", keyfilePtr)

	ctx := context.Background()

	var bqClient *bigquery.Client
	var pubsubClient *pubsub.Client
	var err error
	if keyfilePtr == "default" {
		bq, e := bigquery.NewClient(ctx, projectPtr)
		bqClient = bq
		ps, e := pubsub.NewClient(ctx, projectPtr)
		pubsubClient = ps
		err = e
	} else {
		bq, e := bigquery.NewClient(ctx, projectPtr, option.WithCredentialsFile(keyfilePtr))
		bqClient = bq
		ps, e := pubsub.NewClient(ctx, projectPtr, option.WithCredentialsFile(keyfilePtr))
		pubsubClient = ps
		err = e
	}
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	} else {
		log.Println("Clients ready")
	}
	uploader := bqClient.Dataset(datasetPtr).Table(tablePtr).Uploader()
	subscription := pubsubClient.Subscription(subscriptionPtr)
	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var pos position
		if err := json.Unmarshal(msg.Data, &pos); err != nil {
			log.Printf("could not decode message data: %#v", msg)
			msg.Ack()
			return
		}
		items := []*position{
			&pos,
		}
		if err := uploader.Put(ctx, items); err != nil {
			log.Printf("Failed to insert row: %v", err)
			return
		}
		msg.Ack()
		log.Printf("Inserted %s", msg.ID)
	})
	if err != nil {
		log.Printf("Failed to subscribe: %v", err)
	}
	ctx.Err
}
