package client

import (
	"encoding/json"
	"fmt"
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

	// Check if response status code indicates an error
	if resp.StatusCode != http.StatusOK {
		var rqError RQBitError
		if err := json.NewDecoder(resp.Body).Decode(&rqError); err != nil {
			return nil, fmt.Errorf("failed to decode error response: %v", err)
		}
		return nil, &rqError
	}

	var torrents Torrents
	if err := json.NewDecoder(resp.Body).Decode(&torrents); err != nil {
		return nil, fmt.Errorf("failed to decode torrents response: %v", err)
	}

	return &torrents, nil
}
