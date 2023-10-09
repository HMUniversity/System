package config

import "github.com/HMUniversity/System/modules/github"

type Model struct {
	ListenAddr   string          `json:"listen_addr"`
	NeedAuth     bool            `json:"need_auth"`
	GitHub       github.Instance `json:"github"`
	APIKeys      []string        `json:"api_keys"`
	Organisation string          `json:"org"`

	CertTemplate string `json:"cert_template"`
}
