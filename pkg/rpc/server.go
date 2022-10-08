package rpc

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/datagentleman/tiny-eth/pkg/logger"
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

	j, err := json.Marshal(res)
	if err != nil {
		// TODO: Return this to caller ?
		logger.Error(err)
	}

	io.WriteString(w, string(j))
}

func (r *rpc) Start() error {
	http.HandleFunc("/", r.handleRequest)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		return err
	}

	return nil
}

func (r *rpc) RegisterCommand(name string, c Cmd) {
	r.commands[name] = c
}

func (r *rpc) callCommand(req *http.Request) interface{} {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		logger.Error(err)
		return err.Error()
	}

	p := params{}

	err = json.Unmarshal(b, &p)
	if err != nil {
		logger.Error(err)
		return err.Error()
	}

	return r.commands[p.Method](p.Params)
}
