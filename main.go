package main

import (
	"github.com/Setho0o/exho/audio/data"
	"github.com/Setho0o/exho/audio"
//	"github.com/Setho0o/exho/ui"
)



func main() {
//	audio.Download(audio.Wav_NoPlaylist("https://music.youtube.com/watch?v=rqEbv3Z-i4c"))
	d := data.InitData()
		ch := make(chan audio.Signal)
		p := audio.PlayerInit(ch)

		p.Play(d.SongData[0].Title+d.Exts[d.SongData[0].Title])
}
