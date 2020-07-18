
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# taptap
A simple key recognition game built with kids in mind, written in Go.

## About
This project was started by two parents wanting to teach their young children the basics of using a keyboard, and also used as an opportunity to learn Golang.

The game prints a character on the screen, and the objective for the learner is to find and press the corresponding key on the keyboard. Success should return a "good job!" and present a new challenge, while failure will return "try again".

## Requirements and Usage
Taptap makes use of the [MacOS `say` tool](https://ss64.com/osx/say.html), so OS X is pretty much a hard requirement.

Make sure [you have Go installed](https://golang.org/doc/install) and your [`GOPATH` updated](https://golang.org/doc/gopath_code.html#GOPATH).

If you've cloned the repo, run:

`go install`

Otherwise:

`go get github.com/glhdavenport/taptap`

Then run `taptap` and let the little crumb-ridden hands destroy your tempermental Macbook butterfly keyboard!

## Current Todo's

- [] Prompt for learner name input
- [] Evaluate successes/failures without requiring a carriage return
- [] Make the `esc` exits the program
- [] Add support to make it case-insensitive
- [] Create different modes of difficulty, such as letters, numbers, special characters, words, etc (toggled with `tab`?)
- [] Make text-to-voice audio output an optional parameter
- [] Add "emoji roulette" to cycle through different emoji's on successful keypresses or failures
- [] Add scoring mechanisim

### License
[Apache Open Source License v2.0](https://github.com/glhdavenport/taptap/blob/master/LICENSE)