// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"net/http"

	"github.com/denniskertis/go-mod-core-contracts/v42/common"
	"github.com/denniskertis/go-mod-core-contracts/v42/errors"
	"github.com/denniskertis/go-mod-core-contracts/v42/models"

	"github.com/denniskertis/device-sdk-go/v42/internal/autodiscovery"
	"github.com/denniskertis/device-sdk-go/v42/internal/container"

	"github.com/labstack/echo/v4"
)

func (c *RestController) Discovery(e echo.Context) error {
	request := e.Request()
	writer := e.Response()
	ds := container.DeviceServiceFrom(c.dic.Get)
	if ds.AdminState == models.Locked {
		err := errors.NewCommonEdgeX(errors.KindServiceLocked, "service locked", nil)
		return c.sendEdgexError(writer, request, err, common.ApiDiscoveryRoute)
	}

	configuration := container.ConfigurationFrom(c.dic.Get)
	if !configuration.Device.Discovery.Enabled {
		err := errors.NewCommonEdgeX(errors.KindServiceUnavailable, "device discovery disabled", nil)
		return c.sendEdgexError(writer, request, err, common.ApiDiscoveryRoute)
	}

	driver := container.ProtocolDriverFrom(c.dic.Get)

	go autodiscovery.DiscoveryWrapper(driver, c.lc)
	return c.sendResponse(writer, request, common.ApiDiscoveryRoute, nil, http.StatusAccepted)
}
