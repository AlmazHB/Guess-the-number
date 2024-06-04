package service

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Service struct{}

func New() *Service {
	return &Service{}
}
func (s *Service) GenRandInt() int {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	return target
}

func (s *Service) GetInt() int {
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

func (s *Service) GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)

	return input
}

func (s *Service) CheckNumber(guess, target int) bool {
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
