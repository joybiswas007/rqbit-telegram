package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// get the torrents list
func (rqb *RQBitConfig) GetTorrents() (*Torrents, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s:%d/torrents", rqb.IpAddress, rqb.Port), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var torrents Torrents
	if err := json.Unmarshal(bodyText, &torrents); err != nil {
		return nil, err
	}
	return &torrents, nil
}
