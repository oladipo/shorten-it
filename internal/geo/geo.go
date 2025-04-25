package geo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GeoInfo struct {
	Country string `json:"country"`
	Region  string `json:"regionName"`
	City    string `json:"city"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Query   string `json:"query"`
}

// LookupIP queries ip-api.com for geolocation info about the given IP address.
func LookupIP(ip string) (*GeoInfo, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info GeoInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}
