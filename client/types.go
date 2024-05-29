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

type TorrentDetails struct {
	InfoHash string `json:"info_hash"`
	Name     string `json:"name"`
	Files    []File `json:"files"`
}

type Metadata struct {
	ID           int            `json:"id"`
	Details      TorrentDetails `json:"details"`
	OutputFolder string         `json:"output_folder"`
	SeenPeers    interface{}    `json:"seen_peers"`
}

type RQBitError struct {
	ErrorType         string `json:"error_kind"`
	HumanReadableText string `json:"human_readable"`
	StatusCode        int    `json:"status"`
	StatusText        string `json:"status_text"`
	ID                int    `json:"id"`
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

type TorrentStats struct {
	State         string      `json:"state"`
	FileProgress  []int64     `json:"file_progress"`
	Error         interface{} `json:"error"`
	ProgressBytes int64       `json:"progress_bytes"`
	UploadedBytes int64       `json:"uploaded_bytes"`
	TotalBytes    int64       `json:"total_bytes"`
	Finished      bool        `json:"finished"`
	Live          Live        `json:"live"`
}

type Live struct {
	Snapshot                 Snapshot    `json:"snapshot"`
	AveragePieceDownloadTime Duration    `json:"average_piece_download_time"`
	DownloadSpeed            Speed       `json:"download_speed"`
	UploadSpeed              Speed       `json:"upload_speed"`
	TimeRemaining            interface{} `json:"time_remaining"`
}

type Snapshot struct {
	DownloadedAndCheckedBytes  int64     `json:"downloaded_and_checked_bytes"`
	FetchedBytes               int64     `json:"fetched_bytes"`
	UploadedBytes              int64     `json:"uploaded_bytes"`
	DownloadedAndCheckedPieces int64     `json:"downloaded_and_checked_pieces"`
	TotalPieceDownloadMs       int64     `json:"total_piece_download_ms"`
	PeerStats                  PeerStats `json:"peer_stats"`
}

type PeerStats struct {
	Queued     int `json:"queued"`
	Connecting int `json:"connecting"`
	Live       int `json:"live"`
	Seen       int `json:"seen"`
	Dead       int `json:"dead"`
	NotNeeded  int `json:"not_needed"`
	Steals     int `json:"steals"`
}

type Duration struct {
	Secs  int64 `json:"secs"`
	Nanos int64 `json:"nanos"`
}

type Speed struct {
	Mbps          float64 `json:"mbps"`
	HumanReadable string  `json:"human_readable"`
}
