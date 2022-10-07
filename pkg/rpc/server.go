package rpc

import (
	"encoding/json"
	"io"
	"net/http"
)

type Cmd func(data []interface{}) interface{}

type rpc struct {
	commands map[string]Cmd
}

type params struct {
	Method string
	Params []interface{}
	Id     int
}

type response struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Result  interface{} `json:"result"`
}

func New(commands map[string]Cmd) *rpc {
	if commands == nil {
		commands = map[string]Cmd{}
	}

	return &rpc{commands: commands}
}

func (r *rpc) handleRequest(w http.ResponseWriter, req *http.Request) {
	data := r.callCommand(req)
	res := response{Result: data}

	j, _ := json.Marshal(res)
	io.WriteString(w, string(j))
}

func (r *rpc) Start() {
	http.HandleFunc("/", r.handleRequest)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}

func (r *rpc) RegisterCommand(name string, c Cmd) {
	r.commands[name] = c
}

func (r *rpc) callCommand(req *http.Request) interface{} {
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
