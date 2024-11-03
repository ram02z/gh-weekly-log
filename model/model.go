package model

type Issue struct {
	ID           int64
	Title        string
	URL          string
	PullRequests []PullRequest
}

type PullRequest struct {
	URL string
}
