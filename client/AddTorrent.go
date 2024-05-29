package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// add torrent file or magnet link or http torrent link
func (rqb *RQBitConfig) AddTorrent(input string) (*Metadata, error) {

	var reqBody *bytes.Buffer

	// Check if input is a path to a torrent file
	if fileInfo, err := os.Stat(input); err == nil && !fileInfo.IsDir() {
		fileData, err := os.ReadFile(input)
		if err != nil {
			return nil, fmt.Errorf("failed to read torrent file: %v", err)
		}
		reqBody = bytes.NewBuffer(fileData)
	} else {
		reqBody = bytes.NewBufferString(input)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s:%d/torrents", rqb.IpAddress, rqb.Port), reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-bittorrent")
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

	var metadata Metadata
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		return nil, fmt.Errorf("failed to decode metadata response: %v", err)
	}
	return &metadata, nil
}
