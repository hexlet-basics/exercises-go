package main

import (
	"fmt"
)

// BEGIN

type Config struct {
	Host *string
	Port int
}

func PrintConfig(cfg *Config) {
	if cfg.Host == nil {
		fmt.Println("Host is not set")
	} else {
		fmt.Println("Host:", *cfg.Host)
	}
	fmt.Println("Port:", cfg.Port)
}

// END
