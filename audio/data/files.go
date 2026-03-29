package data

import (
	"fmt"
	"log"
	"os"
	"path"
)

const (
	MusicDir string = "music/"
	JsonDir  string = "audio/data/json/"
	ImgDir   string = "audio/data/img/"
)

// ytdlp will all json and img files in the same dir as the audio files, so just moving them to designated dirs
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

func ClearDataDirs() { // clears audio/data/json & audio/data/img
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

func ClearMusicDir() { // clears music/
	mDir, err := os.ReadDir(MusicDir)
	if err != nil {
		log.Fatal("Music directory not found")
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

func GetBytes(path string) ([]byte, error) {
	path = MusicDir + path
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed reading file at: "+path, err)
	}
	return fileBytes, nil
}
