package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func Pause() {
	fmt.Println("You should be able to pause using interactive commands")
}

func Resume() {
	fmt.Println("You should be able to resume using interactive commands")
}

func Request() {
	color.Green("making a request")
}
