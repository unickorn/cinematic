package cinematic

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"testing"
	"time"
)

func Test(t *testing.T) {
	p := NewPath([]mgl64.Vec3{{1, 0, -5}, {20, 5, 2}, {5, 10, 5}, {8, 2, 0}, {0, 0, 0}}, 1*time.Second, 20*time.Millisecond)

	Write(p, "go_path.json")
	n := FromFile("go_path.json")

	if fmt.Sprintf("%v", p) != fmt.Sprintf("%v", n) {
		t.Errorf("points are not equal")
	}
}
