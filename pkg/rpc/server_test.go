//go:build integration
// +build integration

package rpc

import (
	"testing"
)

type Custom []byte

func TestStart(t *testing.T) {
	r := New()

	r.RegisterCommand("ping", func(data []interface{}) string {
		return "pong"
	})

	r.Start()
}
