package services

import (
	"math/rand"
)

func RandomChoice(choices []string) string {
	return choices[rand.Intn(len(choices))]
}
