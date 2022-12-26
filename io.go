package cinematic

import (
	"encoding/json"
	"fmt"
	"github.com/unickorn/cinematic/act"
	"io"
	"os"
)

// Read reads a scene from an io.Reader.
func Read(reader io.Reader) (Scene, error) {
	var s map[string]interface{}
	err := json.NewDecoder(reader).Decode(&s)
	if err != nil {
		return emptyScene, err
	}
	l := s["actions"].([]interface{})
	actions := make([]act.Act, len(l))
	for k, v := range l {
		a := v.(map[string]interface{})
		ac, ok := act.New(a["type"].(string)).(act.WritableAct)
		if !ok {
			return emptyScene, fmt.Errorf("scene contains non-writable action %T", ac)
		}
		actions[k] = ac.FromMap(a)
		if err != nil {
			return emptyScene, err
		}
	}
	return NewScene(s["name"].(string)).WithActions(actions), nil
}

// ReadFile loads a new Scene from a file.
func ReadFile(path string) (Scene, error) {
	f, err := os.Open(path)
	if err != nil {
		return emptyScene, err
	}
	return Read(f)
}

// Write writes the path to an io.Writer.
func Write(scene Scene, writer io.Writer) error {
	if !checkWritable(scene) {
		return fmt.Errorf("scene %T contains non-writable actions", scene)
	}
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

// checkWritable checks if a scene can be written to disk. If none of the actions in the scene contain
// any callables (like in Forms, Dialogs) then all actions can be json serialized and the scene is considered writable.
func checkWritable(scene Scene) bool {
	for _, a := range scene.Actions {
		if _, ok := a.(act.WritableAct); !ok {
			return false
		}
	}
	return true
}
