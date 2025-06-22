// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"

	"github.com/borderssolutions/packer-plugin-vergeio/builder/vergeio"
	vergeioData "github.com/hashicorp/packer-plugin-vergeio/datasource/vergeio"
	vergeioPP "github.com/hashicorp/packer-plugin-vergeio/post-processor/vergeio"
	vergeioProv "github.com/hashicorp/packer-plugin-vergeio/provisioner/vergeio"
	vergeioVersion "github.com/hashicorp/packer-plugin-vergeio/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("my-builder", new(vergeio.Builder))
	pps.RegisterProvisioner("my-provisioner", new(vergeioProv.Provisioner))
	pps.RegisterPostProcessor("my-post-processor", new(vergeioPP.PostProcessor))
	pps.RegisterDatasource("my-datasource", new(vergeioData.Datasource))
	pps.SetVersion(vergeioVersion.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
