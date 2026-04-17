package main

import (
	//"fmt"

	"github.com/Setho0o/exho/audio/data"
	"github.com/Setho0o/exho/ui"
)

func main() {

	//	audio.Download(audio.Wav_NoPlaylist("https://www.youtube.com/watch?v=9g-TWUn_LM4&list=RD9g-TWUn_LM4&start_radio=1"))
	//	d := data.InitData()
	d := data.Data{}
	//fmt.Println(d)
	ui.Gio(&d)
}
