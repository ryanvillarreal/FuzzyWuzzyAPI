package main

import "github.com/ryanvillarreal/FuzzyWuzzyAPI/cli"

//----------------------------------------------------------
// My attempt at building a API parser/Fuzzer to help automate
// my job so I can do more research and less work. The goal is
// to build a tool that would help teach myself Golang and to
// do something beneficial.  As well golang is nice cause I can
// can cross-compile this on multiple architectures.
//
// This golang project is meant to be an interactive terminal
// application which will let the user make changes and see results
// on the fly.
//
// Please see LICENSE file for license details.

// Global Variables - where we want to start.
var shellMenuContext = "Main"

func main() {
	cli.Shell(shellMenuContext)
}
