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
git clone https://github.com/basemax/kafka-data-processing-go.git
cd kafka-data-processing-go
```

Start the Kafka server and create the necessary topics (assuming you've already installed Kafka):

```bash
# Start the ZooKeeper server (if not already started)
bin/zookeeper-server-start.sh config/zookeeper.properties

# Start the Kafka server
bin/kafka-server-start.sh config/server.properties

# Create the required topics
bin/kafka-topics.sh --create --topic userActivity --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
```

### Usage

This project consists of two main components: Producer and Consumer.

### Producer

The producer generates mock user activity data and sends it to the Kafka topic. To run the producer:

```bash
cd producer
go run main.go
```

The producer will continuously generate and send user activity data to the Kafka topic.

### Consumer

The consumer subscribes to the Kafka topic, processes the user activity data, and performs analytics. To run the consumer:

```bash
cd consumer
go run main.go
```

The consumer will listen for incoming user activity data and process it accordingly.

## Sample Dataset

For your convenience, we've included a sample dataset in the sample_data directory. This dataset contains mock user activity logs that you can use to test the application.

## Contributing

Contributions are welcome! If you encounter any issues or want to add new features, feel free to open a pull request. For significant changes, please open an issue first to discuss your proposed changes.

## License

This project is licensed under the GPL-3.0 License - see the LICENSE file for details.

Copyright 2023, Max Base
