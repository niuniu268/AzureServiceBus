package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var message string
var connectionString string

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send message",
	Long:  ` send message to the service bus at azure`,
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

		var clientInstance = &azclinet{
			Client: GetClient(connectionString),
		}
		clientInstance.SendMsg(message)

		logrus.WithFields(logrus.Fields{
			"content": message,
		}).Info("sending information")

	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.PersistentFlags().StringVar(&message, "msg", "message", "content of the message")

}
