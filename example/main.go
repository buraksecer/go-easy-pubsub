package main

import (
	"fmt"
	"github.com/buraksecer/go-easy-pubsub/pkg/topic"
)

// Const values
const (
	clientId  = "google-client-id"
	topicName = "client-topic-name"
	subName   = "sub-name"
)

func main() {

	// Init new topic operation
	topic.Init(clientId)

	// Create a new topic
	topic.Create(topicName)

	// Get a topic list
	topics, _ := topic.Topics()

	for _, v := range topics.Topics {
		fmt.Println(v)
	}

	// Create a new subscription
	topic.CreateSubscription(topicName, subName)

	// Get subscriptions
	topic.Subscriptions(topicName)

	// Publish a new message
	topic.Publish(topicName, "Fly me to the moon..")

	// Get message to sub
	topic.Receive(subName)

}
