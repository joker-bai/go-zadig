package main

import (
	"fmt"
	"github.com/joker-bai/go-zadig"
	"log"
)

func main() {

	client, err := zadig.NewClient(
		"token", zadig.WithBaseURL("https://xx.com"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	projectList, r, err := client.Project.GetProjectList(&zadig.GetProjectListOptions{
		//Verbosity: "detailed",
		//Verbosity: "detailed",
		Verbosity: "detailed",
	})
	if err != nil {
		return
	}

	fmt.Println(projectList)

	fmt.Println(r)

}
