package data

import (
	"encoding/json"
	"log"
	"os"
)

type Data struct {
	SongData   []SongData
	WavFormMap map[SongData][]int //im initializing all the waveform data on startup so very slow for startup
	ImgMap     map[SongData]string
	Exts       map[string]string //SongData.Title as input
}

// add support for adding songs while running

type SongData struct {
	Title    string
	Artist   string
	Views    string
	Likes    string
	Date     string
	Duration string
	Id       string
}

func InitData() Data { //add data cacheing for preformace
	s := GetSongDataSlice()
	e := GetAllMusicExt()
	return Data{
		SongData:   s,
		WavFormMap: GetWaveFormMap(s,e),
		Exts: e,
	}
}

func JsonToSongData(file string) SongData {
	var data map[string]any
	bytes, err := os.ReadFile(JsonDir + file)
	if err != nil {
		log.Fatal("failed to read file: ", err)
	}

	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal("failed to unmashal json data: ", err)
	}
	
	return SongData{ //preformace hit on this function so fix that later
		Title:    data["title"].(string), //type assertions are bad but it's json sooooo...
		Artist:   data["uploader"].(string),
		Views:    data["meta_view"].(string),
		Likes:    data["meta_likes"].(string),
		Date:     data["upload_date"].(string),
		Duration: data["duration_string"].(string),
		Id:       data["id"].(string),
	}
}

func GetSongDataSlice() []SongData {
	json := GetAllJson()
	var songSlice []SongData
	for _, e := range json {
		s := JsonToSongData(e)
		songSlice = append(songSlice, s)
	}
	return songSlice
}

func GetWaveFormMap(s []SongData, ext map[string]string) map[SongData][]int {
	m := map[SongData][]int{}
	for _, e := range s {
		m[e] = GetWave(e.Title+ext[e.Title])
	}
	return m
}
