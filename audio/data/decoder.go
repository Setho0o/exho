package data

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	//  "github.com/go-flac/go-flac/v2" for now im only doing wav because setting up wave forms are hard
	//  "github.com/hajimehoshi/go-mp3"
	"github.com/youpy/go-wav"
)

type SoundType int

const (
	Wav SoundType = iota
	//	Mp3
	//	Flac
	Nil
)

func Decode(path string) io.Reader {
	st, err := GetSoundType(path)
	if err != nil {
		log.Fatal(err)
	}
	b, err := GetBytes(path)
	if err != nil {
		log.Fatal(err)
	}

	return BytesToReader(st, b)
}

func GetSoundType(path string) (SoundType, error) {
	path = MusicDir + path
	if strings.HasSuffix(path, "wav") {
		return Wav, nil
	}
	return Nil, fmt.Errorf("Invaild file extention. It must be mp3, flac, or wav")
}

func GetBytes(path string) ([]byte, error) {
	path = MusicDir + path
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed reading file at: "+path, err)
	}
	return fileBytes, nil
}

func BytesToReader(t SoundType, b []byte) io.Reader {
	var reader io.Reader
	switch t {
	case Wav:
		reader = wav.NewReader(bytes.NewReader(b))
	}
	return reader
}
