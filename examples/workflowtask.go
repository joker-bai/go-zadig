package main

import (
	"fmt"
	"log"

	"github.com/joker-bai/go-zadig"
)

func listWorkflowTask() {
	client, err := zadig.NewClient("token", zadig.WithBaseURL("https://www.koderover.com"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	w, r, err2 := client.Workflow.GetWorkflowTask(&zadig.GetWorkflowTaskOptions{CommitId: "be3c8b5bb8"})
	if err2 != nil {
		log.Fatalf("get workflow task failed: %v", err2)
	}
	fmt.Println(w, r)
}

func createWorkflowTask() {
	client, err := zadig.NewClient("token", zadig.WithBaseURL("https://www.koderover.com"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	task, r, err := client.Workflow.CreateWorkflowTask(&zadig.CreateWorkflowTaskOptions{
		WorkflowName: "workflowname",
		EnvName:      "envname",
		Targets: []zadig.TargetArgs{
			{
				Name:        "servicename",
				ServiceType: "servicetype",
				Build: zadig.BuildArgs{
					Repos: []zadig.Repository{
						{
							RepoName: "reponame",
							Branch:   "branch",
						},
					},
				},
			},
		},
		Callback: zadig.Callback{
			CallbackUrl: "https://www.koderover.com/api/v1/callback",
		},
	})

	if err != nil {
		log.Fatalf("create workflow faield: %v", err)
	}
	fmt.Println(task, r.Body)
}
