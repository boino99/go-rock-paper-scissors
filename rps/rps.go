package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiseInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"Bien hecho",
	"Excelente eleccion",
	"Hoy estas de suerte, sigue asi",
	"Felicidades mejor ve a comprar un boleto que hoy sacas el premio gordo",
}
var loseMessages = []string{
	"Que lastima",
	"Mas suerte la proxima",
	"Cuidado al salir a la calle, capaz y te cae un piano",
	"Has perdido, que deshonra",
}
var drawMessages = []string{
	"Un empate",
	"Que coincidencia han elegido lo mismo",
	"Nadie gana pero todos perdimos nuestro tiempo",
	"Han quedado empatados, mejor vuelve a intertalo",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligio PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligio TIJERA"
	}

	messageInt := rand.Intn(4)

	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora gana"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiseInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
