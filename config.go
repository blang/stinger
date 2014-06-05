package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	MantisServices []*MantisConfig `json:"mantis"`
}

type MantisConfig struct {
	Host string `json:"host"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

func ReadConfigFromFile(filename string) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
