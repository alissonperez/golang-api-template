package main

import (
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/infra/ioc"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/net"
)

func main() {
	net.SetupServer(ioc.CreateContainer())
}
