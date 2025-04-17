package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)

		switch i {
		case 0:
			fmt.Println("El jugador eligio PIEDRA")
		case 1:
			fmt.Println("El jugador eligio PAPEL")
		case 2:
			fmt.Println("El jugador eligio Tijera")
		}

		fmt.Printf("Mensaje: %s\n", round.Message)
		fmt.Printf("Eleccion de la computadora: %s\n", round.ComputerChoice)
		fmt.Printf("Resultado: %s\n", round.RoundResult)
		fmt.Printf("Eleccion de la computadora en entero: %d\n", round.ComputerChoiseInt)
		fmt.Printf("Puntaje computadora: %s\n", round.ComputerScore)
		fmt.Printf("Puntaje jugador : %s\n", round.PlayerScore)

		fmt.Println("=========================================================================")
	}
}
