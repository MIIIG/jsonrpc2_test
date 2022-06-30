package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Result struct {
	Id      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}
type Request struct {
	Id      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}

type Params struct {
	Name string `json:"name"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var req Request
	// var res Result
	fmt.Printf("server: %s /\n", r.Method)
	fmt.Printf("header: %s /\n", r.Header)
	fmt.Printf("body: %s /\n", r.Body)
	switch r.Method {
	case "GET":
		w.WriteHeader(200)
	case "POST":
		reqbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("server: could not read request body: %s\n", err)
		}
		fmt.Printf("request body: %s\n", reqbody)
		jbody := json.Unmarshal(reqbody, &req)
		if jbody != nil {
			panic(jbody)
		}
		fmt.Println(req)
		// body := json.NewDecoder(r.Body)
		// if err != nil {
		// 	fmt.Sprintf("Bad mashalling: %s", err)
		// }
		// fmt.Printf("Marshalling request: %s", request)
		// err := json.NewDecoder(r.Body).Decode(&req)
		// dec := json.NewDecoder(r.Body)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// fmt.Printf("Request: %s\n", dec)

		fmt.Printf("its Show time\n")
		io.WriteString(w, "This is RPC server!\n")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Mission Impossible!")
	}
}

// func Greeting(ctx context.Context) Result {
// 	r := fmt.Sprintf("Hello, %s\n", ctx.name)
// 	return Result{
// 		result: r,
// 	}
// }

func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
