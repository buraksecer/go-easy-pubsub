package main

import (
	"fmt"
	"github.com/buraksecer/go-easy-pubsub/pkg/topic"
)

const (
	clientId  = "google-client-id"
	topicName = "client-topic-name"
	subName   = "sub-name"
)

func main() {

	fmt.Println("Hi")

	topic.Init(clientId)

	topic.Create(topicName)

	topics := topic.Topics()

	for i, v := range topics.Topics {
		r := fmt.Sprintf("index:/%d/ - value:/%s/ ", i, v)
		fmt.Println(r)
	}

	topic.CreateSubscription(topicName,subName)

	topic.Subscriptions(topicName)

	topic.Publish(topicName,"Fly me to the moon..")

	topic.Receive(subName)

}
