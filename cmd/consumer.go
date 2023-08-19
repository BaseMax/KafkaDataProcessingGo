/*
Copyright © 2023 Alireza Arzehgar <alirezaarzehgar82@gmail.com>
*/
package cmd

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"

	"github.com/BaseMax/KafkaDataProcessingGo/models"
)

var gauges = make(map[string]prometheus.Gauge, 0)

func prepareMonitoring() {
	prom := prometheus.NewRegistry()

	for _, action := range models.USER_ACTIONS {
		gauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "consumer",
			Name:      action,
		})
		gauges[action] = gauge
		prom.MustRegister(gauge)
	}

	go func() {
		for _, gauge := range gauges {
			time.Sleep(time.Second)
			gauge.Set(0)
		}
	}()

	go func() {
		http.Handle("/metrics", promhttp.HandlerFor(prom, promhttp.HandlerOpts{}))
		http.ListenAndServe(":8000", nil)
	}()
}

func StreamAndMonitoring(cmd *cobra.Command) {
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

		for action := range gauges {
			if action == activity.Action {
				gauges[action].Inc()
			}
		}
	}

}

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Consume messages",
	Long:  `Subscribe to kafka topic and monitor data. Then send it to prometheus.`,
	Run: func(cmd *cobra.Command, args []string) {
		prepareMonitoring()

		StreamAndMonitoring(cmd)
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)
	consumerCmd.PersistentFlags().StringP("bootstrap-server", "s", "localhost:9092", "Kafka server address")
	consumerCmd.PersistentFlags().StringP("topic", "t", "activities", "Kafka topic")
	consumerCmd.PersistentFlags().IntP("partitions", "p", 1, "Kafka topic partitions")
}
