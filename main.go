package main

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// Estructura para representar un gesto
type Gesto struct {
	nombre string
}

// Lista de gestos
var gestos = make(map[int]Gesto)

// Mapas para los mensajes específicos
var mensajesEspecificos = make(map[int]string)

// Inicializamos los gestos con sus nombres y mensajes específicos
func initGestos() {
	nombresGestos := []string{
		"Dinamita", "Tornado", "Pistola", "Roca", "Cuchillo",
		"Fuego", "Muro", "Sol", "Cámara", "Cadena",
		"Motor", "Diablo", "Dragón", "Alien", "Helicóptero",
		"Serpiente", "Hacha", "Tornado", "Arenas movedizas", "Pistola",
		"Diablo", "Computadora", "Castillo", "Sangre", "Víbora",
		"Mono", "Rey", "Reina", "Princesa", "Policía",
		"Bebé", "Hombre", "Mujer", "Hogar", "Tren",
		"Hombre lobo", "Escalera", "Puente", "Teléfono", "Árbol",
		"Tiranosaurio", "Lobo", "Gato", "Pájaro", "Pez",
		"Hormiga", "Cocodrilo", "Lobo", "Sombrero", "Araña",
		"Cucaracha", "Corazón", "Vampiro", "Diablo", "Mar",
		"E.T.", "Espacio", "Superhéroe", "Sirena", "Dios",
		"Tiempo", "Energía", "Amor", "Mago", "Mariposa",
		"Fuego", "Tanque", "Helicóptero", "Agua", "Comida",
		"Ojo", "Planeta", "Lluvia", "Arcoíris", "Electricidad",
		"Bomba", "Viento", "Volcán", "Sombra", "Rey",
		"Fantasma", "Príncipe", "Dios", "Mar", "Viento",
		"Comida", "Enfriamiento", "Alien", "Agua", "Montaña",
		"Guitarra", "Láser", "Ojo", "Policía", "Sirena",
		"Vampiro", "Fantasma", "Trampa", "Caja", "Vampiro",
		"Helicóptero",
	}

	for i, nombre := range nombresGestos {
		gestos[i+1] = Gesto{nombre: nombre}
	}

	// Inicializamos los mensajes específicos para cada gesto ganador
	mensajesEspecificos[1] = "¡La dinamita explota a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[2] = "¡El tornado arrasa con la [OBJETO PERDEDOR]!"
	mensajesEspecificos[3] = "¡La pistola dispara a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[4] = "¡La roca aplasta a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[5] = "¡El cuchillo corta a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[6] = "¡El fuego quema a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[7] = "¡El muro bloquea a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[8] = "¡El sol ilumina y supera a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[9] = "¡La cámara captura a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[10] = "¡La cadena atrapa a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[11] = "¡El motor impulsa más que la [OBJETO PERDEDOR]!"
	mensajesEspecificos[12] = "¡El diablo engaña a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[13] = "¡El dragón incinera a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[14] = "¡El alien abduce a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[15] = "¡El helicóptero sobrevuela a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[16] = "¡La serpiente muerde a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[17] = "¡El hacha corta a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[18] = "¡El tornado destroza a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[19] = "¡Las arenas movedizas hunden a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[20] = "¡La pistola dispara a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[21] = "¡El diablo corrompe a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[22] = "¡La computadora supera a la [OBJETO PERDEDOR] con lógica!"
	mensajesEspecificos[23] = "¡El castillo protege contra la [OBJETO PERDEDOR]!"
	mensajesEspecificos[24] = "¡La sangre vitaliza más que la [OBJETO PERDEDOR]!"
	mensajesEspecificos[25] = "¡La víbora envenena a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[26] = "¡El mono se burla de la [OBJETO PERDEDOR]!"
	mensajesEspecificos[27] = "¡El rey gobierna sobre la [OBJETO PERDEDOR]!"
	mensajesEspecificos[28] = "¡La reina comanda a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[29] = "¡La princesa encanta a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[30] = "¡El policía arresta a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[31] = "¡El bebé desarma a la [OBJETO PERDEDOR] con ternura!"
	mensajesEspecificos[32] = "¡El hombre supera a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[33] = "¡La mujer sobrepasa a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[34] = "¡El hogar acoge a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[35] = "¡El tren atropella a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[36] = "¡El hombre lobo aterroriza a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[37] = "¡La escalera supera a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[38] = "¡El puente conecta más que la [OBJETO PERDEDOR]!"
	mensajesEspecificos[39] = "¡El teléfono comunica mejor que la [OBJETO PERDEDOR]!"
	mensajesEspecificos[40] = "¡El árbol da sombra a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[41] = "¡El tiranosaurio devora a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[42] = "¡El lobo acecha a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[43] = "¡El gato juega con la [OBJETO PERDEDOR]!"
	mensajesEspecificos[44] = "¡El pájaro vuela sobre la [OBJETO PERDEDOR]!"
	mensajesEspecificos[45] = "¡El pez nada alrededor de la [OBJETO PERDEDOR]!"
	mensajesEspecificos[46] = "¡La hormiga supera en número a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[47] = "¡El cocodrilo embosca a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[48] = "¡El lobo caza a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[49] = "¡El sombrero cubre a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[50] = "¡La araña atrapa a la [OBJETO PERDEDOR] en su red!"
	mensajesEspecificos[51] = "¡La cucaracha sobrevive a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[52] = "¡El corazón enamora a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[53] = "¡El vampiro drena a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[54] = "¡El diablo corrompe a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[55] = "¡El mar inunda a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[56] = "¡E.T. sorprende a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[57] = "¡El espacio rodea a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[58] = "¡El superhéroe salva el día contra la [OBJETO PERDEDOR]!"
	mensajesEspecificos[59] = "¡La sirena encanta a la [OBJETO PERDEDOR] con su canto!"
	mensajesEspecificos[60] = "¡Dios domina a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[61] = "¡El tiempo supera a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[62] = "¡La energía impulsa más que la [OBJETO PERDEDOR]!"
	mensajesEspecificos[63] = "¡El amor conquista a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[64] = "¡El mago hechiza a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[65] = "¡La mariposa fascina a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[66] = "¡El fuego quema a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[67] = "¡El tanque aplasta a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[68] = "¡El helicóptero sobrevuela a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[69] = "¡El agua apaga a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[70] = "¡La comida satisface más que la [OBJETO PERDEDOR]!"
	mensajesEspecificos[71] = "¡El ojo observa a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[72] = "¡El planeta soporta más que la [OBJETO PERDEDOR]!"
	mensajesEspecificos[73] = "¡La lluvia moja a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[74] = "¡El arcoíris ilumina más que la [OBJETO PERDEDOR]!"
	mensajesEspecificos[75] = "¡La electricidad electrifica a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[76] = "¡La bomba destruye a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[77] = "¡El viento arrastra a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[78] = "¡El volcán cubre de lava a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[79] = "¡La sombra oscurece a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[80] = "¡El rey gobierna sobre la [OBJETO PERDEDOR]!"
	mensajesEspecificos[81] = "¡El fantasma asusta a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[82] = "¡El príncipe supera a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[83] = "¡Dios domina a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[84] = "¡El mar sumerge a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[85] = "¡El viento dispersa a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[86] = "¡La comida atrae a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[87] = "¡El enfriamiento congela a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[88] = "¡El alien abduce a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[89] = "¡El agua inunda a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[90] = "¡La montaña se eleva sobre la [OBJETO PERDEDOR]!"
	mensajesEspecificos[91] = "¡La guitarra armoniza mejor que la [OBJETO PERDEDOR]!"
	mensajesEspecificos[92] = "¡El láser atraviesa a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[93] = "¡El ojo observa a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[94] = "¡El policía detiene a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[95] = "¡La sirena encanta a la [OBJETO PERDEDOR] con su voz!"
	mensajesEspecificos[96] = "¡El vampiro drena la energía de la [OBJETO PERDEDOR]!"
	mensajesEspecificos[97] = "¡El fantasma asusta a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[98] = "¡La trampa captura a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[99] = "¡La caja encierra a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[100] = "¡El vampiro hipnotiza a la [OBJETO PERDEDOR]!"
	mensajesEspecificos[101] = "¡El helicóptero sobrevuela a la [OBJETO PERDEDOR]!"
}

