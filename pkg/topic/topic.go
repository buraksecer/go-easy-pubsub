package topic

import (
	"cloud.google.com/go/pubsub"
	pubsubapp "cloud.google.com/go/pubsub/apiv1"
	"context"
	"fmt"
	"github.com/buraksecer/go-easy-pubsub/internal/error/clienterr"
	"github.com/buraksecer/go-easy-pubsub/internal/error/topicerr"
	"github.com/buraksecer/go-easy-pubsub/internal/model/response/topicres"
	"github.com/buraksecer/go-easy-pubsub/pkg/client"
	"google.golang.org/api/iterator"
	pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"
	"log"
	"time"
)

// Local client id (pointer)
var _clientId *string

// Init must call one time every operation
func Init(clientId string) {
	_clientId = &clientId
}

// Create a new topic in project
func Create(topicName string) error {
	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		return clienterr.ErrClientCannotCreate
	}

	e := c.Client.Topic(topicName)
	ok, err := e.Exists(ctx)
	if err != nil {
		return topicerr.ErrTopicCanNotExists
	}
	if ok {
		return topicerr.ErrTopicDoExist
	}

	topic, err := c.Client.CreateTopic(ctx, topicName)

	if err != nil {
		return topicerr.ErrTopicCanNotCreate
	}

	defer topic.Stop()

	return err
}

// Delete topic in project
func Delete(topicName string) (bool, error) {
	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		return false, clienterr.ErrClientCannotCreate
	}

	topic := c.Client.Topic(topicName)
	defer topic.Stop()
	if err := topic.Delete(ctx); err != nil {
		return false, topicerr.ErrTopicCanNotDelete
	}
	return true, nil
}

// Exists topic control in project
func Exists(topicName string) (bool, error) {
	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		return false, clienterr.ErrClientCannotCreate
	}

	e := c.Client.Topic(topicName)
	return e.Exists(ctx)
}

// Topics list in project
func Topics() (topicres.Topics, error) {
	topics := topicres.Topics{}

	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		return topics, clienterr.ErrClientCannotCreate
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

	return topics, nil
}

// CreateSubscription a new sub in project
func CreateSubscription(topicName string, subName string) error {
	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		return clienterr.ErrClientCannotCreate
	}

	exist, errExists := Exists(topicName)
	if errExists != nil {
		return errExists
	}
	if !exist {
		return topicerr.ErrTopicDoNotExist
	}

	topic := c.Client.Topic(topicName)

	defer topic.Stop()

	sub, err := c.Client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic:            topic,
		AckDeadline:      10 * time.Second,
		ExpirationPolicy: time.Duration(0),
	})

	if err != nil {
		return err
	}
	_ = sub

	return nil
}

// Subscriptions list in a project
func Subscriptions(topicName string) (topicres.Subscriptions, error) {
	subscriptions := topicres.Subscriptions{}

	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		return subscriptions, clienterr.ErrClientCannotCreate
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
		subscriptions.Subscriptions = append(subscriptions.Subscriptions, sub.ID())
	}

	return subscriptions, nil
}

// Publish a new message to sub
func Publish(topicName string, message interface{}) (bool, error) {
	c, ctx, err := client.Create(*_clientId)

	defer client.Close(c)

	if err != nil {
		return false, clienterr.ErrClientCannotCreate
	}

	topic := c.Client.Topic(topicName)

	defer topic.Stop()

	m := fmt.Sprintf("%v", message.(interface{}))

	var results []*pubsub.PublishResult
	r := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(m),
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

// Receive get message to sub
func Receive(subName string) ([]topicres.Receive, error) {
	var receiveResponse []topicres.Receive

	ctx := context.Background()

	c, err := pubsubapp.NewSubscriberClient(ctx)
	if err != nil {
		return receiveResponse, err
	}
	defer c.Close()

	sub := fmt.Sprintf("projects/%s/subscriptions/%s", *_clientId, subName)

	req := &pubsubpb.PullRequest{
		Subscription: sub,
		MaxMessages:  1,
	}
	resp, err := c.Pull(ctx, req)
	if err != nil {
		return receiveResponse, err
	}

	if len(resp.ReceivedMessages) == 0 {
		return receiveResponse, topicerr.ErrTopicReceivedMessagesNotFound
	}

	for _, m := range resp.ReceivedMessages {
		receive := topicres.Receive{
			AckId: m.AckId,
			Data:  m.Message.Data,
		}

		err := c.Acknowledge(ctx, &pubsubpb.AcknowledgeRequest{
			Subscription: sub, AckIds: []string{m.AckId},
		})

		if err != nil {
			receive.Error = err.Error()
		} else {
			receive.Result = true
		}
		receiveResponse = append(receiveResponse, receive)
	}

	return receiveResponse, nil
}
