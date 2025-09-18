package main

import (
	"fmt"

	"github.com/Setho0o/exho/audio/data"
)

//	a "github.com/Setho0o/exho/audio"
//	"github.com/Setho0o/exho/ui"

func main() {
	s := data.DecodeJson()
	fmt.Println(s)
//	ui.Gio()
	//a.Download(a.Mp3_NoPlaylist(a.SampleSwimmingPools))
	/*
		song := "music/"+"It_Was_A_Good_Day [Gb_-rZB2Foc].wav"
		ch := make(chan audio.Signal)
		p := audio.PlayerInit(ch)

		go p.Play(song)


		for i := range 10 {
			i++
			fmt.Println"l")
			ch <- audio.VolDown
			time.Sleep(1 * time.Second)
		}
	*/
}
