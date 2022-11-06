package cinematic

import (
	"encoding/json"
	"github.com/unickorn/cinematic/act"
	"io"
	"os"
	"strconv"
)

// Read reads a scene from an io.Reader.
func Read(reader io.Reader) (Scene, error) {
	var s map[string]interface{}
	err := json.NewDecoder(reader).Decode(&s)
	if err != nil {
		return emptyScene, err
	}
	actions := make(map[int]act.Act)
	for k, v := range s["actions"].(map[string]interface{}) {
		i, _ := strconv.Atoi(k)
		actions[i] = act.New(v.(map[string]interface{})["type"].(string)).FromMap(v.(map[string]interface{}))
		if err != nil {
			return emptyScene, err
		}
	}
	return NewScene(s["name"].(string)).WithActions(actions), nil
}

// ReadFile loads a new Path from a file.
func ReadFile(path string) (Scene, error) {
	f, err := os.Open(path)
	if err != nil {
		return emptyScene, err
	}
	return Read(f)
}

// Write writes the path to an io.Writer.
func Write(scene Scene, writer io.Writer) error {
	return json.NewEncoder(writer).Encode(scene)
}

// WriteFile writes the path to file.
func WriteFile(scene Scene, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	return Write(scene, f)
}
