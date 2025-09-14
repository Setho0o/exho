package audio

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/go-flac/go-flac/v2"
	"github.com/hajimehoshi/go-mp3"
	"github.com/youpy/go-wav"
)

type SoundType int

const (
	Mp3 SoundType = iota
	Wav
	Flac
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
	
	return BytesToReader(st,b)
}

func GetSoundType(path string) (SoundType, error) {
	if strings.HasSuffix(path, "mp3") {
		return Mp3, nil
	} else if strings.HasSuffix(path, "wav") {
		return Wav, nil
	} else if strings.HasSuffix(path, "flac") {
		return Flac, nil
	}
	return Nil, fmt.Errorf("Invaild file extention. It must be mp3, flac, or wav")
}

func GetBytes(path string) ([]byte, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed reading file at: "+path, err)
	}
	return fileBytes, nil
}

func BytesToReader(t SoundType, b []byte) io.Reader {
	var reader io.Reader
	switch t {
	case Mp3:
		m, err := mp3.NewDecoder(bytes.NewReader(b))
		if err != nil {
			log.Fatal("failed decoding mp3", err)
		}
		reader = m

	case Wav:
		reader = wav.NewReader(bytes.NewReader(b))

	case Flac:
		f, err := flac.ParseBytes(bytes.NewReader(b))
		if err != nil {
			log.Fatal("failed decoding flac", err)
		}
		reader = f.Frames
	}
	return reader
}


