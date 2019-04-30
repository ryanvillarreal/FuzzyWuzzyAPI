package utils

import (
	"os"
)

var Debug = false

// CurrentDir is the current directory where Slackord was executed from
var CurrentDir, _ = os.Getwd()

// Version is a constant variable containing the version number for the Slackord package
const Version = "0.0.1 BETA"
