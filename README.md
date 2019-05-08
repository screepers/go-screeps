# go-screeps
A Screeps Library in Go (WIP)

[![Go Report Card](https://goreportcard.com/badge/github.com/screepers/go-screeps)](https://goreportcard.com/report/github.com/screepers/go-screeps)
[![LICENSE](https://img.shields.io/github/license/screepers/go-screeps.svg")](LICENSE)
[![GoDoc](https://godoc.org/github.com/screepers/go-screeps?status.svg)](https://godoc.org/github.com/screepers/go-screeps)

## Basic Usage

Import and use like any other go module

`import "github.com/screepers/go-screeps/client"`

Simple example:

```Go
package main

import (
	"log"

	"github.com/screepers/go-screeps/config"
	"github.com/screepers/go-screeps/screeps"
)

func main() {
	conf := config.NewConfig()
	server := conf.Servers["main"]
	c := screeps.NewClient(server)

	resp, err := c.AuthMe()
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("%v", resp)
	}
}
```
