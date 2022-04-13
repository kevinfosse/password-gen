package main

import (
	"math/rand"
	"time"
)

const (
	letters = "abcdfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers = "0123456789"
	symbols = "!@#$%&*+_-="
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func settingsPassword(lenght int, hasNumber bool, hasLetters bool, hasSymbols bool) string {
	passwordSet := ""

	if hasLetters {
		passwordSet = passwordSet + letters
	}
	if hasNumber {
		passwordSet = passwordSet + numbers
	}
	if hasSymbols {
		passwordSet = passwordSet + symbols
	}

	return generatePassword(lenght, passwordSet)
}

func generatePassword(length int, chars string) string {
	password := ""
	for i := 0; i < length; i++ {
		password += string([]rune(chars)[rand.Intn(len(chars))])
	}
	return password
}
