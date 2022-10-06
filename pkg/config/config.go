package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Load(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	conf := map[string]interface{}{}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
