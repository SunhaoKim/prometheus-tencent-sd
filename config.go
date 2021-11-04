package main

import (
	"time"
)

type Config struct {
	Ak       string        `yaml:"ak"`
	Sk       string        `yaml:"sk"`
	Region   string        `yaml:"region"`
	Port     int           `yaml:"port"`
	Interval time.Duration `yaml:"interval"`
	Filters  []Filter      `yaml:"filters"`
}

type Filter struct {
	Name   string   `yaml:"name"`
	Values []string `yaml:"values"`
}
