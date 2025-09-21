package data

import (
	"encoding/json"
	"log"
	"os"
)

type Data struct {
	SongData   []SongData
	WavFormMap map[SongData][]int
	ImgMap     map[SongData]string
}

// add support for adding songs while running

type SongData struct {
	Title    string
	Artist   string
	Url      string
	Views    string
	Likes    string
	Date     string
	Duration string
	Id       string
	File     string
}

func InitData() Data { //add data cacheing for preformace
	s := GetSongDataSlice()
	return Data{
		SongData: s,
	}
}

func JsonToSongData(file string) SongData { //preformace hit on this function so fix that later
	var data map[string]any
	path := JsonDir

	bytes, err := os.ReadFile(path + file)
	if err != nil {
		log.Fatal("failed to read file: ", err)
	}

	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal("failed to unmashal json data: ", err)
	}

	return SongData{
		Title:    data["title"].(string), //type assertions are bad but it's json sooooo...
		Artist:   data["uploader"].(string),
		Url:      data["url"].(string),
		Views:    data["meta_view"].(string),
		Likes:    data["meta_likes"].(string),
		Date:     data["upload_date"].(string),
		Duration: data["duration_string"].(string),
		Id:       data["id"].(string),
		File:     file,
	}
}

func GetSongDataSlice() []SongData {
	j := GetAllJson()
	var songSlice []SongData
	for i, _ := range j {
		s := JsonToSongData(j[i])
		songSlice = append(songSlice, s)
	}
	return songSlice
}
