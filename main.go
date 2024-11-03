package main

import (
	"context"
	"fmt"
	"github.com/ram02z/gh-weekly-log/api"
	"github.com/ram02z/gh-weekly-log/manager"
)

func main() {
	ctx := context.Background()
	userClient, err := api.DefaultUserClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err := userClient.Get(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("running as %s\n", user.Login)

	issues, err := manager.GetIssues()
	for _, issue := range issues {
		fmt.Printf("issue: %s\n", issue.Title)
		for _, pr := range issue.PullRequests {
			fmt.Printf("pr: %s\n", pr.URL)
		}
	}
}
