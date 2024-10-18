package main

import (
	"context"
	"flag"
	"github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"log"
)

var (
	version string = "dev"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/Dodai-Dodai/proxmox-sdn",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
