package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		SomeStringValue string
		SomeIntValue    int
		Something       something
	}

	something struct {
		maybe bool
	}
)

func main() {
	configName := "config.toml"

	if _, err := os.Stat(configName); err != nil {
		fmt.Printf("Unable to find %v\n", configName)
	}

    var config Config
    _, err := toml.DecodeFile(configName, &config)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    fmt.Printf("stringValue: %v\nintValue: %v\nmaybe: %v", config.SomeStringValue, config.SomeIntValue, config.Something.maybe)

}
