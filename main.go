package main

import (
	"context"
	"fmt"

	"github.com/marktlinn/go_cha_CHIng/application"
)

func main() {
	app := application.New()

	err := app.StartServer(context.TODO())
	if err != nil {
		fmt.Printf("App failed to start %v\n", err)
	}
}
