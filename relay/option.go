/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2017 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package relay

import (
	"net"

	"github.com/intelsdi-x/snap-relay/protocol"
)

type Option func(rm relayMetrics) Option

func UDPConnectionOption(conn *net.UDPConn) Option {
	return func(rm relayMetrics) Option {
		if rm.isStarted {
			log.WithFields(log.Fields{
				"_block": "UDPConnectionOption",
			}).Warn("option cannot be set.  service already started")
			return UDPConnectionOption(nil)
		}
		rm.udp = protocol.NewUDPListener(protocol.UDPConnectionOption(conn))
		return UDPConnectionOption(conn)
	}
}

func UDPListenPortOption(port *int) Option {
	return func(rm relayMetrics) Option {
		if rm.isStarted {
			log.WithFields(log.Fields{
				"_block": "UDPListenPortOption",
				"detail": "service already started",
			}).Warn("option cannot be set")
			return UDPListenPortOption(port)
		}
		rm.udp = protocol.NewUDPListener(protocol.UDPListenPortOption(port))
		return UDPListenPortOption(port)
	}
}

func TCPListenPortOption(port *int) Option {
	return func(rm relayMetrics) Option {
		if rm.isStarted {
			log.WithFields(log.Fields{
				"_block": "TCPListenPortOption",
				"detail": "service already started",
			}).Warn("option cannot be set")
			return TCPListenPortOption(port)
		}
		rm.tcp = protocol.NewTCPListener(protocol.TCPListenPortOption(port))
		return TCPListenPortOption(port)
	}
}

func TCPListenerOption(conn *net.TCPListener) Option {
	return func(rm relayMetrics) Option {
		if rm.isStarted {
			log.WithFields(log.Fields{
				"_block": "TCPConnectionOption",
			}).Warn("option cannot be set.  service already started")
			return TCPListenerOption(nil)
		}
		rm.tcp = protocol.NewTCPListener(protocol.TCPListenerOption(conn))
		return TCPListenerOption(conn)
	}
}
