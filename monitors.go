//go:generate go-bindata -o tpl.go tmpl

package main

import (
	"github.com/zorkian/go-datadog-api"
	"strconv"
)

type Monitor struct {
}

func (m Monitor) getElement(client datadog.Client, id string) (interface{}, error) {
	monitorID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	mon, err := client.GetMonitor(monitorID)
	return mon, err
}

func (m Monitor) getAsset() string {
	return "tmpl/monitor.tmpl"
}

func (m Monitor) getName() string {
	return "monitors"
}

func (m Monitor) String() string {
	return m.getName()
}

func (m Monitor) getAllElements(client datadog.Client) ([]Item, error) {
	var ids []Item
	monitors, err := client.GetMonitors()
	if err != nil {
		return nil, err
	}
	for _, elem := range monitors {
		ids = append(ids, Item{id: strconv.Itoa(*elem.Id), d: Monitor{}})
	}
	return ids, nil
}
