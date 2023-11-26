package main

import (
	"AzureServiceBus/cmd"
)

//const connectionString = "Endpoint=sb://servicebusniuniu.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=9mG1rYLF68v3kvpFBUitFIi+SGbuQzHD2+ASbF1Fd54="

func main() {

	//client := msg.GetClient(connectionString)
	//msg.SendMsg("", client)
	//msg.ReceiveMsg(client, 5)
	//msg.GetDeadLetterMessage(client)

	cmd.Execute()
}
