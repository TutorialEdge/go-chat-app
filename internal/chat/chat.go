package chat

import (
	"context"
	"fmt"
	"log"
	"os"

	stream "github.com/GetStream/stream-chat-go/v5"
)

type Service struct {
	APIKey       string
	APISecret    string
	ServerUserID string
	Channel      *stream.Channel
}

// Creates a new chat service
func New() (*Service, error) {
	var service Service
	var err error
	service.ServerUserID = "my-server-user-id"
	service.APIKey = os.Getenv("getstream_api_key")
	service.APISecret = os.Getenv("getstream_api_secret")

	client, err := stream.NewClient(service.APIKey, service.APISecret)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	service.Channel, err = client.CreateChannel(
		context.Background(),
		"messaging",
		"channel-id",
		service.ServerUserID,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &service, nil
}

func (s *Service) Listen(ctx context.Context) error {
	return nil
}

func (s *Service) WatchChannel(ctx context.Context) error {
	resp := stream.Query(s.Channel)
	fmt.Printf("%+v", resp)
	return nil
}

func (s *Service) AddUser(ctx context.Context, userID string) error {
	resp, err := s.Channel.AddMembers(ctx, []string{userID}, stream.AddMemberOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}

// SendMessage - sends a message woo
func (s *Service) SendMessage(ctx context.Context, userID, msg string) error {
	msg, err := s.Channel.SendMessage(ctx, &stream.Message{Text: msg}, userID)
	if err != nil {
		return err
	}
	return nil
}
