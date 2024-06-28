package main

import (
	"context"
	"demo-backend/application"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	demoApp := application.NewDemoApp(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := demoApp.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
