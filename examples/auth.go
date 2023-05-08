package main

import (
	"fmt"
	"log"

	"github.com/joker-bai/go-zadig"
)

func authExample() {
	client, err := zadig.NewClient("token", zadig.WithBaseURL("https://www.koderover.com"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	fmt.Println(client)
}
