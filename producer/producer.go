package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
	"time"
)

func main() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		Topic:    "topic-B",
		Balancer: &kafka.RoundRobin{},
	}

	i := 1
	for {
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte("Key-A"),
				Value: []byte("Hello World! " + strconv.Itoa(i)),
			},
			//kafka.Message{
			//	Key:   []byte("Key-B"),
			//	Value: []byte("One! " + strconv.Itoa(i)),
			//},
			//kafka.Message{
			//	Key:   []byte("Key-C"),
			//	Value: []byte("Two! " + strconv.Itoa(i)),
			//},
		)

		if err != nil {
			log.Fatal("failed to write messages:", err)
		} else {
			fmt.Println("success to write a messages " + strconv.Itoa(i))
		}
		i++
		time.Sleep(1 * time.Second)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
