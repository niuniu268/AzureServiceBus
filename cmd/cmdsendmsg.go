package cmd

import (
	"AzureServiceBus/msg"
	"fmt"
	"github.com/spf13/cobra"
)

const connectionString = "Endpoint=sb://servicebusniuniu.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=9mG1rYLF68v3kvpFBUitFIi+SGbuQzHD2+ASbF1Fd54="

var message = ""

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service called")
		client := msg.GetClient(connectionString)
		msg.SendMsg(message, client)

	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.PersistentFlags().String("foo", message, "A help for foo")

}
