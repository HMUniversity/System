package config

import "github.com/HMUniversity/System/modules/github"

type ConfigModel struct {
	GitHub       github.Instance `json:"github"`
	APIKeys      []string        `json:"api_keys"`
	Organisation string          `json:"org"`
}
