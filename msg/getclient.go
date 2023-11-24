package msg

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"log"
)

func GetClient(connection string) *azservicebus.Client {
	client, err := azservicebus.NewClientFromConnectionString(connection, nil)

	if err != nil {
		panic(err)
	}

	log.Printf("client information is %v", client)

	return client
}
