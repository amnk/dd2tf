//go:generate go-bindata -o tpl.go tmpl

package main

import (
	"errors"
	"strconv"
	"github.com/zorkian/go-datadog-api"
)

type Monitor struct {
}

func (m Monitor) getElement(client datadog.Client, id interface{}) (interface{}, error) {
	var idInt int
	switch v := id.(type) {
	case string:
		var err error
		idInt, err = strconv.Atoi(v)
		if (err != nil) {
			return "", err
		}
	case int:
		idInt = v
	default:
		return "", errors.New("unsupported id type, should be string or int")
	}
	mon, err := client.GetMonitor(idInt)
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
		ids = append(ids, Item{id: *elem.Id, d: Monitor{}})
	}
	return ids, nil
}
