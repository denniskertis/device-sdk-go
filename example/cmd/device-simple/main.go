// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a simple example of a device service.
package main

import (
	"github.com/denniskertis/device-sdk-go/v42"
	"github.com/denniskertis/device-sdk-go/v42/example/driver"
	"github.com/denniskertis/device-sdk-go/v42/pkg/startup"
)

const (
	serviceName string = "device-simple"
)

func main() {
	sd := driver.SimpleDriver{}
	startup.Bootstrap(serviceName, device.Version, &sd)
}
