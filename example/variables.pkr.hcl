# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

locals {
  foo = data.vergeio-my-datasource.mock-data.foo
  bar = data.vergeio-my-datasource.mock-data.bar
}