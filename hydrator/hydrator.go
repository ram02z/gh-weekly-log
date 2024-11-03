package hydrator

import (
	"context"
	"github.com/ram02z/gh-weekly-log/api"
	"github.com/ram02z/gh-weekly-log/api/schema"
	"github.com/ram02z/gh-weekly-log/model"
)

// HydrateIssue hydrates the issue model with pull requests from the timeline events
func HydrateIssue(issue model.Issue, r schema.Issue) (model.Issue, error) {
	// TODO: Use singleton
	issueClient, err := api.DefaultIssueClient()
	if err != nil {
		return issue, err
	}

	events, err := issueClient.ListTimelineEvents(context.Background(), r.Repository.Owner.Login, r.Repository.Name, r.Number, nil)
	if err != nil {
		return issue, err
	}

	for _, event := range events {
		if event.Event == "cross-referenced" {
			issue.PullRequests = append(issue.PullRequests, model.PullRequest{
				URL: event.Source.Issue.HTMLURL,
			})
		}
	}

	return issue, nil
}
