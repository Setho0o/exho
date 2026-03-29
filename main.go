package main

import (
	"github.com/Setho0o/exho/audio"
	"github.com/Setho0o/exho/audio/data"
)

func main() {
	//	audio.Download(audio.Wav_NoPlaylist("https://www.youtube.com/watch?v=Aibxit_PpAg"))
	d := data.InitData()
	audio.Play(d)
	/*
		d := data.InitData()
		ch := make(chan audio.Signal)
		p := audio.PlayerInit(ch)

		p.Play(d.GetSong())
	*/
}
