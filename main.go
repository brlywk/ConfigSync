package main

import (
	"brlywk/ConfigSync/config"
	"fmt"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
)

// Nobody knows what this function really does...
func main() {
	configName := "config.toml"

	if _, err := os.Stat(configName); err != nil {
		fmt.Printf("Unable to find %v\n", configName)
	}

	var config config.Config
	_, err := toml.DecodeFile(configName, &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	c := reflect.TypeOf(&config).Elem()

	for i := 0; i < c.NumField(); i++ {
		field := c.Field(i)
		value := reflect.ValueOf(&config).Elem().Field(i).Interface()
		fmt.Printf("Field: %v\tValue: %v\n", field, value)
	}

}
