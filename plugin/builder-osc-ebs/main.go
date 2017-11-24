package main

import (
	"github.com/hashicorp/packer/packer/plugin"
	"github.com/remijouannet/packer/packer-osc-plugins/builder/osc/ebs"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(ebs.Builder))
	server.Serve()
}
