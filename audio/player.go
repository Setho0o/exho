package audio

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/Setho0o/exho/audio/data"
	"github.com/ebitengine/oto/v3"
	"github.com/youpy/go-wav"
)

type Player struct {
	Ctx    *oto.Context
	Player *oto.Player
	Ch     <-chan Signal
}

func PlayerInit(ch chan Signal) Player {
	op := &oto.NewContextOptions{}
	op.SampleRate = 48000
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	<-readyChan

	return Player{
		Ctx:    otoCtx,
		Player: nil,
		Ch:     ch,
	}
}

func Play(d data.Data) {
	ch := make(chan Signal)
	p := PlayerInit(ch)

	file, err := os.Open(data.MusicDir + d.GetSong()) //this file needs to stay open untill the song is done since were streaming to save mem

	decodedFile := Decode(file, d)
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

// im planning on add more than just wav for audio extentions, but for now its easier to just stick with it since making a waveform is a pain in the ass
func Decode(file *os.File, d data.Data) io.Reader {
	switch d.GetExt() {
	case ".wav":
		return wav.NewReader(file)
	default:
		log.Fatal("no extention: ", d)

	}
	return nil
}
