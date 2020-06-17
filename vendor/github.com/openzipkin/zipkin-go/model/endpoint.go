// Copyright 2019 The OpenZipkin Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import "net"

// Endpoint holds the network context of a node in the service graph.
type Endpoint struct {
	ServiceName string `json:"serviceName,omitempty"`
	IPv4        net.IP `json:"ipv4,omitempty"`
	IPv6        net.IP `json:"ipv6,omitempty"`
	Port        uint16 `json:"port,omitempty"`
}

// Empty returns if all Endpoint properties are empty / unspecified.
func (e *Endpoint) Empty() bool {
	return e == nil ||
		(e.ServiceName == "" && e.Port == 0 && len(e.IPv4) == 0 && len(e.IPv6) == 0)
}
