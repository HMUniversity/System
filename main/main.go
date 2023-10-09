package main

import (
	"github.com/HMUniversity/System/modules/config"
)

func main() {
	config.Load()
	config.Get().GitHub.Init()
}
