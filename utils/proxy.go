package utils

import (
	"github.com/fatih/color"
	"net/url"
)

// define global variables for PROXY ADDR
var PROXY_ADDR string

func SetProxy(addr string) {
	u, err := url.ParseRequestURI(addr)
	if err != nil {
		color.Red("Incorrect Proxy address Format")
	} else if u.Port() == "" {
		color.Red("Specify Port. ")
	} else {
		PROXY_ADDR = u.String()
	}
}

func UnsetProxy() {
	PROXY_ADDR = ""
}
