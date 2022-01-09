package cinematic

import (
	"github.com/cnkei/gospline"
	"github.com/go-gl/mathgl/mgl32"
	"time"
)

// Splines is a struct wrapping over 3 spline-generated coordinates.
type Splines struct {
	x []float64
	y []float64
	z []float64
}

// NewSplines returns a Splines ready to use from the splines and duration given.
func NewSplines(x, y, z gospline.Spline, duration, interval time.Duration) *Splines {
	d, i := float64(duration), float64(interval)
	return &Splines{
		x: x.Range(0, d, i),
		y: y.Range(0, d, i),
		z: z.Range(0, d, i),
	}
}

// At returns a vector3 at a given step. The step should be an increment of duration / Interval.
func (s *Splines) At(i int) mgl32.Vec3 {
	return mgl32.Vec3{float32(s.x[i]), float32(s.y[i]), float32(s.z[i])}
}

// RotationSplines is a struct wrapping over 3 spline-generated coordinates plus yaw and pitch.
type RotationSplines struct {
	x     []float64
	y     []float64
	z     []float64
	yaw   []float64
	pitch []float64
}

// NewRotationSplines returns a Splines ready to use from the splines and duration given.
func NewRotationSplines(x, y, z, yaw, pitch gospline.Spline, duration, interval time.Duration) *RotationSplines {
	d, i := float64(duration), float64(interval)
	return &RotationSplines{
		x:     x.Range(0, d, i),
		y:     y.Range(0, d, i),
		z:     z.Range(0, d, i),
		yaw:   yaw.Range(0, d, i),
		pitch: pitch.Range(0, d, i),
	}
}

// At returns a vector3 at a given step.
func (s *RotationSplines) At(i int) mgl32.Vec3 {
	return mgl32.Vec3{float32(s.x[i]), float32(s.y[i]), float32(s.z[i])}
}

// Yaw returns a yaw at a given step.
func (s *RotationSplines) Yaw(i int) float32 {
	y := s.yaw[i]
	if y < -180 {
		y += 360
	}
	return float32(y)
}

// Pitch returns a pitch at a given step.
func (s *RotationSplines) Pitch(i int) float32 {
	return float32(s.pitch[i])
}
