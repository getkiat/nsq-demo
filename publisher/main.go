package main

import (
	"fmt"
	"net/http"
	"os"

	configEnv "github.com/joho/godotenv"
	"github.com/nsqio/go-nsq"

	"github.com/getkiat/nsq-demo/publisher/src/handler"
	"github.com/getkiat/nsq-demo/publisher/src/pub"
)

const Topic = "grabfood"

func main() {
	//load environtment variables
	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	nsqHost, ok := os.LookupEnv("NSQ_HOST")
	if !ok {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	publisherConfig := nsq.NewConfig()

	publisher, err := pub.NewPublisher(nsqHost, publisherConfig)

	publisherHandler := handler.NewHttpHandler(Topic, publisher)

	http.Handle("/api/food/send", publisherHandler.PublishMessages())
	http.ListenAndServe(":3000", nil)
}
