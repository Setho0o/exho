package data

import (
	"log"
	"os"
	"path"
)

const (
	MusicDir string = "music/"
	JsonDir  string = "audio/data/json/"
	ImgDir   string = "audio/data/img/"
)

// ytdlp will put the json and img files in the same dir as the audio files so just moving them here
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
func ClearAllData() {
}
