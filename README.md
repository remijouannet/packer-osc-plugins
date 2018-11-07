Packer Builders for Outscale (unofficial)
==================

Requirements
------------

-   [Packer](https://www.packer.io/downloads.html) 1.3.2
-   [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Install
---------------------

```
$ wget https://github.com/remijouannet/packer-osc-plugins/releases/download/v0.1/packer-osc-linux_amd64_v0.1.zip
$ unzip packer-osc-linux_amd64_v0.1.zip
$ mkdir ~/.packer.d/plugins/ && mv packer-osc-linux_amd64_v0.1/* ~/.packer.d/plugins/
$ chmod +x ~/.packer.d/plugins/*
```

What's Working
---------------------

For the moment only osc-ebs has been tested you can follow the documentation https://www.packer.io/docs/builders/amazon.html
it's suppose to work the same way the only difference is to not use "custom_endpoints_ec2" but :

```
"endpoints": {"ec2" : "fcu.eu-west-2.outscale.com"},
```

Build without docker
---------------------

Clone repository to: `$GOPATH/src/github.com/remijouannet/packer-osc-plugins`

```
$ mkdir -p $GOPATH/src/github.com/remijouannet; cd $GOPATH/src/github.com/remijouannet
$ git clone git@github.com:remijouannet/packer-osc-plugins
```

Enter the provider directory and build the provider

```
$ cd $GOPATH/src/github.com/remijouannet/packer-osc-plugins
$ make build
```

Build with docker
---------------------

build the binaries, you'll find all the binaries in bin/

```
$ make docker-bin
```
