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
	t.Run("Topic Create Test", func(t *testing.T) {

		Init(clientId)

		err := Create(topicName)

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Create topic test is Successful.")
		}
	})
}
