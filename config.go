/*
   conntrack-logger
   Copyright (C) 2015 Denis V Chapligin <akashihi@gmail.com>
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"flag"
)

type Configuration struct {
	StatusUrl     string
	Period        int
	MetricsHost   string
	MetricsPort   int
	MetricsPrefix string
}

func config() Configuration {
	statusUrlPtr := flag.String("url", "http://127.0.0.1/server_status", "Url of the haproxy csv status")
	periodPtr := flag.Int("period", 60, "Period of status page polling in seconds")
	metricsHostPtr := flag.String("metrics-host", "127.0.0.1", "Ip address of the Graphite collector host")
	metricsPortPtr := flag.Int("metrics-port", 2003, "Port of the Graphite collector host")
	metricsPrefixPtr := flag.String("metrics-prefix", "server.nginx", "Prefix to prepend metrics names")

	flag.Parse()

	return Configuration{StatusUrl: *statusUrlPtr, Period: *periodPtr, MetricsHost: *metricsHostPtr, MetricsPort: *metricsPortPtr, MetricsPrefix: *metricsPrefixPtr}
}
