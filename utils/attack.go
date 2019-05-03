package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"net/url"
	"time"
)

func Pause() {
	fmt.Println("You should be able to pause using interactive commands")
}

func Resume() {
	fmt.Println("You should be able to resume using interactive commands")
}

func MakeRequest() {
	color.Green("making a request")

	// keep the request from breaking by checking to see if Proxy address is empty
	if PROXY_ADDR != "" {
		// For control over proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport
		proxyUrl, _ := url.Parse(PROXY_ADDR)
		tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			Proxy:              http.ProxyURL(proxyUrl),
		}
		// For control over HTTP client headers,redirect policy, and other settings use a Client
		client := &http.Client{Transport: tr}

		// build the request
		req, err := http.NewRequest(data.Item.Method, data.Item.Protocol+"://"+data.Item.Host.Text, nil)
		if err != nil {
			color.Red(err.Error())
		}

		// call the request.
		resp, err := client.Do(req)
		if err != nil {
			color.Red(err.Error())
		} else {
			color.Green(resp.Status)
		}

	} else {
		// exclude the Proxy Settings so that it doesn't break the request.
		tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}
		// For control over HTTP client headers,redirect policy, and other settings use a Client
		client := &http.Client{Transport: tr}

		// build the request
		req, err := http.NewRequest(data.Item.Method, data.Item.Protocol+"://"+data.Item.Host.Text, nil)
		if err != nil {
			color.Red(err.Error())
		}

		//// call the request.
		resp, err := client.Do(req)
		if err != nil {
			color.Red(err.Error())
		} else {
			color.Green(resp.Status)
		}
	}
}

func Attack() {

}
