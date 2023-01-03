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

	presetInfo, r, err := client.WorkflowPreset.PresetWorkflow(&zadig.GetWorkflowPresetNameOptions{
		Env:          "dev",
		WorkflowName: "show-demo",
		ProjectName:  "nancal-demo",
	})
	if err != nil {
		return
	}

	fmt.Println(presetInfo)

	fmt.Println(r)

}
