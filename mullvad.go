package mullvad

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MullvadType struct {
	Ipaddr     string  `json:"ip"`
	Country    string  `json:"country"`
	City       string  `json:"city"`
	Long       float32 `json:"longitude"`
	Lat        float32 `json:"latitude"`
	Exitip     bool    `json:"mullvad_exit_ip"`
	Exithost   string  `json:"mullvad_exit_ip_hostname"`
	Servertype string  `json:"mullvad_server_type"`
	Org        string  `json:"organization"`

	Blacklisted struct {
		Blacklisted bool        `json:"blacklisted"`
		Results     interface{} `json:"results"`
	} `json:"blacklisted"`
}

var url = "https://am.i.mullvad.net/json"

func (m *MullvadType) Amimullvad() bool	{
	return m.Exitip
}

func GetMullvad() (*MullvadType, error) {

	var result MullvadType
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return &result, fmt.Errorf("%s", err)
	}
	return &result, nil
}
