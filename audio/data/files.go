package data

import (
	"log"
	"os"
	"path"
	"strings"
)

const (
	MusicDir string = "music/"
	JsonDir  string = "audio/data/json/"
	ImgDir   string = "audio/data/img/"
)

// ytdlp will put the json and img files in the same dir as the audio files so just moving them here
// Could move a lot of the directory reading into a function but for now it works
func MoveJsonAndImageFiles() {
	dir, err := os.ReadDir(MusicDir)
	if err != nil {
		log.Fatal("directory not found")
	}

	for _, e := range dir {
		if path.Ext(e.Name()) == ".json" {
			err := os.Rename(MusicDir+e.Name(), JsonDir+e.Name())
			if err != nil {
				log.Fatal("failed to rename json file: ", err)
			}
		} else if path.Ext(e.Name()) == ".png" { // I want to change it to a svg in the future for scaling but we'll see
			err := os.Rename(MusicDir+e.Name(), ImgDir+e.Name())
			if err != nil {
				log.Fatal("failed to rename png file: ", err)
			}
		}
	}
}

func GetAllJson() []string {
	dir, err := os.ReadDir(JsonDir)
	if err != nil {
		log.Fatal("json directory not found")
	}
	var jsonFiles []string
	for _, e := range dir {
		if !e.IsDir() {
			jsonFiles = append(jsonFiles, e.Name())
		}
	}
	return jsonFiles
}

func GetAllImgs() []string {
	dir, err := os.ReadDir(ImgDir)
	if err != nil {
		log.Fatal("img directory not found")
	}
	var ImgFiles []string
	for _, e := range dir {
		if !e.IsDir() {
			ImgFiles = append(ImgFiles, e.Name())
		}
	}
	return ImgFiles
}

func GetAllMusicExt() map[string]string {
	dir, err := os.ReadDir(MusicDir)
	if err != nil {
		log.Fatal("img directory not found")
	}
	ext := map[string]string{}
	for _, e := range dir {
		if !e.IsDir() {
			name := strings.TrimSuffix(e.Name(), path.Ext(e.Name()))
			ext[name] = path.Ext(e.Name())
		}
	}
	return ext
}

func ClearDataDirs() {
	jDir, err := os.ReadDir(JsonDir)
	if err != nil {
		log.Fatal("json directory not found")
	}
	for _, e := range jDir {
		if !e.IsDir() {
			err := os.Remove(JsonDir + e.Name())
			if err != nil {
				log.Fatal("failed to remove "+JsonDir+e.Name()+": ", err)
			}
		}
	}
	iDir, err := os.ReadDir(ImgDir)
	if err != nil {
		log.Fatal("Img directory not found")
	}
	for _, e := range iDir {
		if !e.IsDir() {
			err := os.Remove(ImgDir + e.Name())
			if err != nil {
				log.Fatal("failed to remove "+ImgDir+e.Name()+": ", err)
			}
		}
	}
}

func ClearMusicDir() {
	mDir, err := os.ReadDir(MusicDir)
	if err != nil {
		log.Fatal("json directory not found")
	}
	for _, e := range mDir {
		if !e.IsDir() {
			err := os.Remove(MusicDir + e.Name())
			if err != nil {
				log.Fatal("failed to remove "+MusicDir+e.Name()+": ", err)
			}
		}
	}
}

func ClearAllExit() {
	ClearDataDirs()
	ClearMusicDir()
	os.Exit(0)
}
