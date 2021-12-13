package topic

import (
	"log"
	"testing"
)

// param consts
const (
	clientId  = "google-client-id"
	topicName = "client-topic-name"
	subName   = "sub-name"
)

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

func TestCreateSubscription(t *testing.T) {
	t.Run("Create subscription test", func(t *testing.T) {
		Init(clientId)

		err := CreateSubscription(topicName, subName)

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Create subscription test is successful.")
		}
	})
}

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
