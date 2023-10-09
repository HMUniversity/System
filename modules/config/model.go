package config

import "github.com/HMUniversity/System/modules/github"

type ConfigModel struct {
	ListenAddr   string          `json:"listen_addr"`
	GitHub       github.Instance `json:"github"`
	APIKeys      []string        `json:"api_keys"`
	Organisation string          `json:"org"`
}
