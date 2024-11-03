package transformer

import (
	"fmt"
	"github.com/ram02z/gh-weekly-log/api/schema"
	"github.com/ram02z/gh-weekly-log/hydrator"
	"github.com/ram02z/gh-weekly-log/model"
)

// TransformIssues transforms the issues API response to issues with pull requests
func TransformIssues(issues []schema.Issue) []model.Issue {
	// Create list of issues
	issuesList := make([]model.Issue, 0)
	for _, r := range issues {
		if pr := r.PullRequest; pr != nil {
			// Only populating the issues
			continue
		}
		issue := model.Issue{
			ID:           r.ID,
			Title:        r.Title,
			URL:          r.HTMLURL,
			PullRequests: make([]model.PullRequest, 0),
		}
		issue, err := hydrator.HydrateIssue(issue, r)
		if err != nil {
			fmt.Println(err)
		}
		issuesList = append(issuesList, issue)
	}
	return issuesList
}
