package main

import (
	"fmt"
	"net/http"

	"github.com/JordanRad/notification-system-assignment/internal/message_queue"
	"github.com/JordanRad/notification-system-assignment/internal/service/gateway_service"
)

func main() {
	env, err := configFromEnv()
	if err != nil {
		panic(fmt.Errorf("error reading the config: %w", err))
	}

	kafkaProducer, err := message_queue.NewProducer(env.Kafka.URL)
	if err != nil {
		panic(fmt.Errorf("error creating kafka producer instance: %w", err))
	}
	gateway := gateway_service.NewService(kafkaProducer)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	mux.HandleFunc("/message", gateway.HandleNotification)

	serverAddr := fmt.Sprintf("%s:%d", env.HTTP.Host, env.HTTP.Port)
	fmt.Printf("Server started at %s\n", serverAddr)

	http.ListenAndServe(serverAddr, mux)
}
