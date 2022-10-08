//go:build integration
// +build integration

package rpc

import (
	"testing"

	"github.com/datagentleman/tiny-eth/pkg/config"
	"github.com/datagentleman/tiny-eth/pkg/db"
	"github.com/datagentleman/tiny-eth/pkg/node/api"
)

var TestCommands = map[string]Cmd{
	"ping":            api.Ping,
	"eth_blockNumber": api.BlockNumber,
}

func TestStart(t *testing.T) {
	conf, _ := config.Get("database", "test")
	db.Configure(conf)
	defer db.Close()

	r := New(TestCommands)
	r.Start()
}
