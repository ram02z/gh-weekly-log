package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ram02z/gh-weekly-log/client"
)

func main() {
	ctx := context.Background()
	userClient, err := client.DefaultUserClient()
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

	issueClient, err := client.DefaultIssueClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	params := client.IssueListOptions{
		Filter: "assigned",
		State:  "all",
		Since:  time.Now().AddDate(-1, 0, 0),
	}
	issues, err := issueClient.List(ctx, &params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("number of assigned issues: %d\n", len(issues))
	for _, issue := range issues {
		fmt.Printf(
			"#%d: %s (Created by: %s)\n",
			issue.Number,
			issue.Title,
			issue.User.Login,
		)
	}
}
