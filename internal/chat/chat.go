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
	service.APIKey = os.Getenv("GETSTREAM_API_KEY")
	service.APISecret = os.Getenv("GETSTREAM_API_SECRET")

	client, err := stream.NewClient(service.APIKey, service.APISecret)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := client.CreateChannel(
		context.Background(),
		"messaging",
		"channel-id",
		service.ServerUserID,
		&stream.ChannelRequest{},
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	service.Channel = resp.Channel
	return &service, nil
}

func (s *Service) Listen(ctx context.Context) error {
	return nil
}

func (s *Service) WatchChannel(ctx context.Context) error {
	resp, err := s.Channel.Query(ctx, &stream.QueryRequest{})
	if err != nil {
		return err
	}
	fmt.Printf("%+v", resp)
	return nil
}

func (s *Service) AddUser(ctx context.Context, userID string) error {
	resp, err := s.Channel.AddMembers(ctx, []string{userID})
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}

// SendMessage - sends a message woo
func (s *Service) SendMessage(ctx context.Context, userID, msg string) error {
	resp, err := s.Channel.SendMessage(ctx, &stream.Message{Text: msg}, userID)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}
