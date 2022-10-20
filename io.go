package cinematic

import (
	"encoding/json"
	"os"
)

// FromFile loads a new Path from a file.
func FromFile(path string) *NormalPath {
	// open file and read json
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	// parse json
	p := NormalPath{}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		return &NormalPath{}
	}
	return NewPath(p.Points, p.Duration, p.Interval)
}

// FromFileRotating loads a new RotatingPath from a file.
func FromFileRotating(path string) *RotatingPath {
	// open file and read json
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	// parse json
	p := RotatingPath{}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		return &RotatingPath{}
	}
	return NewRotatingPath(p.Points, p.Duration, p.Interval)
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
