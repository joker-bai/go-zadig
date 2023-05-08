package main

import (
	"fmt"
	"log"

	"github.com/joker-bai/go-zadig"
)

func test_workflow_project() {

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
