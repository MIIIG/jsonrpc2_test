package main

import (
	"fmt"
	"net/rpc"

	"go.neonxp.dev/jsonrpc2/rpc"
)

type Result struct {
	id     string
	result string
}

type Params struct {
	id   string
	name string
}

func Greeting(params Params) Result {
	r_id := params.id
	r := fmt.Sprintf("Hello, %s\n", params.name)
	return Result{
		result: r,
		id:     r_id,
	}
}

func main() {
	s := rpc.NewServer()
}
