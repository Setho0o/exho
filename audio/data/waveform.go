package data

import (
	"bytes"
	"github.com/go-audio/wav"
	"log"
)

func GetWave(song string) []int {
	x, err := GetBytes(song)
	if err != nil {
		log.Fatal("failed to decode song for wavform: ",err)
	}

	w := wav.NewDecoder(bytes.NewReader(x))
	w.FwdToPCM()
	buf, err := w.FullPCMBuffer()
	if err != nil {
		log.Fatal("failed to retrive pcm buffer: ", err)
	}
	//fmt.Print(buf.Data)// length of wav 22199764
	//fmt.Print(len(buf.Data))// length of wav 22199764
	var wayLessThan20M []int
	for i := 0; i <= len(buf.Data)/10*1000; i += 2220 { // gonna need to redo all of this but it works
		if i >= len(buf.Data) {
			break
		}
		wayLessThan20M = append(wayLessThan20M, buf.Data[i])
	}
	return wayLessThan20M
}
