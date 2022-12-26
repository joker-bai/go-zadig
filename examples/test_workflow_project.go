package main

import (
	"fmt"
	"github.com/joker-bai/go-zadig"
	"log"
)

func main() {

	client, err := zadig.NewClient(
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiY3VpeW9uZyIsImVtYWlsIjoiY3VpeW9uZ0BuYW5jYWwuY29tIiwidWlkIjoiY2Y3YjRhN2MtNWE5MS0xMWVkLThlNTktNDZkYTQzOTY0YTUwIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiY3VpeW9uZyIsImZlZGVyYXRlZF9jbGFpbXMiOnsiY29ubmVjdG9yX2lkIjoibGRhcCIsInVzZXJfaWQiOiJjdWl5b25nIn0sImF1ZCI6InphZGlnIiwiZXhwIjo0ODI1MDI3ODUyfQ.cNnMqOQ7pKSNLVvKN2tLfbFKFn4pI7aMSH8691Unwx0",
		zadig.WithBaseURL("https://leyancd.nancalcloud.com"))
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
