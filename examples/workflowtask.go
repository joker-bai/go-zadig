package main

import (
	"fmt"
	"log"

	"github.com/joker-bai/go-zadig"
)

func getWorkflowTask() {
	client, err := zadig.NewClient("token", zadig.WithBaseURL("https://www.koderover.com"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	w, r, err2 := client.Workflow.GetWorkflowTaskStatus(&zadig.GetWorkflowTaskOptions{CommitId: "be3c8b5bb8"})
	if err2 != nil {
		log.Fatalf("get workflow task failed: %v", err2)
	}
	fmt.Println(w, r)
}

func restartWorkflowTask() {
	client, err := zadig.NewClient("token", zadig.WithBaseURL("https://www.koderover.com"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	r, err2 := client.Workflow.RestartWorkflowTask(
		&zadig.RestartWorkflowTaskOptions{
			ID:           1,
			PipelineName: "pipeline_name",
		})
	if err2 != nil {
		log.Fatalf("restart workflow task failed: %v", err2)
	}
	fmt.Println(r)
}

func canalWorkflowTask() {
	client, err := zadig.NewClient("token", zadig.WithBaseURL("https://www.koderover.com"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	r, err2 := client.Workflow.CanalWorkflowTask(
		&zadig.CanalWorkflowTaskOptions{
			ID:           1,
			PipelineName: "pipeline_name",
		})
	if err2 != nil {
		log.Fatalf("canal workflow task failed: %v", err2)
	}
	fmt.Println(r)
}

func createWorkflowTask() {
	client, err := zadig.NewClient("token", zadig.WithBaseURL("https://www.koderover.com"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	task, r, err := client.Workflow.ExecWorkflowTask(&zadig.ExecWorkflowTaskOptions{
		WorkflowName: "workflowname",
		ProjectName:  "projectname",
		Input: zadig.WorkflowInput{
			TargetEnv: "dev",
			Build: zadig.ExecBuildArgs{
				Enabled: true,
				ServiceList: []zadig.BuildServiceInfo{
					{
						ServiceModule: "service_module",
						ServiceName:   "service_name",
						RepoInfo: []zadig.RepositoryInfo{
							{
								CodehostName:  "my_codehost",
								RepoNamespace: "my_repo_namespace",
								RepoName:      "my_repo_name",
								Branch:        "master",
							},
						},
						Inputs: []zadig.UserInput{
							{
								Key:   "my_custom_key",
								Value: "my_custom_value",
							},
						},
					},
				},
			},
			Deploy: zadig.ExecDeployArgs{
				Enabled: true,
				Source:  "zadig",
				ServiceList: []zadig.DeployServiceInfo{
					{
						ServiceModule: "service_module",
						ServiceName:   "service_name",
						Image:         "docker_image",
					},
				},
			},
		},
	})

	if err != nil {
		log.Fatalf("create workflow faield: %v", err)
	}
	fmt.Println(task, r.Body)
}

func getWorkflowTaskDetail() {
	client, err := zadig.NewClient("token", zadig.WithBaseURL("https://www.koderover.com"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	w, r, err2 := client.Workflow.GetWorkflowTaskDetail(
		&zadig.GetWorkflowTaskDetailOptions{
			ID:           1,
			PipelineName: "pipeline_name",
		})
	if err2 != nil {
		log.Fatalf("get workflow task detail failed: %v", err2)
	}
	fmt.Println(w, r)
}
