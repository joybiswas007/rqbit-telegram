package botutils

import (
	"fmt"

	"github.com/joybiswas007/rqbit-telegram/client"
)

// format the torrent stats message
func FormatTorrentStats(stats *client.TorrentStats) string {
	timeRemaining := "NaN"
	if stats.Live.TimeRemaining != nil {
		timeRemaining = fmt.Sprintf("%v", stats.Live.TimeRemaining)
	}

	return fmt.Sprintf(
		"<b>Torrent Status:</b>\n"+
			"State: %s\n"+
			"Progress: %d / %d bytes\n"+
			"Uploaded: %d bytes\n"+
			"Finished: %t\n"+
			"Download Speed: %s\n"+
			"Upload Speed: %s\n"+
			"Peers: %d seen, %d live\n"+
			"Average Piece Download Time: %ds %dns\n"+
			"Time Remaining: %s",
		stats.State,
		stats.ProgressBytes,
		stats.TotalBytes,
		stats.UploadedBytes,
		stats.Finished,
		stats.Live.DownloadSpeed.HumanReadable,
		stats.Live.UploadSpeed.HumanReadable,
		stats.Live.Snapshot.PeerStats.Seen,
		stats.Live.Snapshot.PeerStats.Live,
		stats.Live.AveragePieceDownloadTime.Secs,
		stats.Live.AveragePieceDownloadTime.Nanos,
		timeRemaining,
	)
}
