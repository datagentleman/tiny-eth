//go:build integration
// +build integration

package rpc

import (
	"testing"
)

type Custom []byte

var TestCommands = map[string]cmd{
	"ping": func(data []interface{}) string { return "pong" },
}

func TestStart(t *testing.T) {
	r := New(TestCommands)
	r.Start()
}
