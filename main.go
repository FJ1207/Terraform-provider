package main

import (
	"context"
	"flag"
	//"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"log"
	demo "terraform-provider-yunjidft/demo"
)

func main() {

	var debugMode bool

	// Set this flag to true if you want the provider to run in debug mode. Leaving it as is will cause it to run
	// normally.
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        debugMode,
		ProviderAddr: "registry.terraform.io/dell/redfish",
		ProviderFunc: demo.Provider,
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/dell/redfish", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
	//plugin.Serve(&plugin.ServeOpts{
	//	ProviderFunc: demo.Provider,
	//})
}
