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
	topics := topic.Topics()

	for i, v := range topics.Topics {
		r := fmt.Sprintf("index:/%d/ - value:/%s/ ", i, v)
		fmt.Println(r)
	}

	// Create a new subscription
	topic.CreateSubscription(topicName,subName)

	// Get subscriptions
	topic.Subscriptions(topicName)

	// Publish a new message
	topic.Publish(topicName,"Fly me to the moon..")

	// Get message to sub
	topic.Receive(subName)

}
