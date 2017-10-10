package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Profiles map[string]profile `yaml:",flow"`
}

type profile struct {
	Link []map[string]options `yaml:",omitempty,flow"`
}

type options struct {
	Before string `yaml:",omitempty,flow"`
	After  string `yaml:",omitempty,flow"`
}

func NewConfig(fileName string) (*Config, error) {

	var c Config

	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	yaml.Unmarshal(buf, &c)

	return &c, nil
}
