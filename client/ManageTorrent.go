package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// pause, resume and delete torrent from client
func (rqb *RQBitConfig) ManageTorrent(tID int, apiPath string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s:%d/torrents/%d/%s", rqb.IpAddress, rqb.Port, tID, apiPath), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if response status code indicates an error
	if resp.StatusCode != http.StatusOK {
		var rqError RQBitError
		if err := json.NewDecoder(resp.Body).Decode(&rqError); err != nil {
			return fmt.Errorf("failed to decode error response: %v", err)
		}
		return &rqError
	}
	return nil
}
