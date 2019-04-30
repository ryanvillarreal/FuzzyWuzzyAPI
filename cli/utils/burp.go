package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"os"
)

func BurpRequest(filename string) {
	file, err := os.Open(filename) // For read access.
	if err != nil {
		color.Red(err.Error())
	}
	// Check File Type

	contentType, err := GetFileContentType(file)
	if err != nil {
		color.Red(err.Error())
	}
	switch contentType {
	case "text/xml; charset=utf-8":
		// do things with the XML file here.
		fmt.Println("XML")
	default:
		fmt.Println("Unknown File Type, Burp exports Requests as XML.")
	}
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func Base64Decode(payload string) {
	data, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// change this to do something with the decoded data later
	fmt.Println(data)
}

func Base64Encode(payload string) {
	data := []byte(payload)
	str := base64.StdEncoding.EncodeToString(data)
	// change this to do something with the encode data later
	fmt.Println(str)
}
