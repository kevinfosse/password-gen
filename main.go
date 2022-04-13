package main

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/inancgumus/screen"
)

var qs = []*survey.Question{
	{
		Name:      "length",
		Prompt:    &survey.Input{Message: "Password length : "},
		Validate:  survey.Required,
		Transform: survey.Title,
	},

	{
		Name:   "hasLetters",
		Prompt: &survey.Confirm{Message: "Do you want letters?", Default: true},
	},
	{
		Name:   "hasSymbols",
		Prompt: &survey.Confirm{Message: "Do you want symbols?", Default: true},
	},
	{
		Name:   "hasNumbers",
		Prompt: &survey.Confirm{Message: "Do you want numbers?", Default: true},
	},
}

func main() {

	randcolor := colorsText()
	answers := struct {
		Length     int
		HasLetters bool
		HasSymbols bool
		HasNumbers bool
	}{}

	log.SetPrefix("Error: ")
	log.SetFlags(0)

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	screen.Clear()
	screen.MoveTopLeft()
	password := settingsPassword(answers.Length, answers.HasNumbers, answers.HasLetters, answers.HasSymbols)
	fmt.Println("Password :", string(randcolor), password, "\033[0m")
}
