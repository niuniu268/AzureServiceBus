package cmd

import (
	"AzureServiceBus/msg"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var message string
var connectionString string

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Read input.yaml
		yamlData, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			log.Fatalf("Error reading input.yaml: %v", err)
		}

		// Unmarshal YAML data
		var config map[string]string
		err = yaml.Unmarshal(yamlData, &config)
		if err != nil {
			log.Fatalf("Error unmarshalling YAML: %v", err)
		}

		// Get the connection string
		connectionString = config["connectionString"]
		fmt.Println("service called")
		client := msg.GetClient(connectionString)
		msg.SendMsg(message, client)

	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.PersistentFlags().String("foo", message, "A help for foo")

}
