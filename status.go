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
	"encoding/csv"
	"io"
	"strings"
)

type Status struct {
	Type        string
	Name        string
	QCur        string
	SCur        string
	STot        string
	BytesIn     string
	BytesOut    string
	EReq        string
	Econ        string
	EResp       string
	Chkfail     string
	ChkDown     string
	Downtime    string
	ChkDur      string
	HRSP1       string
	HRSP2       string
	HRSP3       string
	HRSP4       string
	HRSP5       string
	HRSPO       string
	ReqTot      string
	QTime       string
	CTime       string
	RTime       string
	TTime       string
	MaxReqRate  string
	CurrReqRate string
}

func parse(page io.ReadCloser) ([]Status, error) {
	var result = []Status{}

	r := csv.NewReader(page)
	records, err := r.ReadAll()
	page.Close()
	if err != nil {
		log.Error("Can't parse status page: %v", err)
		return result, err
	}
	for _, entry := range records {
		var item = Status{}
		item.Type = entry[32]
		item.Name = strings.Join(entry[0:2], "-")

		item.SCur = entry[4]
		item.STot = entry[7]
		item.BytesIn = entry[8]
		item.BytesOut = entry[9]

		switch item.Type { //Type
		case "type": //Header
			continue
		case "0": //Frontend
			item.Type = "Frontend"
			item.EReq = entry[12]
			item.HRSP1 = entry[39]
			item.HRSP2 = entry[40]
			item.HRSP3 = entry[41]
			item.HRSP4 = entry[42]
			item.HRSP5 = entry[43]
			item.HRSPO = entry[44]
			item.CurrReqRate = entry[46]
			item.MaxReqRate = entry[47]
			item.ReqTot = entry[48]
			break
		case "1": //Backend
			item.Type = "Backend"
			item.QCur = entry[2]
			item.Econ = entry[13]
			item.EResp = entry[14]
			item.ChkDown = entry[22]
			item.Downtime = entry[24]
			item.HRSP1 = entry[39]
			item.HRSP2 = entry[40]
			item.HRSP3 = entry[41]
			item.HRSP4 = entry[42]
			item.HRSP5 = entry[43]
			item.HRSPO = entry[44]
			item.QTime = entry[58]
			item.CTime = entry[59]
			item.RTime = entry[60]
			item.TTime = entry[61]
			break
		case "2": //Server
			item.Type = "Server"
			item.QCur = entry[2]
			item.Econ = entry[13]
			item.EResp = entry[14]
			item.Chkfail = entry[21]
			item.ChkDown = entry[22]
			item.Downtime = entry[24]
			item.HRSP1 = entry[39]
			item.HRSP2 = entry[40]
			item.HRSP3 = entry[41]
			item.HRSP4 = entry[42]
			item.HRSP5 = entry[43]
			item.HRSPO = entry[44]
			item.QTime = entry[58]
			item.CTime = entry[59]
			item.RTime = entry[60]
			item.TTime = entry[61]
			break
		case "3": //Listener
			item.Type = "Listener"
			item.EReq = entry[12]
			break
		}

		result = append(result, item)
	}

	return result, nil
}
