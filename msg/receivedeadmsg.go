package msg

import (
	context2 "context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"golang.org/x/net/context"
)

func GetDeadLetterMessage(client *azservicebus.Client) {
	receiver, err := client.NewReceiverForQueue(
		"myqueue",
		&azservicebus.ReceiverOptions{
			SubQueue: azservicebus.SubQueueDeadLetter,
		},
	)
	if err != nil {
		panic(err)
	}
	defer func(receiver *azservicebus.Receiver, ctx context2.Context) {
		err := receiver.Close(ctx)
		if err != nil {

		}
	}(receiver, context.TODO())

	messages, err := receiver.ReceiveMessages(context.TODO(), 1, nil)
	if err != nil {
		panic(err)
	}

	for _, message := range messages {
		fmt.Printf("DeadLetter Reason: %s\nDeadLetter Description: %s\n", *message.DeadLetterReason, *message.DeadLetterErrorDescription)
		err := receiver.CompleteMessage(context.TODO(), message, nil)
		if err != nil {
			panic(err)
		}
	}
}
