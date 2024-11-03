package manager

import (
	"context"
	"fmt"
	"github.com/ram02z/gh-weekly-log/api"
	"github.com/ram02z/gh-weekly-log/model"
	"github.com/ram02z/gh-weekly-log/transformer"
)

func GetIssues() ([]model.Issue, error) {
	ctx := context.Background()
	issueClient, err := api.DefaultIssueClient()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	params := api.IssueListOptions{
		Filter: "assigned",
		State:  "all",
	}
	issues, err := issueClient.List(ctx, &params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return transformer.TransformIssues(issues), nil
}
