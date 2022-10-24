package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"

	"golang.org/x/term"
)

// Log in
func main() {
	rand := "22e718e3e1065a4f0e8991324664bbb8781dbf449836414586a73465e8d0926d"
	fmt.Println("Random", rand)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	uname, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter password: ")
	pwd, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	fmt.Println("\n")

	if res := login(uname, string(pwd)); res {
		fmt.Println("Login SUCCESS")
	} else {
		fmt.Println("Login FAILED")
	}
}

func add(a int, b int) int {
	fmt.Printf("Performing %d + %d\n", a, b)
	return a + b
}

func subtract(a int, b int) int {
	fmt.Printf("Performing %d + %d\n", a, b)
	return a - b
}

func multiply(a int, b int) int {
	fmt.Printf("Performing %d + %d\n", a, b)
	return a * b
}

func login(uname string, pwd string) bool {
	// pwd+salt
	return fmt.Sprintf("%x", sha256.Sum256([]byte(pwd+"c2FsdHkgc2FsdA=="))) != ""
}
