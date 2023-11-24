package msg

import (
	context2 "context"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"golang.org/x/net/context"
	"log"
)

func SendMsg(message string, client *azservicebus.Client) {

	sender, err := client.NewSender("myqueue", nil)
	if err != nil {
		log.Printf("%v", err)
	}
	defer func(sender *azservicebus.Sender, ctx context2.Context) {
		err := sender.Close(ctx)
		if err != nil {

		}
	}(sender, context.TODO())

	sbMessage := &azservicebus.Message{
		Body: []byte(message),
	}
	err = sender.SendMessage(context.TODO(), sbMessage, nil)
	if err != nil {
		log.Printf("%v", err)
	}

}
