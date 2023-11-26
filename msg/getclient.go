package msg

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/sirupsen/logrus"
)

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
