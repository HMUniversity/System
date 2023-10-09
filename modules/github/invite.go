package github

import (
	"context"
	"errors"
	"github.com/HMUniversity/System/modules/err_handler"
	go_github "github.com/google/go-github/v55/github"
)

func (i *Instance) InviteByEmail(email string, org string) error {
	inv, _, err := i.Cli.Organizations.CreateOrgInvitation(context.Background(), org, &go_github.CreateOrgInvitationOptions{
		Email: &email,
	})
	return processInviteErr(inv, err)
}

func (i *Instance) InviteById(id int64, org string) error {
	inv, _, err := i.Cli.Organizations.CreateOrgInvitation(context.Background(), org, &go_github.CreateOrgInvitationOptions{
		InviteeID: &id,
	})
	return processInviteErr(inv, err)

}

var ErrGetUserInfo = errors.New("Cannot get user info")
var ErrInviteUser = errors.New("Cannot invite user")

func (i *Instance) InviteByUsername(username string, org string) error {
	info, _, err := i.Cli.Users.Get(context.Background(), username)
	if err != nil {
		err_handler.HandleError(err)
		return ErrGetUserInfo
	}
	inv, _, err := i.Cli.Organizations.CreateOrgInvitation(context.Background(), org, &go_github.CreateOrgInvitationOptions{
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
