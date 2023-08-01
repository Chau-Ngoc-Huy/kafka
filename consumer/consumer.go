package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Set up the Kafka consumer configuration
	config := sarama.NewConfig()
	config.Version = sarama.V3_3_2_0
	config.Consumer.Return.Errors = true
	config.Consumer.Group.InstanceId = "first_group"

	// Create a new Kafka consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9094"}, config)
	if err != nil {
		log.Fatalln("Error creating Kafka consumer:", err)
	}
	defer consumer.Close()

	// Example topic
	topic := "topic-A"

	// Subscribe to the Kafka topic
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)

	if err != nil {
		log.Fatalln("Error subscribing to topic:", err)
	}
	defer partitionConsumer.Close()

	// Handle Kafka messages
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Received message: %s\n", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Println("Error:", err.Error())
		case <-signals:
			fmt.Println("Interrupted. Shutting down...")
			return
		}
	}
}
