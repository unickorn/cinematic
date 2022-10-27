package cinematic

import (
	"encoding/json"
	"os"
)

// FromFile loads a new Path from a file.
func FromFile(path string) Path {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	var n NormalPath
	err = json.NewDecoder(f).Decode(&n)
	if err == nil {
		return &n
	}
	var r RotatingPath
	err = json.NewDecoder(f).Decode(&r)
	if err == nil {
		return &r
	}
	return emptyPath
}

// Write writes the path to file.
func Write(path Path, file string) {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(f).Encode(path)
	if err != nil {
		panic(err)
	}
}
