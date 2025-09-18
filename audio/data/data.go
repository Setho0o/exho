package data

import (
	"encoding/json"
	"log"
	"os"
)

type SongData struct {
	title    string
	artist   string
	url      string
	views    string
	likes    string
	date     string
	duration string
	id       string
}

func InitSongData() map[string]SongData { //add data caching for preformace
	return make(map[string]SongData)
}

func DecodeJson() SongData { //preformace hit on this function so fix that later
	var data map[string]any
	path := "audio/data/json/"

	bytes, err := os.ReadFile(path + "Foo_Fighters___Everlong_(Official_HD_Video).info.json")
	if err != nil {
		log.Fatal("failed to read file: ", err)
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal("failed to unmashal json data: ", err)
	}

	return SongData{
		title:    data["title"].(string), //type assertions bad but its json sooooo...
		artist:   data["uploader"].(string),
		url:      data["url"].(string),
		views:    data["meta_view"].(string),
		likes:    data["meta_likes"].(string),
		date:     data["upload_date"].(string),
		duration: data["duration_string"].(string),
		id:       data["id"].(string),
	}
}
