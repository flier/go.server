package main

import (
	"github.com/flier/go.server"
)

func main() {
	server.Bootstrap("echo").Create()
}
