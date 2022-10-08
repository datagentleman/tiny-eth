package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type section map[string]json.RawMessage

var conf = map[string]section{}

type Config struct {
	data json.RawMessage
}

func Load(key, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	sec := section{}
	err = json.Unmarshal(data, &sec)
	if err != nil {
		return err
	}

	conf[key] = sec
	return nil
}

func Get(keys ...string) (*Config, error) {
	if len(keys) == 2 {
		section, ok := conf[keys[0]]
		if !ok {
			return nil, errors.New(fmt.Sprintf("Section '%s' doesn't exist", keys[0]))
		}

		key, ok := section[keys[1]]
		if !ok {
			return nil, errors.New(fmt.Sprintf("Section '%s' doesn't have '%s' key", keys[0], keys[1]))
		}

		c := &Config{data: key}
		return c, nil
	}

	return nil, errors.New("You must specify section (ex: 'database') and key (ex: 'prod')")
}

func (c *Config) Decode(src any) error {
	err := json.Unmarshal(c.data, &src)
	if err != nil {
		return err
	}

	return nil
}
