package main

import (
	"fmt"
	"log"

	"github.com/joker-bai/go-zadig"
)

func test_environment() {

	client, err := zadig.NewClient(
		"token",
		zadig.WithBaseURL("http://xx.com"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	projectList, r, err := client.Environment.GetEvnByProjectName(&zadig.GetEvnByProjectNameOptions{
		PorjectName: "java-demo",
	})
	if err != nil {
		return
	}

	fmt.Println(projectList)

	fmt.Println(r)

}
