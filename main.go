package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
	"os"
)

type Message struct {
	Id          string `json:"id"`
	FirstName      string `json:"firstName"`
	LastName int64  `json:"lastName"`
}
func handler(c context.Context, e events.SNSEvent) {

	for _, record := range e.Records {
		snsRecord := record.SNS

		FirstName, ok := snsRecord.MessageAttributes["FirstName"].(map[string]interface{})
		if ok && "pcp" == FirstName["Value"] {
			log.Infof("FirstName: %s\n", FirstName["Value"])
			message := snsRecord.Message
			var msg = Message{}
			err := json.Unmarshal([]byte(message), &msg)
			if err != nil {
				log.Errorf("Error Unmarshalling message %v", err.Error())
			}
			log.Infof("Message passed %v", msg)

		}

	}
}
func main() {
	cloudenv := os.Getenv("CLOUD_ENVIRONMENT")
	log.Infof("Started Topic subscriber lambda in %s", cloudenv)
	lambda.Start(handler)
}
