package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// get information about specific torrent
func (rqb *RQBitConfig) GetTorrentInfo(tID int) (*TorrentDetails, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s:%d/torrents/%d", rqb.IpAddress, rqb.Port, tID), nil)
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

	var details TorrentDetails
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return nil, fmt.Errorf("failed to decode torrents response: %v", err)
	}

	return &details, nil
}