// Muestra la lista de gestos
func mostrarGestos() {
	fmt.Println("\nLista de gestos:")
	for i := 1; i <= len(gestos); i++ {
		fmt.Printf("%d. %s\n", i, gestos[i].nombre)
	}
}

// Obtiene la elección del jugador.
// Si hideInput es true, oculta la entrada (usado para jugadores en modo dos jugadores).
func obtenerEleccionJugador(numJugador int, hideInput bool) int {
	if hideInput {
		return obtenerEleccionOculta(numJugador)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Jugador %d, elige tu gesto (1-101): ", numJugador)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		eleccion, err := strconv.Atoi(input)
		if err == nil && eleccion >= 1 && eleccion <= 101 {
			return eleccion
		}
		fmt.Println("Entrada inválida. Por favor, ingresa un número entre 1 y 101.")
	}
}

// Obtiene una elección oculta para un jugador (sin mostrar la entrada en la consola)
func obtenerEleccionOculta(numJugador int) int {
	fmt.Printf("Jugador %d, elige tu gesto (1-101): ", numJugador)
	byteInput, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println() // Nueva línea después de la entrada oculta
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return obtenerEleccionJugador(numJugador, false) // Fallback a entrada visible
	}

	input := strings.TrimSpace(string(byteInput))
	eleccion, err := strconv.Atoi(input)
	if err == nil && eleccion >= 1 && eleccion <= 101 {
		return eleccion
	}
	fmt.Println("Entrada inválida. Por favor, ingresa un número entre 1 y 101.")
	return obtenerEleccionJugador(numJugador, false)
}

