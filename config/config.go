package config

import (
	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
)

var mutex = sync.RWMutex{}
var config *Configuration

const (
	DefaultPath string = "/Users/macbook/Documents/config.yml"
)

type Configuration struct {
	path string

	Mongo mongoConfiguration `yaml:"mongo"`
}

type mongoConfiguration struct {
	Host     string `yaml:"host" default:"localhost"`
	Port     int    `yaml:"port" default:"27017"`
	Username string `yaml:"username" default:"admin1"`
	Password string `yaml:"password" default:"123"`
	Database string `yaml:"database" default:"backend-app"`
}

func Setup() error {
	c, err := NewWithPath(DefaultPath)

	if err != nil {
		return err
	}

	if _, err = os.Stat(c.path); os.IsNotExist(err) {
		// File does not exist
		if err = c.WriteToFile(); err != nil {
			return err
		}
	} else {
		if err = ReadFromFile(c); err != nil {
			return err
		}
	}

	config = c

	return nil
}

func ReadFromFile(base *Configuration) error {
	mutex.RLock()
	defer mutex.RUnlock()

	path := base.path

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, base)
	if err != nil {
		return err
	}

	return nil
}

func NewWithPath(path string) (*Configuration, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var c Configuration

	if err := defaults.Set(&c); err != nil {
		return nil, err
	}

	c.path = path
	return &c, nil
}

func (c Configuration) WriteToFile() error {
	mutex.Lock()
	defer mutex.Unlock()

	s, _ := yaml.Marshal(c)

	if err := ioutil.WriteFile(c.path, s, 0666); err != nil {
		return err
	}

	return nil
}

func Get() Configuration {
	mutex.RLock()
	defer mutex.RUnlock()


	c := *config
	return c
}

func Update(update func(configuration Configuration) *Configuration) error {
	cfg := update(*config)
	config = cfg

	if err := config.WriteToFile(); err != nil {
		return err
	}

	return nil
}