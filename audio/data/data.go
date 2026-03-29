package data

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"strings"
)

type Data struct {
	Current    int
	SongData   []SongData
	WavFormMap map[SongData][]int //im initializing all the waveform data on startup so very slow for startup
	ImgMap     map[SongData]string
}

func (d *Data) GetSong() string {
	return d.SongData[d.Current].Title + d.SongData[d.Current].Ext
}
func (d *Data) GetTitle() string {
	return d.SongData[d.Current].Title
}
func (d *Data) GetExt() string {
	return d.SongData[d.Current].Ext
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
	Ext      string
}

func InitData() Data { //add data cacheing for preformace
	s := GetSongDataSlice()
	return Data{
		SongData:   s,
		WavFormMap: GetWaveFormMap(s),
	}
}

func GetSongDataSlice() []SongData {
	json := GetAllJson()
	ex := GetAllMusicExt()
	var songSlice []SongData
	for _, e := range json {
		s := JsonToSongData(e, ex)
		songSlice = append(songSlice, s)
	}
	return songSlice
}

func JsonToSongData(file string, exts map[string]string) SongData {
	var data map[string]any
	bytes, err := os.ReadFile(JsonDir + file)
	if err != nil {
		log.Fatal("failed to read file: ", err)
	}

	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal("failed to unmashal json data: ", err)
	}

	s := SongData{ //preformace hit on this function so fix that later
		Title:    data["title"].(string), //type assertions are bad but it's json sooooo...
		Artist:   data["uploader"].(string),
		Views:    data["meta_view"].(string),
		Likes:    data["meta_likes"].(string),
		Date:     data["upload_date"].(string),
		Duration: data["duration_string"].(string),
		Id:       data["id"].(string),
	}
	s.Ext = exts[s.Title]
	return s
}

// you'd think id be able to get the music ext from the json file but no it only lists the mp4 video file
// so its better just to read the music dir and get the exts seperately, also Im using a map since i cant guaranty the order
// between each dir.
func GetAllMusicExt() map[string]string { // takes in songdata.Title
	dir, err := os.ReadDir(MusicDir)
	if err != nil {
		log.Fatal("music directory not found")
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

func GetWaveFormMap(s []SongData) map[SongData][]int {
	m := map[SongData][]int{}
	for _, e := range s {
		m[e] = GetWave(e.Title + e.Ext)
	}
	return m
}
