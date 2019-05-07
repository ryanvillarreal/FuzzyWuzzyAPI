package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// define global variables
var threads = 5
var timeout = 5 * time.Second // will eventually add in functionality to change this.
var wg sync.WaitGroup

func Pause() {
	fmt.Println("You should be able to pause using interactive commands")
}

func Resume() {
	fmt.Println("You should be able to resume using interactive commands")
}

func SetThreads(change int) {
	threads = change
}

func GetThreads() (num string) {
	return strconv.Itoa(threads)
}

func MakeRequest(payload string) {
	if data.Item.Host.Text == "" {
		color.Red("No data set. Import Data first. \n")
	} else {
		// keep the request from breaking by checking to see if Proxy address is empty
		if PROXY_ADDR != "" {
			// For control over proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport
			proxyUrl, _ := url.Parse(PROXY_ADDR)
			tr := &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    timeout,
				DisableCompression: true,
				TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
				Proxy:              http.ProxyURL(proxyUrl),
			}
			// For control over HTTP client headers,redirect policy, and other settings use a Client
			client := &http.Client{Transport: tr}

			// build the request
			req, err := http.NewRequest(data.Item.Method, data.Item.Protocol+"://"+data.Item.Host.Text+data.Item.Path+"/"+payload, nil)

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
				IdleConnTimeout:    timeout,
				DisableCompression: true,
				TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			}
			// For control over HTTP client headers,redirect policy, and other settings use a Client
			client := &http.Client{Transport: tr}

			// build the request
			req, err := http.NewRequest(data.Item.Method, data.Item.Protocol+"://"+data.Item.Host.Text+data.Item.Path+"/"+payload, nil)
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
	time.Sleep(10 * time.Second)
	defer wg.Done()
}

func SingleRequest() {
	wg.Add(1)
	go MakeRequest("")
	wg.Wait()
}

func Attack() {
	// build options here
	runtime.GOMAXPROCS(threads)

	if GetPayloadLen() <= 0 {
		color.Red("No Payloads imported yet.")
	} else {
		for _, attack := range payloads {
			wg.Add(1)
			go MakeRequest(attack)
		}
		wg.Wait()
	}
}
