package node

import (
	"fmt"

	"github.com/datagentleman/tiny-eth/pkg/node/api"
	"github.com/datagentleman/tiny-eth/pkg/rpc"
)

var logo = `
   __      __                            ________  ________  __    __ 
  /  |    /  |                          /        |/        |/  |  /  |
 _$$ |_   $$/  _______   __    __       $$$$$$$$/ $$$$$$$$/ $$ |  $$ |
/ $$   |  /  |/       \ /  |  /  |      $$ |__       $$ |   $$ |__$$ |
$$$$$$/   $$ |$$$$$$$  |$$ |  $$ |      $$    |      $$ |   $$    $$ |
  $$ | __ $$ |$$ |  $$ |$$ |  $$ |      $$$$$/       $$ |   $$$$$$$$ |
  $$ |/  |$$ |$$ |  $$ |$$ \__$$ |      $$ |_____    $$ |   $$ |  $$ |
  $$  $$/ $$ |$$ |  $$ |$$    $$ |      $$       |   $$ |   $$ |  $$ |
   $$$$/  $$/ $$/   $$/  $$$$$$$ |      $$$$$$$$/    $$/    $$/   $$/ 
                        /  \__$$ |
                        $$    $$/                                     
                         $$$$$$/                                      
`

type Node struct{}

var commands = map[string]rpc.Cmd{
	"ping":            api.Ping,
	"eth_blockNumber": api.BlockNumber,
}

func New() *Node {
	return &Node{}
}

func (n *Node) Start() {
	fmt.Println("\033[36m", logo, "\033[0m")

	rpc.New(commands).Start()
}
