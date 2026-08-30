package main

import (
	"maps"
	"os"

	"github.com/subosito/gotenv"
)

type Config map[string]string

func ReadConfig(filename string) (Config, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return make(Config), nil
	}
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	config := make(Config)
	maps.Copy(config, gotenv.Parse(fd))
	return config, nil
}
