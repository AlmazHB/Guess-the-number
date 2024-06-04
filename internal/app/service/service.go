package service

import (
	"math/rand"
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
