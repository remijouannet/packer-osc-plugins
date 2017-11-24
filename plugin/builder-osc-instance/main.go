package main

import (
	"github.com/remijouannet/packer/packer-osc-plugins/builder/osc/instance"
	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(instance.Builder))
	server.Serve()
}
