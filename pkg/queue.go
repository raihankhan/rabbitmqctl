package pkg

import (
	"errors"
	"fmt"
	rabbithole "github.com/michaelklishin/rabbit-hole/v3"
)

func GetClient(uri string, username string, password string) (*rabbithole.Client, error) {
	client, err := rabbithole.NewClient(uri, username, password)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error getting queues: %s", err))
	}
	return client, nil
}

func GetQueues(client *rabbithole.Client) ([]rabbithole.QueueInfo, error) {
	queues, err := client.ListQueues()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error getting queues: %s", err))
	}
	return queues, nil
}
