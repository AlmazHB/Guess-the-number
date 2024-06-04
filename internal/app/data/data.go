package data

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Service interface {
	GenRandInt() int
}
type Data struct {
	s Service
}

func New(s Service) *Data {
	return &Data{
		s: s,
	}
}

func (d *Data) CheckNumber(guess, target int) bool {
	success := false
	if guess < target {
		fmt.Println("Oops.Your guess was LOW.")
	} else if guess > target {
		fmt.Println("Oops. Your guess was HIGH")
	} else {
		success = true
		fmt.Println("Good job! You guessed it!")
		return success
	}
	return success
}

func (d *Data) GetInt() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	inputINT, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return inputINT
}
