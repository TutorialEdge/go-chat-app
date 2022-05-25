package main

import (
	"context"
	"fmt"

	"github.com/TutorialEdge/go-chat-app/internal/chat"
)

func Run() error {
	ctx := context.Background()

	chatSvc, err := chat.New()
	if err != nil {
		return fmt.Errorf("failed to create chat service: %w", err)
	}

	err = chatSvc.AddUser(ctx, "1")
	if err != nil {
		fmt.Println(err)
	}

	err = chatSvc.SendMessage(ctx, "1", "Hello World")
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Failed to start app")
	}
}
