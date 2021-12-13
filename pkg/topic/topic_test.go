package topic

import (
	"log"
	"testing"
)

// param consts
const (
	clientId  = "personaltraining-333910"
	topicName = "client-topic-name"
	subName   = "sub-name"
)

//TestCreate test to create topic
func TestCreate(t *testing.T) {
	t.Run("Topic create test", func(t *testing.T) {
		Init(clientId)

		err := Create(topicName)

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Create topic test is successful.")
		}
	})
}

//TestDelete test to delete topic
func TestDelete(t *testing.T) {
	t.Run("Topic delete test", func(t *testing.T) {
		Init(clientId)

		_, err := Delete(topicName)

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Delete topic test is successful.")
		}
	})
}

//TestExists test to exist topic
func TestExists(t *testing.T) {
	t.Run("Topic exists test", func(t *testing.T) {
		Init(clientId)

		_, err := Exists(topicName)

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Exists topic test is successful.")
		}
	})
}

//TestTopics test to get topic list
func TestTopics(t *testing.T) {
	t.Run("Get topic list test", func(t *testing.T) {
		Init(clientId)

		_, err := Topics()

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Get topic list test is successful.")
		}
	})
}

//TestCreateSubscription test to create subs
func TestCreateSubscription(t *testing.T) {
	t.Run("Create subscription test", func(t *testing.T) {
		Init(clientId)

		err := Create(topicName)

		if err != nil {
			t.Error(err)
		}

		err = CreateSubscription(topicName, subName)

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Create subscription test is successful.")
		}
	})
}

//TestSubscriptions test to get subs list
func TestSubscriptions(t *testing.T) {
	t.Run("Get subscription list test", func(t *testing.T) {
		Init(clientId)

		_, err := Subscriptions(topicName)

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Get subscription list test is successful.")
		}
	})
}

//TestPublish test to publish new message
func TestPublish(t *testing.T) {
	t.Run("Publish message test", func(t *testing.T) {
		Init(clientId)

		_, err := Publish(topicName, "Fly me to the moon..")

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Publish message test is successful.")
		}
	})
}

//TestReceive test to receive new message
func TestReceive(t *testing.T) {
	t.Run("Receive message test", func(t *testing.T) {
		Init(clientId)

		_, err := Receive(subName)

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Receive message test is successful.")
		}
	})
}

//TestExistsSubscription test to exist subs
func TestExistsSubscription(t *testing.T) {
	t.Run("Exists subscription test", func(t *testing.T) {
		Init(clientId)

		_, err := ExistsSubscription(subName)

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Exists subs test is successful.")
		}
	})
}
