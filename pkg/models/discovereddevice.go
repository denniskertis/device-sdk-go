//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

import "github.com/denniskertis/go-mod-core-contracts/v42/models"

// DiscoveredDevice defines the required information for a found device.
type DiscoveredDevice struct {
	Name        string
	Protocols   map[string]models.ProtocolProperties
	Description string
	Labels      []string
}
