package github

import "github.com/google/go-github/v55/github"

type GitHubObj struct {
	Token    string         `json:"token"`
	Username string         `json:"username"`
	Cli      *github.Client `json:"-"`
}

func (g *GitHubObj) Init() {
	g.Cli = github.NewClient(nil).WithAuthToken(g.Token)
}
