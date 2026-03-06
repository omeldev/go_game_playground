# Super Gopher Bros

Ein einfaches 2D Plattformer-Spiel in Go mit raylib.

[Cheatsheet](https://www.raylib.com/cheatsheet/cheatsheet.html)

---

## Features

- **Player-Follow Kamera**: Kamera folgt dem Spieler horizontal mit fixierter Ground-Linie
- **Unendliches Scrolling**: Hintergrund wird als nahtlose gekachelte Tiles gerendert
- **Collectibles (Cookies)**: Sammle Cookies, der Spieler zeigt kurz eine angry Animation
- **State Machine**: Player-Animationen (idle, walk, jump, angry) über State-System

## Controls

- **A / D**: Links/Rechts laufen
- **W / SPACE**: Springen

## Project Structure

```
go_game_playground/
├── main.go                          # Entry point und Game Loop
├── game/
│   └── game.go                      # Game-Instanz, Updates und Collisions
├── internal/
│   ├── camera/
│   │   └── camera.go                # Player-Follow Kamera
│   ├── rendering/
│   │   └── background.go            # Infinite Background Rendering
│   └── utils.go                     # Utility-Funktionen
├── math/
│   ├── gravity.go                   # Physik-Konstanten und Funktionen
│   ├── position.go                  # Position Struktur
│   └── velocity.go                  # Velocity Struktur
├── obj/
│   ├── collectible/
│   │   ├── collectible.go           # Base Collectible Klasse
│   │   └── cookie.go                # Cookie Klasse
│   ├── entity/
│   │   ├── player.go                # Player Klasse
│   │   └── player_states.go         # Player State Machine States
│   ├── sprite/
│   │   ├── animation.go             # Animation System
│   │   └── sprite.go                # Sprite Base Klasse
│   └── state/
│       ├── state.go                 # State Interface
│       └── state_machine.go         # State Machine Implementation
└── resources/
    ├── bg.png                       # Background Tile
    ├── cookie/
    │   └── cookie.png               # Cookie Sprite (transparent)
    └── gophers/
        ├── walk_left.png
        ├── walk_right.png
        ├── angry_left.png
        ├── angry_right.png
        └── idle.png
```

## Requirements

```bash
sudo apt-get install libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev libwayland-dev libxkbcommon-dev
```

## Build & Run

```bash
go run main.go
```

## Attribution

**Background Image**: https://opengameart.org/content/country-side-platform-tiles  
Author: ansimuz  
License: CC0

**Gopher Images**: https://github.com/MariaLetta/free-gophers-pack  
Author: Maria Letta  
License: CC0


