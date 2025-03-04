package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	API struct {
		Url string `yaml:"url"`
		Key string `yaml:"key"`
	} `yaml:"api"`
	DB struct {
		Conn  string `yaml:"conn"`
		Table string `yaml:"table"`
	} `yaml:"dao"`
	Kafka struct {
		Brokers string `yaml:"brokers"`
		Topic   string `yaml:"topic"`
	} `yaml:"kafka"`
}

var AppConfig Config

func LoadConfig(configFile string) error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		return err
	}
	return nil
}
