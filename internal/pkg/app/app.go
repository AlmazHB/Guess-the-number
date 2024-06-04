package app

import (
	"errors"
	"fmt"

	"random/internal/app/service"
)

type App struct {
	s *service.Service
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	return a, nil
}

func (a *App) Run() {
	target := a.s.GenRandInt()

	fmt.Println("I've chosen a random random number between 1 and 100.")
	fmt.Println("Can you guess it")

	success := false
	for guess := 0; guess < 10; guess++ {
		fmt.Println("You have", 10-guess, "guess left.")

		fmt.Print("Make a guess:")
		inputINT := a.s.GetInt()
		success = a.s.CheckNumber(inputINT, target)
		if success {
			break
		}

	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}

}

func (a *App) Start() error {
	fmt.Println("Priwet eta Mini Game s chislami\n Chtoby nachat widite 'start'")

	input := a.s.GetInput()
	if input != "start" {
		err := errors.New("Ops neozhidannyy otwet.\n Prowerti swoy otwet poprobuyti zapustit igru zanowa...")
		return err
	}
	return nil
}
