package data 

import (
	"os"
	"log"
	"path"
)

//ytdlp will put the json files in the same dir as the audio files so just moving them here 
func MoveJson() { 
	var musicDir string = "music/"
	var jsonDir string = "audio/data/json/"
	
	dir, err := os.ReadDir("music")
	if err != nil {
		log.Fatal("directory not found")
	}	
	
	for _,e := range dir {
		if path.Ext(e.Name()) == ".json" {
			err := os.Rename(musicDir+e.Name(),jsonDir+e.Name())
			if err != nil {
				log.Fatal(err)
			}		
		}
	}
}

