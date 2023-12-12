/*
Copyright © 2023 Alireza Arzehgar <alirezaarzehgar82@gmail.com>
*/
package cmd

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"

	"github.com/BaseMax/KafkaDataProcessingGo/models"
)

// producerCmd represents the producer command
var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "Publish dataset to Kafka",
	Long:  `Read dataset and publish its content to Apache Kafka`,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("bootstrap-server")
		topic, _ := cmd.Flags().GetString("topic")
		partition, _ := cmd.Flags().GetInt("partition")
		path, _ := cmd.Flags().GetString("dataset")
		fakeDelay, _ := cmd.Flags().GetBool("fake-delay")

		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		conn, err := kafka.DialLeader(context.Background(), "tcp", server, topic, partition)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		var activities []models.UserActivity
		json.NewDecoder(f).Decode(&activities)
		for _, activity := range activities {
			data, err := json.Marshal(activity)
			if err != nil {
				log.Println(err)
				continue
			}

			_, err = conn.WriteMessages(kafka.Message{Value: data})
			if err != nil {
				log.Println(err)
			}

			if fakeDelay {
				time.Sleep(time.Nanosecond * time.Duration(rand.Intn(300)+200))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)
	producerCmd.PersistentFlags().StringP("bootstrap-server", "s", "localhost:9092", "Kafka server address")
	producerCmd.PersistentFlags().StringP("topic", "t", "activities", "Kafka topic")
	producerCmd.PersistentFlags().IntP("partitions", "p", 1, "Kafka topic partitions")
	producerCmd.PersistentFlags().StringP("dataset", "d", "sample_data.json", "Dataset path")
	producerCmd.PersistentFlags().BoolP("fake-delay", "f", false, "Dataset path")
}
