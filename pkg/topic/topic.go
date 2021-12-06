package topic

import (
	"cloud.google.com/go/pubsub"
	pubsubapp "cloud.google.com/go/pubsub/apiv1"
	"context"
	"encoding/json"
	"fmt"
	"github.com/buraksecer/go-easy-pubsub/internal/error/clienterr"
	"github.com/buraksecer/go-easy-pubsub/internal/error/convererr"
	"github.com/buraksecer/go-easy-pubsub/internal/error/topicerr"
	"github.com/buraksecer/go-easy-pubsub/internal/model/response/topicres"
	"github.com/buraksecer/go-easy-pubsub/pkg/client"
	"google.golang.org/api/iterator"
	pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"
	"log"
	"time"
)

var _clientId *string

func Init(clientId string) {
	_clientId = &clientId
}

func Create(topicName string) error {

	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		return clienterr.ClientCannotCreate
	}

	e := c.Client.Topic(topicName)
	ok, err := e.Exists(ctx)
	if err != nil {
		return topicerr.TopicCanNotExists
	}
	if ok {
		return topicerr.TopicDoExist
	}

	topic, err := c.Client.CreateTopic(ctx, topicName)

	if err != nil {
		return topicerr.TopicCanNotCreate
	}

	defer topic.Stop()

	return err
}

func Delete(topicName string) (bool, error) {

	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		return false, clienterr.ClientCannotCreate
	}

	topic := c.Client.Topic(topicName)
	defer topic.Stop()
	if err := topic.Delete(ctx); err != nil {
		return false, topicerr.TopicCanNotDelete
	} else {
		return true, nil
	}
}

func Exists(topicName string) (bool, error) {

	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		return false, clienterr.ClientCannotCreate
	}

	e := c.Client.Topic(topicName)
	return e.Exists(ctx)
}

func Topics() topicres.TopicsModel {

	topics := topicres.TopicsModel{}

	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		log.Println(clienterr.ClientCannotCreate)
		return topics
	}

	it := c.Client.Topics(ctx)

	for {
		t, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
		}

		topics.Topics = append(topics.Topics, t.String())
	}

	return topics
}

func CreateSubscription(topicName string, subName string) {
	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		log.Println(clienterr.ClientCannotCreate)
		return
	}

	exist, errExists := Exists(topicName)
	if errExists != nil {
		log.Println(errExists)
		return
	}
	if !exist {
		log.Println(topicerr.TopicDoNotExist)
		return
	}

	topic := c.Client.Topic(topicName)

	defer topic.Stop()

	sub, err := c.Client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic:            topic,
		AckDeadline:      10 * time.Second,
		ExpirationPolicy: time.Duration(0),
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(sub)
}

func Subscriptions(topicName string) {

	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		log.Println(clienterr.ClientCannotCreate)
	}

	topic := c.Client.Topic(topicName)

	defer topic.Stop()

	for subs := topic.Subscriptions(ctx); ; {

		sub, err := subs.Next()

		if err == iterator.Done {
			log.Println(err)
			break
		}
		if err != nil {
			log.Printf("Error: %s\n\n", err)
		}
		log.Println(sub)
	}
}

func Publish(topicName string, message interface{}) (bool, error) {

	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		return false, clienterr.ClientCannotCreate
	}

	topic := c.Client.Topic(topicName)

	defer topic.Stop()

	m, err := json.Marshal(message)
	if err != nil {
		return false, convererr.JsonMarshalCanNotBeParse
	}

	var results []*pubsub.PublishResult
	r := topic.Publish(ctx, &pubsub.Message{
		Data: m,
	})

	results = append(results, r)

	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			log.Printf("Publication did not occur, message ID: %s\n\n", id)
			log.Printf("Error: %s\n\n", err)
		}
		log.Printf("Published a message with a message ID: %s\n\n", id)
	}

	return true, nil
}

func Receive(subName string) {

	ctx := context.Background()

	c, err := pubsubapp.NewSubscriberClient(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	sub := fmt.Sprintf("projects/%s/subscriptions/%s", *_clientId, subName)

	req := &pubsubpb.PullRequest{
		Subscription: sub,
		MaxMessages:  1,
	}
	resp, err := c.Pull(ctx, req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Success - ", resp.GetReceivedMessages())
}
