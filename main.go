package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func randomOutput() string {
	var randS string
	// allChars := "ABCDEFG"
	// alpha :=	astrings.Split(allChars, "")
	alpha := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for i := 1; i < 11; i++ {
		alpha = append(alpha, strconv.Itoa(i))
	}

	rand.Seed(time.Now().UnixNano())
	randS = alpha[rand.Intn(len(alpha))]
	println(randS)
	return randS
}

func main() {
	name := promptForName()
	// found := false
	for {
		// if found {
		randS := randomOutput()
		// found = false
		// }
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := (scanner.Text())

			if input == string(randS) {
				fmt.Printf("Great work, %s! 🥳\n", name)
				cmd := exec.Command("say", fmt.Sprintf("Great work, %s! That was the %s key!", name, input))
				go cmd.Start()
				// found = true
				randS = randomOutput()
			} else if input != string(randS) {
				fmt.Printf("Try again, %s! 😝\n%s\n", name, randS)
				cmd := exec.Command("say", fmt.Sprintf("Try again, %s! You pressed %s. . . Can you find %s? \n", name, input, randS))
				go cmd.Start()
			}
		}
	}
}

func promptForName() string {
	fmt.Println("👋 Hi, what is your name?")
	cmd := exec.Command("say", "Hi, what is your name?")
	go cmd.Start()

	var name string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		name = strings.TrimSpace(scanner.Text())

		if name == "" {
			fmt.Println("👋 Hi, what is your name?")
			cmd := exec.Command("say", "Hi, what is your name?")
			go cmd.Start()
			continue
		}

		break
	}

	return name
}
