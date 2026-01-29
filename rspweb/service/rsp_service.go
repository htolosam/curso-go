package service

import (
	"fmt"
	"math/rand"
)

const (
	ROCK     = 0 // piedra. vence a las tijerads. (tijerads +1) % 3 = 0
	PAPER    = 1 // papel. vence a las piedras. (piedra + 1) % 3 = 1
	SCISSORS = 2 // tijeras ven al papel. (papel + 1) % 3 = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"Bien hecho!",
	"Buen trabajo",
	"Deberias comprar un boleto de loteria",
}

var loseMessages = []string{
	"Que lastima",
	"Intentalo de nuevo!",
	"Hoy simplemente no es tu día!",
}

var drawMessages = []string{
	"loos mas grandes piensan igual",
	"Oh Oh intentalod e nuevo",
	"Nadie gana, pero puedes intentarlo de nuevo",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)
	var computerChoice, roundResult string
	var computerChoiceInt int
	switch computerValue {
	case ROCK:
		computerChoice = "La computadora eligio piedra"
		computerChoiceInt = ROCK
	case PAPER:
		computerChoice = "La computadora eligio papel"
		computerChoiceInt = PAPER
	case SCISSORS:
		computerChoice = "La computadora eligio tijeras"
		computerChoiceInt = SCISSORS
	}
	messageInt := rand.Intn(3)
	var message string
	if playerValue == computerValue {
		roundResult = "Empate"
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
	return Round{message, computerChoice, roundResult, computerChoiceInt, fmt.Sprintf("%d", ComputerScore), fmt.Sprintf("%d", PlayerScore)}
}
