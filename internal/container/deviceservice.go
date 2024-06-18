// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"github.com/denniskertis/go-mod-bootstrap/v42/di"
	"github.com/denniskertis/go-mod-core-contracts/v42/models"

	"github.com/denniskertis/device-sdk-go/v42/pkg/interfaces"
)

// DeviceServiceName contains the name of device service struct in the DIC.
var DeviceServiceName = di.TypeInstanceToName(models.DeviceService{})

// ProtocolDriverName contains the name of protocol driver implementation in the DIC.
var ProtocolDriverName = di.TypeInstanceToName((*interfaces.ProtocolDriver)(nil))

// AutoEventManagerName contains the name of autoevent manager implementation in the DIC
var AutoEventManagerName = di.TypeInstanceToName((*interfaces.AutoEventManager)(nil))

// DeviceServiceFrom helper function queries the DIC and returns device service struct.
func DeviceServiceFrom(get di.Get) *models.DeviceService {
	return get(DeviceServiceName).(*models.DeviceService)
}

// ProtocolDriverFrom helper function queries the DIC and returns protocol driver implementation.
func ProtocolDriverFrom(get di.Get) interfaces.ProtocolDriver {
	return get(ProtocolDriverName).(interfaces.ProtocolDriver)
}

// AutoEventManagerFrom helper function queries the DIC and returns autoevent manager implementation
func AutoEventManagerFrom(get di.Get) interfaces.AutoEventManager {
	return get(AutoEventManagerName).(interfaces.AutoEventManager)
}
