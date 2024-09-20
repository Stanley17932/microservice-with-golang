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

	signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
