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

	workflowList, r, err := client.WorkflowProject.GetWorkflowByPorectName(&zadig.GetWorkflowProjectNameOptions{
		PorjectName: "PorjectName",
	})
	if err != nil {
		return
	}

	fmt.Println(workflowList)

	fmt.Println(r)

}
