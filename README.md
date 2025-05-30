# Go WebAssembly Texas Hold'em Poker with Gio UI

This project implements a Texas Hold'em poker game using Go, WebAssembly, and Gio UI, with a mock implementation for SpaceTimeDB integration.

## Project Structure

```
go-wasm-poker/
├── cmd/
│   ├── poker/      # Main application entry point for WASM
│   └── server/     # Simple HTTP server for serving the WASM app
├── pkg/
│   ├── game/       # Core poker game logic
│   ├── ui/         # Gio UI components
│   └── db/         # Database integration (currently mocked)
├── web/           # Web assets and HTML
├── build.sh       # Build script
└── README.md      # This file
```

## Features

- Texas Hold'em poker game logic
- Gio UI for rendering the game
- WebAssembly compilation for browser deployment
- Mock SpaceTimeDB integration (due to lack of official Go client)

## Requirements

- Go 1.21 or later
- Internet connection for dependency downloads

## Building and Running

1. Make the build script executable:
   ```
   chmod +x build.sh
   ```

2. Run the build script:
   ```
   ./build.sh
   ```

3. Start the server:
   ```
   ./server
   ```

4. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

## SpaceTimeDB Integration

Currently, this project uses a mock implementation of SpaceTimeDB as there is no official Go client library for SpaceTimeDB that supports WebAssembly. The mock implementation provides the following features:

- Simulated database connection
- Game state persistence
- Player profile management
- Game history tracking

When an official Go client for SpaceTimeDB becomes available, the mock implementation can be replaced with the real client. The mock implementation is located in `pkg/db/mock_spacetime.go`.

## Future Improvements

- Complete hand evaluation logic implementation
- Add animations and visual effects
- Implement real SpaceTimeDB integration when a Go client becomes available
- Add multiplayer functionality
- Improve UI responsiveness and mobile support

## License

This project is provided as-is without any warranty. You are free to modify and distribute it as needed.
