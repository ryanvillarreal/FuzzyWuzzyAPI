package utils

import (
	"encoding/xml"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"os"
)

// define global variables
var data Items

//
func BurpImport(filename string) {
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

		color.Yellow("Parsing the XML File")

		// convert file into a byte array for parsing?
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			color.Red(err.Error())
		}

		// Parse XML now
		xml.Unmarshal(content, &data) // data holds everything, but it's becoming accessible after the function exits.

		color.Green("Finished Importing...")

	default:
		color.Red("Unknown File Type, Burp exports Requests as XML.")
	}
}

// Structure created using the online XML parser
// https://www.onlinetool.io/xmltogo/
type Items struct {
	XMLName     xml.Name `xml:"items"`
	Text        string   `xml:",chardata"`
	BurpVersion string   `xml:"burpVersion,attr"`
	ExportTime  string   `xml:"exportTime,attr"`
	UserAgent   string   `xml:"useragent,attr"`

	Item struct {
		Text string `xml:",chardata"`
		Time string `xml:"time"`
		URL  string `xml:"url"`

		Host struct {
			Text string `xml:",chardata"`
			Ip   string `xml:"ip,attr"`
		} `xml:"host"`
		Port      string `xml:"port"`
		Protocol  string `xml:"protocol"`
		Method    string `xml:"method"`
		Path      string `xml:"path"`
		Extension string `xml:"extension"`

		Request struct {
			Text   string `xml:",chardata"`
			Base64 string `xml:"base64,attr"`
		} `xml:"request"`
		Status         string `xml:"status"`
		Responselength string `xml:"responselength"`
		Mimetype       string `xml:"mimetype"`

		Response struct {
			Text   string `xml:",chardata"`
			Base64 string `xml:"base64,attr"`
		} `xml:"response"`
		Comment string `xml:"comment"`
	} `xml:"item"`
}

// Check FileType to make sure it's approved.
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

func CheckData() (check bool) {
	if &data == nil {
		color.Red("Data Hasn't been loaded yet.")
		return false
	} else {
		return true
	}
}

func PrintInfo(attribute string) {
	switch attribute {
	case "all":
		color.Green("Export Time: " + data.ExportTime)
		color.Green("Hostname: " + data.Item.Host.Text)
		color.Green("IP: " + data.Item.Host.Ip)
		color.Green("Port: " + data.Item.Port)
		color.Green("Proto: " + data.Item.Protocol)
		color.Green("Method: " + data.Item.Method)
		color.Green("Path: " + data.Item.Path)
		color.Green("Extension: " + data.Item.Extension)
		color.Green("Status: " + data.Item.Status)
		color.Green("Response Length: " + data.Item.Responselength)
		color.Green("Mime Type: " + data.Item.Mimetype)
		color.Green("Comments: " + data.Item.Comment)
	case "time":
		color.Green(data.ExportTime)
	case "url":
		color.Green(data.Item.Host.Text)
	case "host":
		color.Green(data.Item.Host.Text)
	case "ip":
		color.Green(data.Item.Host.Ip)
	case "port":
		color.Green(data.Item.Port)
	case "protocol":
		color.Green(data.Item.Protocol)
	case "method":
		color.Green(data.Item.Method)
	case "path":
		color.Green(data.Item.Path)
	case "extension":
		color.Green(data.Item.Extension)
	case "status":
		color.Green(data.Item.Status)
	case "length":
		color.Green(data.Item.Responselength)
	case "mime":
		color.Green(data.Item.Mimetype)
	case "response":
		color.Red("Hasn't been implemented Yet")
		color.Green(data.Item.Response.Text)
	case "request":
		color.Red("Hasn't ben implemented Yet")
		color.Green(data.Item.Request.Text)
	case "comment":
		color.Green(data.Item.Comment)
	default:
		color.Red("The attribute provided was not found.")
	}

}

func EditInfo(attribute string, change string) {
	switch attribute {
	case "time":
		color.Red("Export time is Static.")
	case "url":
		data.Item.Host.Text = change
		color.Yellow(data.Item.Host.Text)
	case "host":
		data.Item.Host.Text = change
		color.Yellow(data.Item.Host.Text)
	case "ip":
		data.Item.Host.Ip = change
		color.Yellow(data.Item.Host.Ip)
	case "port":
		data.Item.Port = change
		color.Yellow(data.Item.Port)
	case "protocol":
		data.Item.Protocol = change
		color.Yellow(data.Item.Protocol)
	case "method":
		data.Item.Method = change
		color.Yellow(data.Item.Method)
	case "path":
		data.Item.Path = change
		color.Yellow(data.Item.Path)
	case "extension":
		data.Item.Extension = change
		color.Yellow(data.Item.Extension)
	case "status":
		color.Red("Response Status is Static.")
	case "length":
		color.Red("Response Length is Static.")
	case "mime":
		data.Item.Extension = change
		color.Yellow(data.Item.Mimetype)
	case "response":
		color.Red("Response Payload is Static.")
	case "request":
		color.Red("Hasn't been implemented yet.")
		color.Yellow(data.Item.Request.Text)
	case "comment":
		data.Item.Comment = change
		color.Yellow(data.Item.Comment)
	default:
		color.Red("The attribute provided was not found.")
	}
}

func DefineInsertions() {

}

func AddAuth(auth string) {
	color.Red("Decode the Payload, add the auth string, and then re-encode")
	// temp := base64decode(data.Item.Request.Text)
	// temp = temp + auth string
	// data.Item.Request.Text = temp
}

func Manual() {
	// while there are not two empty lines continual ask for Input from the user.

}
