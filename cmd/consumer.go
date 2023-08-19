/*
Copyright © 2023 Alireza Arzehgar <alirezaarzehgar82@gmail.com>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/BaseMax/KafkaDataProcessingGo/models"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Consume messages",
	Long:  `Subscribe to kafka topic and monitor data. Then send it to prometheus.`,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("bootstrap-server")
		topic, _ := cmd.Flags().GetString("topic")
		partition, _ := cmd.Flags().GetInt("partition")

		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{server},
			Topic:     topic,
			Partition: partition,
			MinBytes:  10e3,
			MaxBytes:  10e6,
		})
		r.SetOffset(0)
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Println(err)
				continue
			}

			var activity models.UserActivity
			err = json.Unmarshal(m.Value, &activity)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Printf("%#v\n", activity)
		}
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)
	consumerCmd.PersistentFlags().StringP("bootstrap-server", "s", "localhost:9092", "Kafka server address")
	consumerCmd.PersistentFlags().StringP("topic", "t", "activities", "Kafka topic")
	consumerCmd.PersistentFlags().IntP("partitions", "p", 1, "Kafka topic partitions")
}
