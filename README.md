# cinematic

A [dragonfly](https://github.com/df-mc/dragonfly) tool for moving players in a smooth fashion over given points and a duration.

## Usage

Create a scene with [actions](#actions).
```go
path := act.NewPath(
    []mgl.Vec3{{0, 0, 0}, {0, 10, 5}, {0, 0, 10}}, // coordinates (x, y, z)
    time.Second * 10, // total path duration
    time.Second / 20, // frequency of position updates (higher-smoother, lower-better performance)
)

scene := cinematic.NewScene("introduction").WithActions(map[int]act.Act{ 
	// millisecond timestamp: act
        0: path, 
	1000: act.NewMessage("One second has passed.")
    })
scene.Play(player)
```

### Actions
Path moves the player over a path.
```go
path := act.NewPath(
	[]mgl.Vec3{{0, 0, 0}, {0, 10, 5}, {0, 0, 10}}, // coordinates (x, y, z)
	time.Second * 10, // total path duration
	time.Second / 20, // frequency of position updates (higher-smoother, lower-better performance)
)
```

Rotating path moves the player over a path with rotation.
```go
rpath := act.NewRotatingPath(
    []mgl.Vec3{{0, 0, 0, 0, 0}, {0, 10, 5, 20, 20}, {0, 0, 10, 0, 0}}, // coordinates (x, y, z, yaw, pitch)
    time.Second * 10, // total path duration
    time.Second / 20, // frequency of position updates (higher-smoother, lower-better performance)
)
```

Message sends a message to the player.
```go
message := act.NewMessage("Hello world!")
```

Teleport teleports the player to a position.
```go
teleport := act.NewTeleport(mgl.Vec3{0, 0, 0}, 0, 0)
```

### Saving / Reading
Save/read to/from files.
```go
cinematic.WriteFile(scene, "scene.json")
read := cinematic.ReadFile("scene.json")
```
