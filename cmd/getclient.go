package cmd

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"log"
	"time"
)

type azclinet struct {
	*azservicebus.Client
}

func (client *azclinet) SendMsg(message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sender, err := client.NewSender("myqueue", nil)
	if err != nil {
		log.Printf("%v", err)
	}
	defer func(sender *azservicebus.Sender, ctx context.Context) {
		err := sender.Close(ctx)
		if err != nil {
		}
	}(sender, ctx)

	sbMessage := &azservicebus.Message{
		Body: []byte(message),
	}
	err = sender.SendMessage(ctx, sbMessage, nil)
	if err != nil {
		log.Printf("%v", err)
	}

}

func GetClient(connection string) *azservicebus.Client {
	client, err := azservicebus.NewClientFromConnectionString(connection, nil)

	if err != nil {
		panic(err)
	}

	//log.Printf("client information is %v", client)
	logrus.WithFields(logrus.Fields{
		"info": client,
	}).Info("client information: ")

	return client
}
