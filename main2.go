package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func GetClient() *azservicebus.Client {
	namespace, ok := os.LookupEnv("AZURE_SERVICEBUS_HOSTNAME") //ex: myservicebus.servicebus.windows.net
	if !ok {
		panic("AZURE_SERVICEBUS_HOSTNAME environment variable not found")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	client, err := azservicebus.NewClient(namespace, cred, nil)
	if err != nil {
		panic(err)
	}
	return client
}

func SendMessage(message string, client *azservicebus.Client) {
	sender, err := client.NewSender("myqueue", nil)
	if err != nil {
		panic(err)
	}
	defer func(sender *azservicebus.Sender, ctx context.Context) {
		err := sender.Close(ctx)
		if err != nil {

		}
	}(sender, context.TODO())

	sbMessage := &azservicebus.Message{
		Body: []byte(message),
	}
	err = sender.SendMessage(context.TODO(), sbMessage, nil)
	if err != nil {
		panic(err)
	}
}

func SendMessageBatch(messages []string, client *azservicebus.Client) {
	sender, err := client.NewSender("myqueue", nil)
	if err != nil {
		panic(err)
	}
	defer func(sender *azservicebus.Sender, ctx context.Context) {
		err := sender.Close(ctx)
		if err != nil {

		}
	}(sender, context.TODO())

	batch, err := sender.NewMessageBatch(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	for _, message := range messages {
		err := batch.AddMessage(&azservicebus.Message{Body: []byte(message)}, nil)
		if errors.Is(err, azservicebus.ErrMessageTooLarge) {
			fmt.Printf("Message batch is full. We should send it and create a new one.\n")
		}
	}

	if err := sender.SendMessageBatch(context.TODO(), batch, nil); err != nil {
		panic(err)
	}
}

func GetMessage(count int, client *azservicebus.Client) {
	receiver, err := client.NewReceiverForQueue("myqueue", nil)
	if err != nil {
		panic(err)
	}
	defer func(receiver *azservicebus.Receiver, ctx context.Context) {
		err := receiver.Close(ctx)
		if err != nil {

		}
	}(receiver, context.TODO())

	messages, err := receiver.ReceiveMessages(context.TODO(), count, nil)
	if err != nil {
		panic(err)
	}

	for _, message := range messages {
		body := message.Body
		fmt.Printf("%s\n", string(body))

		err = receiver.CompleteMessage(context.TODO(), message, nil)
		if err != nil {
			panic(err)
		}
	}
}

func DeadLetterMessage(client *azservicebus.Client) {
	deadLetterOptions := &azservicebus.DeadLetterOptions{
		ErrorDescription: to.Ptr("exampleErrorDescription"),
		Reason:           to.Ptr("exampleReason"),
	}

	receiver, err := client.NewReceiverForQueue("myqueue", nil)
	if err != nil {
		panic(err)
	}
	defer func(receiver *azservicebus.Receiver, ctx context.Context) {
		err := receiver.Close(ctx)
		if err != nil {

		}
	}(receiver, context.TODO())

	messages, err := receiver.ReceiveMessages(context.TODO(), 1, nil)
	if err != nil {
		panic(err)
	}

	if len(messages) == 1 {
		err := receiver.DeadLetterMessage(context.TODO(), messages[0], deadLetterOptions)
		if err != nil {
			panic(err)
		}
	}
}

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
	defer func(receiver *azservicebus.Receiver, ctx context.Context) {
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

func main() {
	client := GetClient()

	fmt.Println("send a single message...")
	SendMessage("firstMessage", client)

	fmt.Println("send two messages as a batch...")
	messages := [2]string{"secondMessage", "thirdMessage"}
	SendMessageBatch(messages[:], client)

	fmt.Println("\nget all three messages:")
	GetMessage(3, client)

	fmt.Println("\nsend a message to the Dead Letter Queue:")
	SendMessage("Send message to Dead Letter", client)
	DeadLetterMessage(client)
	GetDeadLetterMessage(client)
}
