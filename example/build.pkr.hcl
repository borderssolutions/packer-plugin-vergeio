# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

packer {
  required_plugins {
    vergeio = {
      version = ">=v0.1.0"
      source  = "github.com/hashicorp/vergeio"
    }
  }
}

source "vergeio-my-builder" "foo-example" {
  mock = local.foo
}

source "vergeio-my-builder" "bar-example" {
  mock = local.bar
}

build {
  sources = [
    "source.vergeio-my-builder.foo-example",
  ]

  source "source.vergeio-my-builder.bar-example" {
    name = "bar"
  }

  provisioner "vergeio-my-provisioner" {
    only = ["vergeio-my-builder.foo-example"]
    mock = "foo: ${local.foo}"
  }

  provisioner "vergeio-my-provisioner" {
    only = ["vergeio-my-builder.bar"]
    mock = "bar: ${local.bar}"
  }

  post-processor "vergeio-my-post-processor" {
    mock = "post-processor mock-config"
  }
}
