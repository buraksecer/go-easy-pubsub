package client

import (
	"github.com/buraksecer/go-easy-pubsub/internal/error/clienterr"
	"log"
	"testing"
)

// clientId const
const (
	clientId = "google-client-id"
)

// TestCreate testing create the client
func TestCreate(t *testing.T) {

	t.Run("Create Client Test", func(t *testing.T) {
		c, ctx, err := Create(clientId)

		if err != nil {
			t.Error(clienterr.ErrClientCannotCreate)
		}

		_ = ctx

		log.Println("Create client test is  Successful.")

		defer Close(c)
	})

}

// TestClose  testing close the client
func TestClose(t *testing.T) {

	t.Run("Close Client Test", func(t *testing.T) {
		c, ctx, err := Create(clientId)

		_ = ctx

		if err != nil {
			t.Error(err)
		}

		err = Close(c)

		if err != nil {
			t.Error(err)
		}

		log.Println("Close client test is  Successful.")
	})
}
