//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package corekeeper

// key delimiter for edgex keeper
const KeyDelimiter = "/"

// Constants related to defined routes in the v2 service APIs
const ApiBase = "/api/v2"
const ApiKVRoute = ApiBase + "/kv"
const ApiPingRoute = ApiBase + "/ping"

// Constants related to defined url path names and parameters in the v2 service APIs
const (
	Flatten = "flatten"
	Keys    = "keys"
	Raw     = "raw"
)