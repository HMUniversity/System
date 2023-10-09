package github_invitor

import (
	"context"
	"errors"
	"github.com/HMUniversity/System/main/modules/err_handler"
	"github.com/HMUniversity/System/main/modules/github"
	go_github "github.com/google/go-github/v55/github"
)

type GitHubInvitor struct {
	GitHub *github.GitHubObj
}

func (i *GitHubInvitor) InviteByEmail(email string, org string) error {
	inv, _, err := i.GitHub.Cli.Organizations.CreateOrgInvitation(context.Background(), org, &go_github.CreateOrgInvitationOptions{
		Email: &email,
	})
	return processInviteErr(inv, err)
}

func (i *GitHubInvitor) InviteById(id int64, org string) error {
	inv, _, err := i.GitHub.Cli.Organizations.CreateOrgInvitation(context.Background(), org, &go_github.CreateOrgInvitationOptions{
		InviteeID: &id,
	})
	return processInviteErr(inv, err)

}

var ErrGetUserInfo = errors.New("Cannot get user info")
var ErrInviteUser = errors.New("Cannot invite user")

func (i *GitHubInvitor) InviteByUsername(username string, org string) error {
	info, _, err := i.GitHub.Cli.Users.Get(context.Background(), username)
	if err != nil {
		err_handler.HandleError(err)
		return ErrGetUserInfo
	}
	inv, _, err := i.GitHub.Cli.Organizations.CreateOrgInvitation(context.Background(), org, &go_github.CreateOrgInvitationOptions{
		InviteeID: info.ID,
	})
	return processInviteErr(inv, err)
}

func processInviteErr(inv *go_github.Invitation, err error) error {
	if err != nil {
		err_handler.HandleError(err)
		return ErrInviteUser
	}
	return nil
}
