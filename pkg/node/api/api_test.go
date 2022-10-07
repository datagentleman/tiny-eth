package api

import (
	"testing"

	"github.com/datagentleman/tiny-eth/pkg/config"
	"github.com/datagentleman/tiny-eth/pkg/db"
)

func TestFindHeader(t *testing.T) {
	conf, _ := config.Load("../../../config/database.json")
	db.Configure(conf)

	BlockNumber(nil)
}
