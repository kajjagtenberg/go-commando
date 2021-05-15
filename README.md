# go-commando

Simple package for command handling pipeline in Go.

This code was lifted verbatim from my project [EventflowDB](https://github.com/kajjagtenberg/eventflowdb) since I found it a clean way to integrate a command pipeline into an application without introducing direct coupling with the transport protocol. This allows you to write commands once, and a port for different transports once. Then changes to commands need to only be made in a single place.

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
		Name: "add", // "a" would be sufficient too, since it's the shorthand version for the same command
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
