# cinematic

A [dragonfly](https://github.com/df-mc/dragonfly) tool for moving players in a smooth fashion over given points and a duration.

### Usage

Move a player over a path.
```go
path := cinematic.NewPath(
	[]mgl.Vec3{{0, 0, 0}, {0, 10, 5}, {0, 0, 10}}, // coordinates (x, y, z)
	time.Second * 10, // total path duration
	time.Second / 20, // frequency of position updates (higher-smoother, lower-better performance)
)
path.Move(p) // *player.Player
```

Move a player over a path with rotation.
```go
rpath := cinematic.NewRotatingPath(
    []mgl.Vec3{{0, 0, 0, 0, 0}, {0, 10, 5, 20, 20}, {0, 0, 10, 0, 0}}, // coordinates (x, y, z, yaw, pitch)
    time.Second * 10, // total path duration
    time.Second / 20, // frequency of position updates (higher-smoother, lower-better performance)
)
path.Move(p) // *player.Player
```

Save/read to/from files.
```go
cinematic.Write(path, "path.json")
cinematic.Write(rpath, "rotating_path.json")

normal := cinematic.FromFile("path.json")
rotating := cinematic.FromFileRotating("rotating_path.json")
```
