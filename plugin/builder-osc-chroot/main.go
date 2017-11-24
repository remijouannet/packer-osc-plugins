package main

import (
	"github.com/hashicorp/packer/packer/plugin"
	"github.com/remijouannet/packer-osc-plugins/builder/osc/chroot"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(chroot.Builder))
	server.Serve()
}
