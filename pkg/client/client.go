package client

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/buraksecer/go-easy-pubsub/internal/error/clienterr"
)

//Client is a pub and sub struct
type Client struct {
	Client *pubsub.Client
}

// Create a new client in project
func Create(clientId string) (Client, context.Context, error) {

	ctx := context.Background()

	c, err := pubsub.NewClient(ctx, clientId)

	if err != nil {
		err = clienterr.ClientCannotCreate
	}

	client := Client{
		Client: c,
	}

	return client, ctx, err
}

//Close is client operations
func Close(client Client) {
	client.Client.Close()
}
