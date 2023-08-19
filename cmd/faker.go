/*
Copyright © 2023 Alireza Arzehgar <alirezaarzehgar82@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/spf13/cobra"

	"github.com/BaseMax/KafkaDataProcessingGo/models"
)

// fakerCmd represents the faker command
var fakerCmd = &cobra.Command{
	Use:   "faker",
	Short: "Create Fake Dataset",
	Long:  `Create a fake dataset for sending to Apache Kafka`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("output")

		f, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		var activities []models.UserActivity
	loop:
		for timeout := time.After(time.Second * 10); ; {
			select {
			case <-timeout:
				break loop
			default:
			}

			var activity models.UserActivity
			err := faker.FakeData(&activity)
			if err != nil {
				log.Println(err)
				continue
			}
			activities = append(activities, activity)
		}

		json.NewEncoder(f).Encode(activities)
	},
}

func init() {
	rootCmd.AddCommand(fakerCmd)
	fakerCmd.PersistentFlags().StringP("output", "o", "sample_data.json", "Path of sample dataset")
}
