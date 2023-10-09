package config

import (
	"encoding/json"
	"github.com/KevinZonda/GoX/pkg/iox"
)

var cfg *Model

func Get() *Model {
	return cfg
}

func Load() {
	bs, err := iox.ReadAllByte("config.json")
	if err != nil {
		panic(err)
	}
	cfg = &Model{}
	err = json.Unmarshal(bs, cfg)
	if err != nil {
		panic(err)
	}
}
