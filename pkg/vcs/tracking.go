package vcs

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
)

func Track(basePath string, file string, tracked *[]TrackedItem) {
	data, err := os.ReadFile(filepath.Join(basePath, file))
	if err != nil {
		Log.Error(err.Error())
	}
	split := bytes.Split(data, []byte{13})
	var item = TrackedItem{file, split}
	*tracked = append(*tracked, item)
}

func LoadPrevious(basePath string, tracked *[]TrackedItem) {
	data, err := os.ReadFile(filepath.Join(basePath, "\\.gvc\\BASE.json"))
	if err != nil {
		Log.Error(err.Error())
	}
	jErr := json.Unmarshal(data, &tracked)
	if jErr != nil {
		Log.Error(err.Error())
	}

	// TODO: Read and build off of each commit/diff

}

func LoadStaged(basePath string, tracked *[]TrackedItem) {
	data, err := os.ReadFile(filepath.Join(basePath, "\\.gvc\\STAGING.json"))
	if err != nil {
		Log.Error(err.Error())
	}
	jErr := json.Unmarshal(data, &tracked)
	if jErr != nil {
		Log.Error(err.Error())
	}

}

func RemoveStaged(basePath string) {
	err := os.Remove(filepath.Join(basePath, "\\.gvc\\STAGING.json"))
	if err != nil {
		Log.Error(err.Error())
	}
}

func Write(basePath string, name string, data interface{}) {
	j, er := json.MarshalIndent(data, "", "    ")
	if er != nil {
		Log.Error(er.Error())
	}

	writeEr := os.WriteFile(filepath.Join(basePath, ".gvc\\"+name+".json"), j, 0644)
	if writeEr != nil {
		Log.Error(writeEr.Error())
	}
}

func SearchTracked(tracked []TrackedItem, name string) (TrackedItem, bool) {
	for _, item := range tracked {
		if item.Name == name {
			return item, true
		}
	}
	return TrackedItem{}, false
}
