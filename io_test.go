package cinematic

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/unickorn/cinematic/act"
	"testing"
	"time"
)

// TestIO
func TestIO(t *testing.T) {

	// Create a new path
	p := act.NewPath([]mgl64.Vec3{{1, 0, -5}, {20, 5, 2}, {5, 10, 5}, {8, 2, 0}, {0, 0, 0}}, 2*time.Second, 20*time.Millisecond)
	// New scene using the path
	s := NewScene("test").WithActions(map[int]act.Act{
		0:    p,
		500:  act.NewMessage("500ms"),
		1500: act.NewMessage("Hello world! (1500ms)"),
		2000: act.NewMessage("2000ms! complete"),
	})
	// Write to file
	err := WriteFile(s, "test_path.json")
	if err != nil {
		t.Fatal(err)
	}
	//Read from file
	n, err := ReadFile("test_path.json")
	if err != nil {
		t.Fatal(err)
	}

	if fmt.Sprintf("%v", s) != fmt.Sprintf("%v", n) {
		t.Logf("expected: %v", s)
		t.Logf("got: %v", n)
		t.Errorf("scenes are not equal")
	}
}