// Genera una elección aleatoria para la computadora
func obtenerEleccionComputadora() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(101) + 1
}

// Determina el resultado entre dos gestos según la regla general
func determinarResultado(gesto1, gesto2 int) {
	fmt.Printf("\nJugador 1 eligió \"%s\", Jugador 2 eligió \"%s\".\n", gestos[gesto1].nombre, gestos[gesto2].nombre)

	diferencia := (gesto2 - gesto1) % 101
	if diferencia < 0 {
		diferencia += 101
	}

	if diferencia == 0 {
		fmt.Println("Resultado: ¡Empate!")
	} else if diferencia >= 1 && diferencia <= 50 {
		// Gesto 1 vence a Gesto 2
		mensaje := obtenerMensajeEspecifico(gesto1, gesto2)
		fmt.Printf("Resultado: %s Jugador 1 gana.\n", mensaje)
	} else if diferencia >= 51 && diferencia <= 100 {
		// Gesto 1 pierde contra Gesto 2
		mensaje := obtenerMensajeEspecifico(gesto2, gesto1)
		fmt.Printf("Resultado: %s Jugador 2 gana.\n", mensaje)
	}
}

// Función para obtener un mensaje específico o genérico
func obtenerMensajeEspecifico(gestoGanador, gestoPerdedor int) string {
	// Verificamos si hay un mensaje específico para el gesto ganador
	if plantilla, existe := mensajesEspecificos[gestoGanador]; existe {
		// Reemplazamos [OBJETO PERDEDOR] por el nombre del gesto perdedor
		mensaje := strings.Replace(plantilla, "[OBJETO PERDEDOR]", gestos[gestoPerdedor].nombre, -1)
		return mensaje
	}
	// Si no hay un mensaje específico, usamos uno genérico
	return fmt.Sprintf("¡%s vence a %s!", gestos[gestoGanador].nombre, gestos[gestoPerdedor].nombre)
}

// Pregunta si el jugador quiere jugar otra ronda
func jugarOtraRonda() bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n¿Quieres jugar otra ronda? (s/n): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "s" {
			return true
		} else if input == "n" {
			return false
		} else {
			fmt.Println("Entrada inválida. Por favor, ingresa 's' o 'n'.")
		}
	}
}

// Muestra el menú inicial y retorna el modo de juego seleccionado
func mostrarMenu() int {
	fmt.Println("Bienvenido a RPS-101, elige un modo de juego:")
	fmt.Println("1. Jugar contra otro jugador")
	fmt.Println("2. Jugar contra la computadora")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Selecciona una opción (1 o 2): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		modo, err := strconv.Atoi(input)
		if err == nil && (modo == 1 || modo == 2) {
			return modo
		}
		fmt.Println("Entrada inválida. Por favor, elige 1 o 2.")
	}
}

func main() {
	initGestos()
	modoJuego := mostrarMenu()

	for {
		mostrarGestos()
		gestoJugador1 := 0
		gestoJugador2 := 0

		if modoJuego == 1 {
			// Modo de dos jugadores
			fmt.Println("\n--- Jugador 1 ---")
			gestoJugador1 = obtenerEleccionJugador(1, true) // Oculta la entrada
			fmt.Println("\n--- Jugador 2 ---")
			gestoJugador2 = obtenerEleccionJugador(2, false) // Entrada visible
		} else {
			// Modo contra la computadora
			fmt.Println("\n--- Jugador 1 ---")
			gestoJugador1 = obtenerEleccionJugador(1, false) // Entrada visible
			gestoJugador2 = obtenerEleccionComputadora()
			fmt.Printf("Jugador 2 (Computadora) eligió \"%s\".\n", gestos[gestoJugador2].nombre)
		}

		determinarResultado(gestoJugador1, gestoJugador2)

		if !jugarOtraRonda() {
			fmt.Println("\n¡Gracias por jugar RPS-101!")
			break
		}
	}
}
