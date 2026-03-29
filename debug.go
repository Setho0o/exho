package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/Setho0o/exho/audio"
	"github.com/Setho0o/exho/audio/data"
	"github.com/youpy/go-wav"
)

type DebugOps struct {
	Clrscn bool // linux only
	Go     bool
}

func clrscn() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (ops DebugOps) PrintSong(d data.Data) {
	if ops.Go { // using recursion to make this concurrent
		ops.Go = false
		go ops.PrintSong(d)
	} else if ops.Clrscn {
		clrscn()
	}
	fmt.Println(d.GetSong())
}

func DebugPlay() {
	d := data.InitData()

	ch := make(chan audio.Signal)
	p := audio.PlayerInit(ch)

	file, err := os.Open(data.MusicDir + d.GetSong()) //this file needs to stay open untill the song is done since were streaming to save mem

	decodedFile := audio.Decode(file, d)
	wav.NewReader(file)

	p.Player = p.Ctx.NewPlayer(decodedFile)
	p.Player.Play()

	for p.Player.IsPlaying() {
		p.CheckSignals()
		time.Sleep(time.Millisecond * 1)
	}

	err = file.Close()
	if err != nil {
		log.Fatal("failed close song", err)
	}

	err = p.Player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}
