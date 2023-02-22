package main

import (
	"fmt"
	"github.com/piresrui/orb/config"
	orb2 "github.com/piresrui/orb/orb"
)

func main() {

	conf, _ := config.ProvideConfig()
	orb := orb2.VirtualOrb{
		Config: *conf,
	}

	_ = orb.Status()
	_ = orb.Signup("")
	fmt.Println(conf)
}
