package main

import (
	/*
		"fmt"
		"time"
	*/
	a "github.com/Setho0o/exho/audio"
)

func main() {
	a.Download(a.Mp3_NoPlaylist(a.SampleEverLong))
	/*
		song := "music/"+"It_Was_A_Good_Day [Gb_-rZB2Foc].wav"
		ch := make(chan audio.Signal)
		p := audio.PlayerInit(ch)

		go p.Play(song)


		for i := range 10 {
			i++
			fmt.Println("l")
			ch <- audio.VolDown
			time.Sleep(1 * time.Second)
		}
	*/
}
