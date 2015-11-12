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
	"fmt"
	"github.com/marpaia/graphite-golang"
)

func sendMetrics(status []Status, config Configuration) {
	var Graphite, err = graphite.NewGraphite(config.MetricsHost, config.MetricsPort)
	if err != nil {
		log.Error("Can't connect to graphite collector: %v", err)
		return
	}

	for _, entry := range status {
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".queued_requests"), entry.QCur)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".current_sessions"), entry.SCur)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".total_connections"), entry.STot)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".bytes_in"), entry.BytesIn)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".bytes_out"), entry.BytesOut)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".errors_request"), entry.EReq)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".errors_connecting"), entry.Econ)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".errors_response"), entry.EResp)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".checks_failed"), entry.Chkfail)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".switches_to_down"), entry.ChkDown)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".total_downtime"), entry.Downtime)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".checks_duration"), entry.ChkDur)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".hrsp_1xx"), entry.HRSP1)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".hrsp_2xx"), entry.HRSP2)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".hrsp_3xx"), entry.HRSP3)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".hrsp_4xx"), entry.HRSP4)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".hrsp_5xx"), entry.HRSP5)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".hrsp_other"), entry.HRSPO)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".total_requests"), entry.ReqTot)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".time_queue"), entry.QTime)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".time_connect"), entry.CTime)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".time_response"), entry.RTime)
		Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".haproxy.", entry.Type, ".", entry.Name, ".time_total_session"), entry.TTime)
	}
}
