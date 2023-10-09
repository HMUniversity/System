package config

import (
	"encoding/json"
	"github.com/KevinZonda/GoX/pkg/iox"
)

var cfg *ConfigModel

func Get() *ConfigModel {
	return cfg
}

func Load() {
	bs, err := iox.ReadAllByte("config.json")
	if err != nil {
		panic(err)
	}
	cfg = &ConfigModel{}
	err = json.Unmarshal(bs, cfg)
	if err != nil {
		panic(err)
	}
}
