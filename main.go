package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
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
				fmt.Println("Great work, Woody! 🥳")
				cmd := exec.Command("say", fmt.Sprintf("Great work, Woody! That was the %s key!", input))
				go cmd.Start()
				// found = true
				randS = randomOutput()
			} else if input != string(randS) {
				fmt.Printf("Try again, Woody! 😝\n%s\n", randS)
				cmd := exec.Command("say", fmt.Sprintf("Try again, Woody! You pressed %s. . . Can you find %s? \n", input, randS))
				go cmd.Start()
			}
		}
	}
}
