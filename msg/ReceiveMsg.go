package msg

import (
	context2 "context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"golang.org/x/net/context"
	"log"
)

func ReceiveMsg(client *azservicebus.Client) {
	receiver, err := client.NewReceiverForQueue("myqueue", nil) //Change myqueue to env var
	if err != nil {
		log.Printf("%v", err)
	}
	defer func(receiver *azservicebus.Receiver, ctx context2.Context) {
		err := receiver.Close(ctx)
		if err != nil {

		}
	}(receiver, context.TODO())

	messages, err := receiver.ReceiveMessages(context.TODO(), 2, nil)
	if err != nil {
		log.Printf("%v", err)
	}

	for _, message := range messages {
		body := message.Body
		fmt.Printf("%s\n", string(body))

		err = receiver.CompleteMessage(context.TODO(), message, nil)
		if err != nil {
			log.Printf("%v", err)
		}
	}
}
