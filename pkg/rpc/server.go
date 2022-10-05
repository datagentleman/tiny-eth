package rpc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type cmd func(data []interface{}) string

type rpc struct {
	commands map[string]cmd
}

type params struct {
	Method string
	Params []interface{}
	Id     int
}

func New() *rpc {
	return &rpc{commands: map[string]cmd{}}
}

func (r *rpc) handleRequest(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, r.callCommand(req))
}

func (r *rpc) Start() {
	fmt.Println("\033[36m", logo, "\033[0m")
	http.HandleFunc("/", r.handleRequest)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}

func (r *rpc) RegisterCommand(name string, c cmd) {
	r.commands[name] = c
}

func (r *rpc) callCommand(req *http.Request) string {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	p := params{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		return err.Error()
	}

	return r.commands[p.Method](p.Params)
}
