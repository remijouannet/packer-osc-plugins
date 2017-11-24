package main

import (
	"github.com/hashicorp/packer/packer/plugin"
	"github.com/remijouannet/packer-osc-plugins/builder/osc/ebssurrogate"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(ebssurrogate.Builder))
	server.Serve()
}
