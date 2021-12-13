package client

import (
	"log"
	"testing"
)

// clientId const
const (
	clientId = "google-client-id"
)

// TestCreate test to create the client
func TestCreate(t *testing.T) {
	t.Run("Create client test", func(t *testing.T) {
		c, ctx, err := Create(clientId)

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Create client test is successful.")
		}
		_ = ctx

		defer c.Client.Close()
	})
}

// TestClose  test to close the client
func TestClose(t *testing.T) {
	t.Run("Close client test", func(t *testing.T) {
		c, ctx, err := Create(clientId)

		if err != nil {
			t.Error(err)
		}

		err = Close(c)

		if err != nil {
			t.Error(err)
		} else {
			log.Println("Close client test is successful.")
		}

		_ = ctx
	})
}
