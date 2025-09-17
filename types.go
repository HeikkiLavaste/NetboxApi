package netboxapi

import "time"

type Result struct {
	Count    int              `json:"count"`
	Next     int              `json:"next"`
	Previous int              `json:"previous"`
	Results  []map[string]any `json:"results"`
}

type L2VPN struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID         int    `json:"id"`
		URL        string `json:"url"`
		DisplayURL string `json:"display_url"`
		Display    string `json:"display"`
		Identifier int    `json:"identifier"`
		Name       string `json:"name"`
		Slug       string `json:"slug"`
		Type       struct {
			Value string `json:"value"`
			Label string `json:"label"`
		} `json:"type"`
		Status struct {
			Value string `json:"value"`
			Label string `json:"label"`
		} `json:"status"`
		ImportTargets []interface{} `json:"import_targets"`
		ExportTargets []interface{} `json:"export_targets"`
		Description   string        `json:"description"`
		Comments      string        `json:"comments"`
		Tenant        interface{}   `json:"tenant"`
		Tags          []Tags        `json:"tags"`
		CustomFields  struct {
			VPRN   string `json:"VPRN"`
			Routed bool   `json:"routed"`
			Subnet string `json:"subnet"`
			Vlan   int    `json:"Vlan"`
		} `json:"custom_fields"`
		Created     time.Time `json:"created"`
		LastUpdated time.Time `json:"last_updated"`
	} `json:"results"`
}

type Tags struct {
	Color      string `json:"color"`
	Display    string `json:"display"`
	DisplayURL string `json:"display_url"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Url        string `json:"url"`
}
