# go-commando

Simple package for command handling pipeline in Go.

## Installation

```shell
go get -u github.com/kajjagtenberg/go-commando
```

## Usage

```go
package main

import (
	"encoding/json"
	"log"

	"github.com/kajjagtenberg/go-commando"
)

type AddRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

func AddHandler(cmd commando.Command) (interface{}, error) {
	var req AddRequest

	if err := json.Unmarshal(cmd.Args, &req); err != nil {
		return nil, err
	}

	return req.A + req.B, nil
}

func main() {
	// Registering command handlers
	dispatcher := commando.NewCommandDispatcher()
	dispatcher.Register("add", "a", AddHandler) // Name is "add", shorthand version is "a"

	// Executing a command
	req := AddRequest{
		A: 3,
		B: 5,
	}
	args, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	cmd := commando.Command{
		Name: "add",
		Args: args,
	}

	result, err := dispatcher.Handle(cmd)
	if err != nil {
		panic(err)
	}

	sum, ok := result.(int)
	if !ok {
		panic("Invalid cast")
	}

	log.Println(sum) // Returns 8
}
```

## License

This project is licensed under the MIT license - Copyright (c) 2021 Kaj Jagtenberg. See the [LICENSE.md](LICENSE.md) file for details
