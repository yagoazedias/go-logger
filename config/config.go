package config

import (
	"encoding/json"
	"fmt"
	"github.com/yagoazedias/go-logger/logger"
	"io/ioutil"
	"os"
)

type Settings struct {
	Prefix string
	AppName string
}

type Config struct {
	Keys map[string]string
}

func BuildFromJSON(settings Settings) (c *Config, log *logger.Log) {
	log = logger.NewLogger(settings.AppName)
	c = &Config{}
	jsonFile, err := ioutil.ReadFile("env.json")
	err = json.Unmarshal(jsonFile, &c.Keys)

	if err != nil {
		log.Error(fmt.Sprintf("Could not load json config file, check your 'env.json' file"), err)
	}

	return c, log
}

func Build(settings Settings, keys []string) (c *Config, log *logger.Log) {
	c = &Config{}
	log = logger.NewLogger(settings.AppName)
	c.Keys = make(map[string]string)

	for _, key := range keys {
			c.Keys[key] = os.Getenv(key)
			log.Info(fmt.Sprintf("[CONFIG] Getting key value %s from os env", os.Getenv(key)))
	}

	c.Dumps(settings.AppName)

	return c, log
}

func (c Config) Dumps(appName string) {
	data, err := json.Marshal(c.Keys)
	log := logger.NewLogger(appName)

	if err != nil {
		panic(err)
	}

	log.Info(fmt.Sprintf("[CONFIG] %s", string(data)))
}
