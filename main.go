package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Stanley17932/microservice-with-golang/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
