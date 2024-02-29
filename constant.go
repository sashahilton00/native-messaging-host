// constant.go - All constants in a file.
// Copyright (c) 2018 - 2024  Sasha Hilton <sashahilton00@users.noreply.github.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package host

// The Http connection and timeout configurations.
const (
	HttpContinueTimeout   = 5
	HttpKeepAlive         = 600
	HttpDialTimeout       = 10
	HttpOverallTimeout    = 15
	IdleTimeout           = 90
	MaxConnections        = 100
	ResponseHeaderTimeout = 10
	TLSDialTimeout        = 15
)
