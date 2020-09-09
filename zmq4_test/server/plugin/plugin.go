package plugin

import (
	"encoding/json"
	"fmt"
)

var Plugins []*Plugin

type Plugin struct {
	ID     string `json:"id"`
	Host   string `json:"host"`
	Desc   string `json:"desc"`
	Signal int    `json:"signal"`
	API    []API  `json:"api"`
}

type API struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Method []string `json:"method"`
	Path   string   `json:"path"`
}

func Umarshal(content string) (*Plugin, error) {
	plugin := new(Plugin)
	err := json.Unmarshal([]byte(content), &plugin)
	if err != nil {
		return nil, err
	}
	return plugin, err
}

func (p *Plugin) Valid() bool {
	if len(p.API) == 0 || p.Host == "" {
		return false
	}

	return true
}

func (p *Plugin) String() string {
	return fmt.Sprintf("%v", p.API)
}

func (p *Plugin) Append() {
	Plugins = append(Plugins, p)
}
