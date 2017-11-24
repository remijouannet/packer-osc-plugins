package main

import (
	"github.com/remijouannet/packer/packer-osc-plugins/builder/osc/chroot"
	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(chroot.Builder))
	server.Serve()
}
