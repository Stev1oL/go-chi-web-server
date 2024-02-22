package main

import (
	"context"
	"fmt"

	"github.com/Stev1oL/basic-web-server/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start application: ", err)
	}
}
