# Kafka Data Processing with Go

Welcome to the Kafka Data Processing with Go project! This project showcases how to use Apache Kafka in combination with the Go programming language to build a data processing application. By following this example, you'll learn how to produce and consume data using Kafka topics, allowing you to develop scalable and efficient data processing pipelines.

In this project, we'll create a data processing application using the Go programming language and Apache Kafka. Imagine you're building a system that processes user activity data from a website and performs real-time analytics on it. Kafka will serve as the backbone for data streaming, enabling the efficient transfer of data between different components.

## Prerequisites

Before you begin, make sure you have the following prerequisites:

- Go programming language (Installation guide: Getting Started with Go)
- Apache Kafka (Installation guide: Kafka Quickstart)
- Git

## Setup

Clone this repository:

```bash
git clone https://github.com/basemax/KafkaDataProcessingGo.git
cd KafkaDataProcessingGo
```

Start the Kafka server and create the necessary topics (assuming you've already installed Kafka):

```bash
# Start the ZooKeeper server (if not already started)
bin/zookeeper-server-start.sh config/zookeeper.properties

# Start the Kafka server
bin/kafka-server-start.sh config/server.properties

# Create the required topics
bin/kafka-topics.sh --create --topic activities --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
```

Start prometheus server and config like this:
```yaml
scrape_configs:
  - job_name: KafkaProcessing
    metrics_path: /metrics
    honor_labels: false
    honor_timestamps: true
    scheme: http
    scrape_interval: 1s
    follow_redirects: true
    body_size_limit: 0
    sample_limit: 0
    label_limit: 0
    label_name_length_limit: 0
    label_value_length_limit: 0
    target_limit: 0
    static_configs:
      - targets:
          - "HOSTNAME:8000"
```

You can optionally using Grafana. Grafana configurations placed on `grafana/` directory.

### Usage

This project consists of two main components: Producer and Consumer.

### Run via docker

#### Docker setup

You can change .env configurations.

#### Start app

Start application and it's dependencies using docker compose.
```bash
make run
```
or
```bash
docker-compose up -d
```

You can check some endpoints to ensure checking health of system.
 - localhost:8000/metrics - Consumer application
 - localhost:29092 - Kafka server
 - localhost:9090 - Prometheus server
 - localhost:3000 - Grafana server

#### Config Grafana

Import datasource configurations using following command:
```bash
make grafana_import_ds
```

Import Grafana dashboards manually. You can copy json configuration on `grafana/dashboards/` directory and import it to Grafana.

NOTE: If you wanna build application image multiple times you can use `go mod vendor` command to keeping dependencies on container. With this technique build process will speed up.

### Producer

The producer generates mock user activity data and sends it to the Kafka topic. To run the producer:

```bash
go run . producer
```

Use `-f` option for creating fake delay on publishing Kafka messages.

The producer will continuously generate and send user activity data to the Kafka topic.

### Consumer

The consumer subscribes to the Kafka topic, processes the user activity data, and performs analytics. To run the consumer:

```bash
go run . consumer
```

The consumer will listen for incoming user activity data and process it accordingly.

### Faker

The faker create sample dataset for producer.

```bash
go run . faker
```

## Sample Dataset

<!-- For your convenience, we've included a sample dataset in the sample_data directory. This dataset contains mock user activity logs that you can use to test the application. -->

This application deals with tracking user activities on an e-commerce website. Here's a simple example of a JSON-based user activity dataset:

```json
[
  {
    "user_id": "user123",
    "timestamp": "2023-08-18T10:00:00Z",
    "action": "view",
    "product_id": "prod456"
  },
  {
    "user_id": "user456",
    "timestamp": "2023-08-18T11:30:00Z",
    "action": "add_to_cart",
    "product_id": "prod123"
  },
  {
    "user_id": "user789",
    "timestamp": "2023-08-18T12:15:00Z",
    "action": "purchase",
    "product_id": "prod789"
  },
  // More entries...
]
```

You can create a sample dataset file like sample_data.json in the root directory of your project with multiple such entries.

Please note that this is just a basic representation, and you can extend it with additional fields and more complex data as needed for your application.

For generating a larger dataset, you might consider using libraries like Faker (for generating realistic fake data) in combination with Go's built-in JSON handling capabilities. Here's a rough example of how you could generate a larger dataset using Faker:

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
)

type UserActivity struct {
	UserID     string    `json:"user_id"`
	Timestamp  time.Time `json:"timestamp"`
	Action     string    `json:"action"`
	ProductID  string    `json:"product_id"`
}

func main() {
	var activities []UserActivity

	for i := 0; i < 1000; i++ {
		activity := UserActivity{
			UserID:     faker.UUIDHyphenated(),
			Timestamp:  faker.DateUnix(),
			Action:     faker.RandomChoice([]string{"view", "add_to_cart", "purchase"}),
			ProductID:  faker.UUIDHyphenated(),
		}
		activities = append(activities, activity)
	}

	file, err := os.Create("sample_data.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(activities); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Sample data generated and saved to sample_data.json")
}
```

Remember that this is just a basic example to get you started. Depending on your application's needs, you might want to generate more complex data with a wider range of possible actions, timestamps, and user profiles.

## Contributing

Contributions are welcome! If you encounter any issues or want to add new features, feel free to open a pull request. For significant changes, please open an issue first to discuss your proposed changes.

## License

This project is licensed under the GPL-3.0 License - see the LICENSE file for details.

Copyright 2023, Max Base
