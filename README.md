# StoneWays

StoneWays is an engaging and strategic game, inspired by the classic mechanics of Cairn, reimagined and rebuilt in Go. This implementation brings the essence of strategy and mystique to your command line, offering a unique blend of ancient lore and modern programming.

## Installation

To install StoneWays, you'll need to have Go installed on your system. You can then install StoneWays using `go get`:

```bash
go get github.com/Github-Reneon/stoneways
```

## Usage

After installation, you can run StoneWays directly from the command line:

```bash 
./stoneways
```

Follow the on-screen instructions to play the game. The game is turn-based, so you can take your time to strategize your moves.

### Starting a Game

To start a new game, simply run the command:

```bash
./stoneways new
```

### Game Commands

- `move <area>`: Move your piece to the specified coordinates.
- `search`: Search for something in the area you're in.
- `attack`: Attack the creature that has declared on you.
- `attack <creature name>`: Attack the creature in the same room with you with the passed name.
- `say <phrase>`: Say something in the room.
- `rp <phrase>`: Inconsequentially do something.
- `status`: Display the current status of your character.
- `help`: Display a list of available commands and game rules.

## Features

- **Command-Line Interface**: Easy to navigate command-line interface for a smooth game experience.
- **Strategic Gameplay**: Incorporates the strategic elements of Cairn, challenging players to think ahead and plan their moves.
- **Single Player Mode**: Play against a sophisticated AI that challenges both beginners and experienced players.

## Contributing

Contributions to StoneWays are welcome! Whether it's bug fixes, new features, or improvements to existing features, your input is valued.

To contribute:
1. Fork the repository.
2. Create a new branch for your feature (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

## License

StoneWays is released under the MIT License. See the LICENSE file for more details.

## Acknowledgements

- Inspired by the original Cairn game mechanics.
- Thanks to the Go community for their invaluable resources and support.
