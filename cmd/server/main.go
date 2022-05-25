package main

import (
	"context"
	"fmt"
	"time"

	"github.com/TutorialEdge/go-chat-app/internal/chat"
)

func Run() {
	ctx := context.Background()
	chatSvc, err := chat.New()
	if err != nil {
		fmt.Println(err)
	}

	for {
		chatSvc.WatchChannel(ctx)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	Run()
}
