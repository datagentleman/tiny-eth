package api

import (
	"fmt"

	"github.com/datagentleman/tiny-eth/pkg/block"
	"github.com/datagentleman/tiny-eth/pkg/common"
	"github.com/datagentleman/tiny-eth/pkg/db"
)

type S struct {
	Params interface{}
}

// Returns the number of most recent block
func BlockNumber(s []interface{}) interface{} {
	hash, _ := db.Get([]byte("LastBlock"))

	h := common.NewHash(hash)
	header, _ := block.FindHeader(h)

	return fmt.Sprintf("0x%x", header.Number)
}
