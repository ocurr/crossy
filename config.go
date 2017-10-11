package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Profiles map[string][]map[string]options `yaml:",flow"`
}

type options struct {
	Link   string `yaml:",omitempty,flow"`
	Before string `yaml:",omitempty,flow"`
	After  string `yaml:",omitempty,flow"`
}

func NewConfig(fileName string, c *Config) (*Config, error) {

	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	yaml.Unmarshal(buf, c)

	return c, nil
}
