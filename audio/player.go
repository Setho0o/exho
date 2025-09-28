package audio

import (
	"fmt"
	"time"

	"github.com/Setho0o/exho/audio/data"
	"github.com/ebitengine/oto/v3"
)

type Player struct {
	ctx    *oto.Context
	player *oto.Player
	ch     <-chan Signal
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
		ctx:    otoCtx,
		player: nil,
		ch:     ch,
	}
}

func (p *Player) Play(song string) {
	p.player = p.ctx.NewPlayer(data.Decode(song))
	p.player.Play()

	for p.player.IsPlaying() {
		fmt.Println(p.player.BufferedSize())
		p.CheckSignals()
		time.Sleep(time.Millisecond)
	}

	err := p.player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}
