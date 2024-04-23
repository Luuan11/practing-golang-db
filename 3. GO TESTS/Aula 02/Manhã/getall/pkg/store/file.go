package store

import (
	"encoding/json"
	"errors"
	"os"
)

type Store interface {
	Read(data interface{}) (interface{}, error)
	Write(data interface{}) (interface{}, error)
}

type Type string

const (
	FileType Type = "file"
)

type FileStore struct {
	FileName string
}


func (f *FileStore) Read(data interface{}) (interface{}, error) {
	jsonR, err := os.ReadFile(f.FileName)
	if err != nil {
		return nil, errors.New("not possible to open the file to read the data")
	}
	return data, json.Unmarshal(jsonR, data)
}

func (f *FileStore) Write(data interface{}) (interface{}, error) {
	jsonW, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return nil, errors.New("failed to marshal the data list")
	}
	return data, os.WriteFile(f.FileName, jsonW, 0644)
}

func NewFileStore(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName}
	}
	return nil
}