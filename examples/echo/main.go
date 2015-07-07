package main

import (
	"os"

	"github.com/flier/go.server"
)

func main() {
	server.NewBootstrap("echo").App().Run(os.Args)
}
