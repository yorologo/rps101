# RPS-101 Game in Go

**RPS-101** is an extended version of the classic Rock-Paper-Scissors game, featuring **101 unique gestures**. Implemented in the Go programming language, this game allows players to compete either against another human or against the computer (AI). The game emphasizes fairness by hiding each player's gesture choice from the other when playing in two-player mode.

**Note:** Both the script and the game interface are currently available **only in Spanish**.

## Features

- **101 Unique Gestures:** A vast array of gestures to choose from, enhancing the complexity and fun of the game.
- **Two Game Modes:**
  - **Two-Player Mode:** Play against another human player with hidden gesture inputs to ensure fairness.
  - **Single-Player Mode:** Compete against the computer (AI) with random gesture selections.
- **Customizable Messages:** Unique, thematic winning messages for each gesture interaction.
- **Input Validation:** Ensures that players enter valid gesture numbers.
- **Replay Option:** Allows players to engage in multiple rounds without restarting the game.

## Prerequisites

Before running the **RPS-101** game, ensure you have the following installed on your system:

- **Go Programming Language:** Version 1.16 or later.
  - [Download Go](https://golang.org/dl/)
- **Go Modules Enabled:** Required for dependency management.

## Installation

1. **Clone the Repository:**

   If your script is part of a repository, clone it using:

   ```bash
   git clone https://github.com/yourusername/rps101.git
   cd rps101
   ```

   *If the script is not in a repository, ensure you're in the directory containing `main.go`.*

2. **Initialize Go Module:**

   Initialize a new Go module in your project directory. Replace `rps101` with an appropriate module name if desired.

   ```bash
   go mod init rps101
   ```

3. **Install Dependencies:**

   The game requires the `golang.org/x/term` package to handle hidden input for two-player mode.

   ```bash
   go install golang.org/x/term@latest
   ```

   This command downloads and installs the necessary package and its dependencies.


## Running the Game

After completing the installation steps, you can run the game using the following command:

```bash
go run main.go
```

Alternatively, you can compile the program and then execute the binary:

1. **Compile the Program:**

   ```bash
   go build main.go
   ```

   This command generates an executable file named `main` (or `main.exe` on Windows) in your project directory.

2. **Run the Executable:**

   ```bash
   ./main
   ```

   *(Use `main.exe` if you're on Windows.)*

## Game Mechanics

**RPS-101** operates on an extended version of the traditional Rock-Paper-Scissors rules:

- **Gestures:** There are **101 gestures**, each assigned a unique number from 1 to 101.
- **Winning Logic:** Each gesture beats the next 50 gestures in the list and loses to the previous 50 gestures, following a cyclical pattern.
- **Tie Condition:** If both players choose the same gesture, the result is a tie.

**Example:**

- **Player 1** selects Gesture 1 (**Dinamita**).
- **Player 2** selects Gesture 28 (**Reina**).
- The game calculates the difference and determines that Gesture 1 beats Gesture 28, resulting in Player 1's victory with a unique winning message.

## Code Structure

The **RPS-101** script is organized into several key components:

1. **Imports:**

   - Standard Go packages: `bufio`, `fmt`, `math/rand`, `os`, `strconv`, `strings`, `syscall`, `time`.
   - External package: `golang.org/x/term` for handling hidden input.

2. **Data Structures:**

   - **Gesto:** A struct representing each gesture with a `nombre` (name).
   - **gestos:** A map linking gesture numbers to their corresponding `Gesto`.
   - **mensajesEspecificos:** A map linking gesture numbers to their specific winning messages.

3. **Functions:**

   - **initGestos():** Initializes the list of gestures and their specific winning messages.
   - **mostrarGestos():** Displays the list of available gestures to the players.
   - **obtenerEleccionJugador(numJugador int, hideInput bool):** Retrieves a player's gesture choice, with the option to hide input for fairness.
   - **obtenerEleccionOculta(numJugador int):** Handles hidden input for two-player mode using `golang.org/x/term`.
   - **obtenerEleccionComputadora():** Randomly selects a gesture for the computer opponent.
   - **determinarResultado(gesto1, gesto2 int):** Determines the outcome of the round based on the players' gestures.
   - **obtenerMensajeEspecifico(gestoGanador, gestoPerdedor int):** Retrieves and formats the specific winning message.
   - **jugarOtraRonda():** Asks players if they wish to play another round.
   - **mostrarMenu():** Displays the initial game mode selection menu.

4. **Main Function:**

   - Initializes gestures.
   - Displays the menu and captures the selected game mode.
   - Manages the game loop, handling gesture selections, determining outcomes, and managing replay options.

## Limitations

- **Language:** The entire game, including prompts and messages, is currently implemented **only in Spanish**. There is no support for other languages at this time.
- **Gesture Uniqueness:** Some gestures share the same name but have different numbers (e.g., multiple instances of "Tornado" or "Lobo"). Each gesture number is treated uniquely in the game logic.
- **User Interface:** The game operates solely through the console/terminal with no graphical interface.


## Future Enhancements

To enhance the **RPS-101** game, consider implementing the following features:

1. **Multilingual Support:**
   - Add support for multiple languages, allowing users to select their preferred language at the start.
   - Implement language files or use localization packages to manage translations.

2. **Graphical User Interface (GUI):**
   - Develop a GUI using frameworks like [Fyne](https://fyne.io/) or [Gio](https://gioui.org/) to make the game more visually appealing.

3. **Gesture Categories:**
   - Organize gestures into categories or themes to add strategic depth to the game.

4. **Persistent Score Tracking:**
   - Implement a system to track and save player scores across multiple game sessions.

5. **Online Multiplayer:**
   - Allow players to compete over the internet by integrating networking capabilities.

6. **Expanded Gesture Descriptions:**
   - Provide detailed descriptions or animations for each gesture to enhance user engagement.
