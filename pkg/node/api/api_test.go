package api

import (
	"testing"

	"github.com/datagentleman/tiny-eth/pkg/config"
	"github.com/datagentleman/tiny-eth/pkg/db"
)

func TestFindHeader(t *testing.T) {
	config.Load("database", "../../../config/database.json")
	conf, _ := config.Get("database", "test")

	db.Configure(conf)
	BlockNumber(nil)
}
