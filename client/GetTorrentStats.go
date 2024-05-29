package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// pass torrent id and returns specific torrents details
func (rqb *RQBitConfig) GetTorrentStats(tID int) (*TorrentStats, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s:%d/torrents/%d/stats/v1", rqb.IpAddress, rqb.Port, tID), nil)
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

	var stats TorrentStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to decode torrent stats response: %v", err)
	}

	return &stats, nil
}
