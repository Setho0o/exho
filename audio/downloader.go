package audio

import (
	_ "embed"
	"fmt"
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

const SampleLifeLivingYou string = "https://www.youtube.com/watch?v=SwW5QGEBHDo&list=RDSwW5QGEBHDo"
const SampleSwimmingPools string = "https://www.youtube.com/watch?v=B5YNiCfWC3A&list=RDB5YNiCfWC3A"
const SampleEverLong string = "https://www.youtube.com/watch?v=eBG7P-K-r1Y&list=RDeBG7P-K-r1Y"

func Download(s SongDownload) {
	var no_playlist string
	if !s.playlist {
		no_playlist = "--no-playlist"
	}

	c := exec.Command("bash", "/dev/stdin", s.format, s.url, s.dir, no_playlist)

	c.Stdin = strings.NewReader(script)
	c.Stderr = os.Stderr

	b, e := c.Output()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(b))
	data.MoveJson() // move json after every download
}

// will probably remove wav since i cant embed the thumnail with ydlp
func Wav_NoPlaylist(url string) SongDownload {
	return SongDownload{
		format:   "wav",
		url:      url,
		dir:      "",
		playlist: false,
	}
}

func Mp3_NoPlaylist(url string) SongDownload {
	return SongDownload{
		format:   "mp3",
		url:      url,
		dir:      "",
		playlist: false,
	}
}

func Flac_Playlist(url, dir string) SongDownload {
	return SongDownload{
		format:   "flac",
		url:      url,
		dir:      dir,
		playlist: true,
	}
}
