package app

import (
	"fmt"
	"random/internal/app/data"
	"random/internal/app/service"
)

type App struct {
	s *service.Service
	d *data.Data
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.d = data.New(a.s)
	return a, nil
}

func (a *App) Run() {
	target := a.s.GenRandInt()
	fmt.Println(target)
	fmt.Println("I've chosen a random random number between 1 and 100.")
	fmt.Println("Can you guess it")

	success := false
	for guess := 0; guess < 10; guess++ {
		fmt.Println("You have", 10-guess, "guess left.")

		fmt.Print("Make a guess:")
		inputINT := a.d.GetInt()
		success = a.d.CheckNumber(inputINT, target)
		if success {
			break
		}

	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}

}
