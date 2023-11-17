package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/marktlinn/go_cha_CHIng/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.StartServer(ctx)
	if err != nil {
		fmt.Printf("App failed to start %v\n", err)
	}

}
