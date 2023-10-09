package github

import (
	"context"
	"github.com/HMUniversity/System/modules/err_handler"
)

func (i *Instance) IsInOrg(org string, username string) (bool, error) {
	member, _, err := i.Cli.Organizations.IsMember(context.Background(), org, username)
	if err != nil {
		err_handler.HandleError(err)
		return false, ErrGetUserInfo
	}
	return member, nil
}
