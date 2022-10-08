package config

import (
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	Load("database", "./config_test.json")
	conf, _ := Get("database", "test")

	conf1 := map[string]interface{}{}
	conf2 := map[string]interface{}{"path": "root", "read_only": true}

	conf.Decode(&conf1)

	if !reflect.DeepEqual(conf1, conf2) {
		t.Errorf("config: expected %v got %v", conf2, conf1)
	}
}
