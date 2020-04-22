package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/sethvargo/go-password/password"
)

func main() {

	// Make sure we got an argument
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "No Default Password Supplied")
		return
	}

	// Password Arg
	pass := os.Args[1]

	// Generate Salt
	salt, err := GenRandom()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating salt: %e", err)
		return
	}

	// Base64 for later
	saltb64 := base64.StdEncoding.EncodeToString([]byte(salt))

	// Hash the password
	h := sha512.New()
	h.Write([]byte(pass))
	h.Write([]byte(salt))

	// Base64
	passHash := base64.StdEncoding.EncodeToString(h.Sum(nil))

	fmt.Fprintf(os.Stdout, "0$%s$%s", saltb64, passHash)
}

// GenRandom ...
func GenRandom() (rand string, err error) {
	// Rand input
	randInput := password.GeneratorInput{}
	randInput.UpperLetters = password.UpperLetters
	randInput.Digits = password.Digits
	randGen, err := password.NewGenerator(&randInput)
	if err != nil {
		return
	}

	// Generate Salt
	return randGen.Generate(6, 4, 0, false, true)
}
