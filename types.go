package netboxapi

type Result struct {
	Count    int           `json:"count"`
	Next     int           `json:"next"`
	Previous int           `json:"previous"`
	Results  []interface{} `json:"results"`
}
