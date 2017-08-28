// Copyright 2016 The go-daylight Authors
// This file is part of the go-daylight library.
//
// The go-daylight library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-daylight library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-daylight library. If not, see <http://www.gnu.org/licenses/>.

package controllers

import (
	"strconv"

	"github.com/EGaaS/go-egaas-mvp/packages/model"
	"github.com/EGaaS/go-egaas-mvp/packages/utils"
)

type listOfTablesPage struct {
	Alert     string
	Lang      map[string]string
	WalletID  int64
	CitizenID int64
	TxType    string
	TxTypeID  int64
	TimeNow   int64
	Global    string
	Tables    []map[string]string
}

// ListOfTables show the list of custom tables
func (c *Controller) ListOfTables() (string, error) {

	var err error

	global := c.r.FormValue("global")
	prefix := c.StateIDStr
	if global == "1" {
		prefix = "global"
	}
	t := &model.Table{}
	t.SetTablePrefix(prefix)
	data, err := t.GetAll(prefix)
	if err != nil {
		return "", utils.ErrInfo(err)
	}
	tables := make([]map[string]string, 0)
	for _, table := range data {
		mappedTable := table.ToMap()
		count, err := model.GetRecordsCount(string(table.Name))
		if err != nil {
			return "", utils.ErrInfo(err)
		}
		mappedTable["count"] = strconv.FormatInt(count, 10)
		tables = append(tables, mappedTable)
	}

	TemplateStr, err := makeTemplate("list_of_tables", "listOfTables", &listOfTablesPage{
		Alert:     c.Alert,
		Lang:      c.Lang,
		Global:    global,
		WalletID:  c.SessWalletID,
		CitizenID: c.SessCitizenID,
		Tables:    tables})
	if err != nil {
		return "", utils.ErrInfo(err)
	}
	return TemplateStr, nil
}
