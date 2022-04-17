package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/suvam720/api/pkg/controller"
)

func main() {
	fmt.Println("Enter a user Id")
	var userId string
	fmt.Scan(&userId)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	if err := controller.Controller(ctx, userId); err != nil {
		log.Fatal(err)
	}
}
