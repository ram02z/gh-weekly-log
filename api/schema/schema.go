package schema

import "time"

// Issue represents a GitHub issue on Repository
type Issue struct {
	ID                    int64           `json:"id"`
	NodeID                string          `json:"node_id"`
	URL                   string          `json:"url"`
	RepositoryURL         string          `json:"repository_url"`
	LabelsURL             string          `json:"labels_url"`
	CommentsURL           string          `json:"comments_url"`
	EventsURL             string          `json:"events_url"`
	HTMLURL               string          `json:"html_url"`
	Number                int             `json:"number"`
	State                 string          `json:"state"`
	StateReason           *string         `json:"state_reason,omitempty"`
	Title                 string          `json:"title"`
	Body                  *string         `json:"body,omitempty"`
	User                  *User           `json:"user"`
	Labels                []Label         `json:"labels"`
	Assignee              *User           `json:"assignee"`
	Assignees             []User          `json:"assignees"`
	Milestone             *Milestone      `json:"milestone,omitempty"`
	Locked                bool            `json:"locked"`
	ActiveLockReason      *string         `json:"active_lock_reason,omitempty"`
	Comments              int             `json:"comments"`
	PullRequest           *PullRequest    `json:"pull_request,omitempty"`
	ClosedAt              *time.Time      `json:"closed_at,omitempty"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`
	Draft                 bool            `json:"draft"`
	ClosedBy              *User           `json:"closed_by,omitempty"`
	BodyHTML              string          `json:"body_html"`
	BodyText              string          `json:"body_text"`
	TimelineURL           string          `json:"timeline_url"`
	Repository            Repository      `json:"repository"`
	PerformedViaGitHubApp *GitHubApp      `json:"performed_via_github_app,omitempty"`
	AuthorAssociation     string          `json:"author_association"`
	Reactions             *ReactionRollup `json:"reactions,omitempty"`
}

// User represents a GitHub user.
type User struct {
	Name              *string `json:"name,omitempty"`
	Email             *string `json:"email,omitempty"`
	Login             string  `json:"login"`
	ID                int64   `json:"id"`
	NodeID            string  `json:"node_id"`
	AvatarURL         string  `json:"avatar_url"`
	GravatarID        *string `json:"gravatar_id,omitempty"`
	URL               string  `json:"url"`
	HTMLURL           string  `json:"html_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReposURL          string  `json:"repos_url"`
	EventsURL         string  `json:"events_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	Type              string  `json:"type"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
	UserViewType      string  `json:"user_view_type"`
}

// Label represents a GitHub issue/pull request label.
type Label struct {
	ID          int64   `json:"id"`
	NodeID      string  `json:"node_id"`
	URL         string  `json:"url"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
	Default     bool    `json:"default"`
}

// Milestone represents a collection of issues and pull requests.
type Milestone struct {
	URL          string     `json:"url"`
	HTMLURL      string     `json:"html_url"`
	LabelsURL    string     `json:"labels_url"`
	ID           int        `json:"id"`
	NodeID       string     `json:"node_id"`
	Number       int        `json:"number"`
	State        string     `json:"state"`
	Title        string     `json:"title"`
	Description  *string    `json:"description,omitempty"`
	Creator      *User      `json:"creator,omitempty"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	ClosedAt     *time.Time `json:"closed_at,omitempty"`
	DueOn        *time.Time `json:"due_on,omitempty"`
}

// PullRequest represents a GitHub pull request.
type PullRequest struct {
	MergedAt *time.Time `json:"merged_at,omitempty"`
	DiffURL  *string    `json:"diff_url,omitempty"`
	HTMLURL  *string    `json:"html_url,omitempty"`
	PatchURL *string    `json:"patch_url,omitempty"`
	URL      *string    `json:"url,omitempty"`
}

