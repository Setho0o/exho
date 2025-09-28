package audio

import (
	_ "embed"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Setho0o/exho/audio/data"
)

//go:embed ytdlp.sh
var script string

type SongDownload struct {
	format   string
	url      string
	dir      string
	playlist bool
}

func Download(s SongDownload) {
	var no_playlist string
	if !s.playlist {
		no_playlist = "--no-playlist"
	}

	c := exec.Command("bash", "/dev/stdin", s.format, s.url, s.dir, no_playlist) //ytdlp.sh cmds

	c.Stdin = strings.NewReader(script)
	c.Stderr = os.Stderr

	_, e := c.Output()
	if e != nil {
		log.Fatal(e)
	}

	data.MoveJsonAndImageFiles() // move json and imgs after every download
}

func Wav_NoPlaylist(url string) SongDownload {
	return SongDownload{
		format:   "wav",
		url:      url,
		dir:      "",
		playlist: false,
	}
}
func Wav_Playlist(url string) SongDownload {
	return SongDownload{
		format:   "wav",
		url:      url,
		dir:      "",
		playlist: true,
	}
}
