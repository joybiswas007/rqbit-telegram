package client

import "fmt"

type Config struct {
	Bot   BotConfig
	RQBit RQBitConfig
}

type BotConfig struct {
	Token  string `toml:"token"`
	SudoID int64  `toml:"sudoId"`
}

type RQBitConfig struct {
	IpAddress string `toml:"ip"`
	Port      int    `toml:"port"`
}

type File struct {
	Name       string   `json:"name"`
	Components []string `json:"components"`
	Length     int      `json:"length"`
	Included   bool     `json:"included"`
}

type Details struct {
	InfoHash string `json:"info_hash"`
	Name     string `json:"name"`
	Files    []File `json:"files"`
}

type Metadata struct {
	ID           int         `json:"id"`
	Details      Details     `json:"details"`
	OutputFolder string      `json:"output_folder"`
	SeenPeers    interface{} `json:"seen_peers"`
}

type RQBitError struct {
	ErrorType         string `json:"error_kind"`
	HumanReadableText string `json:"human_readable"`
	StatusCode        int    `json:"status"`
	StatusText        string `json:"status_text"`
}

func (e *RQBitError) Error() string {
	return fmt.Sprintf("RQBitError: %s - %s", e.ErrorType, e.HumanReadableText)
}

type Torrents struct {
	Torrents []Torrent `json:"torrents"`
}

type Torrent struct {
	TorrentID int    `json:"id"`
	InfoHash  string `json:"info_hash"`
}