// Repository represents a GitHub repository.
type Repository struct {
	ID          int64      `json:"id"`
	NodeID      string     `json:"node_id"`
	Name        string     `json:"name"`
	FullName    string     `json:"full_name"`
	Private     bool       `json:"private"`
	Owner       User       `json:"owner"`
	HTMLURL     string     `json:"html_url"`
	Description *string    `json:"description,omitempty"`
	Fork        bool       `json:"fork"`
	URL         string     `json:"url"`
	License     *License   `json:"license,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	PushedAt    *time.Time `json:"pushed_at,omitempty"`
}

// License represents a repository license.
type License struct {
	Key     string  `json:"key"`
	Name    string  `json:"name"`
	URL     *string `json:"url,omitempty"`
	SPDXID  *string `json:"spdx_id,omitempty"`
	NodeID  string  `json:"node_id"`
	HTMLURL string  `json:"html_url"`
}

// GitHubApp represents a GitHub app.
type GitHubApp struct {
	ID          int         `json:"id"`
	NodeID      string      `json:"node_id"`
	Owner       *User       `json:"owner,omitempty"`
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
	ExternalURL string      `json:"external_url"`
	HTMLURL     string      `json:"html_url"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Permissions Permissions `json:"permissions"`
	Events      []string    `json:"events"`
}

// Permissions represents GitHub app permissions.
type Permissions struct {
	Issues      string `json:"issues"`
	Checks      string `json:"checks"`
	Metadata    string `json:"metadata"`
	Contents    string `json:"contents"`
	Deployments string `json:"deployments"`
}

// ReactionRollup represents reactions to an issue or comment.
type ReactionRollup struct {
	URL        string `json:"url"`
	TotalCount int    `json:"total_count"`
	PlusOne    int    `json:"+1"`
	MinusOne   int    `json:"-1"`
	Laugh      int    `json:"laugh"`
	Confused   int    `json:"confused"`
	Heart      int    `json:"heart"`
	Hooray     int    `json:"hooray"`
	Eyes       int    `json:"eyes"`
	Rocket     int    `json:"rocket"`
}

// Timeline represents a timeline event.
type Timeline struct {
	// Event identifies the actual type of Event that occurred. Possible values
	// are:
	//
	//     assigned
	//       The issue was assigned to the assignee.
	//
	//     closed
	//       The issue was closed by the actor. When the commit_id is present, it
	//       identifies the commit that closed the issue using "closes / fixes #NN"
	//       syntax.
	//
	//     commented
	//       A comment was added to the issue.
	//
	//     committed
	//       A commit was added to the pull request's 'HEAD' branch. Only provided
	//       for pull requests.
	//
	//     cross-referenced
	//       The issue was referenced from another issue. The 'source' attribute
	//       contains the 'id', 'actor', and 'url' of the reference's source.
	//
	//     demilestoned
	//       The issue was removed from a milestone.
	//
	//     head_ref_deleted
	//       The pull request's branch was deleted.
	//
	//     head_ref_restored
	//       The pull request's branch was restored.
	//
	//     labeled
	//       A label was added to the issue.
	//
	//     locked
	//       The issue was locked by the actor.
	//
	//     mentioned
	//       The actor was @mentioned in an issue body.
	//
	//     merged
	//       The issue was merged by the actor. The 'commit_id' attribute is the
	//       SHA1 of the HEAD commit that was merged.
	//
	//     milestoned
	//       The issue was added to a milestone.
	//
	//     referenced
	//       The issue was referenced from a commit message. The 'commit_id'
	//       attribute is the commit SHA1 of where that happened.
	//
	//     renamed
	//       The issue title was changed.
	//
	//     reopened
	//       The issue was reopened by the actor.
	//
	//     reviewed
	//       The pull request was reviewed.
	//
	//     subscribed
	//       The actor subscribed to receive notifications for an issue.
	//
	//     unassigned
	//       The assignee was unassigned from the issue.
	//
	//     unlabeled
	//       A label was removed from the issue.
	//
	//     unlocked
	//       The issue was unlocked by the actor.
	//
	//     unsubscribed
	//       The actor unsubscribed to stop receiving notifications for an issue.
	//
	Event  string  `json:"event"`
	Source *Source `json:"source,omitempty"`
}

// Source represents a reference's source.
type Source struct {
	ID    *int64  `json:"id,omitempty"`
	URL   *string `json:"url,omitempty"`
	Actor *User   `json:"actor,omitempty"`
	Type  *string `json:"type,omitempty"`
	Issue *Issue  `json:"issue,omitempty"`
}
