package msg

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"golang.org/x/net/context"
	"log"
	"time"
)

func SendMsg(message string, client *azservicebus.Client) {
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
