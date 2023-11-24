package main

import (
	"AzureServiceBus/msg"
)

const connectionString = "Endpoint=sb://servicebusniuniu.servicebus.windows.net/;SharedAccessKeyName=test;SharedAccessKey=zXIhcfVeVr3h+ToeC1GHNV1cNGtQVCRvE+ASbOI91Ks=;EntityPath=myqueue"

func main() {

	client := msg.GetClient(connectionString)
	msg.SendMsg("12345", client)
	msg.ReceiveMsg(client)

}
