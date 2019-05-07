package utils

import (
	"archive/zip"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// define the global array for Payloads
var payloads []string

func Unload() {
	payloads = nil
	color.Green("Total Payloads: " + PrintPayloadLen())
}

// import the Payload files here.
func LoadPayloads(filename string) {
	color.Yellow("Loading file: " + filename)
	file, err := os.Open(filename)
	if err != nil {
		color.Red(err.Error())
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			payloads = append(payloads, scanner.Text())
		}
	}
	color.Green("Total Payloads: " + PrintPayloadLen())
}

func PrintPayloadLen() (length string) {
	return strconv.Itoa(len(payloads))
}

func GetPayloadLen() (length int) {
	return len(payloads)
}

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer
// interface and we can pass this into io.TeeReader() which will report progress on each
// write cycle.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	color.Green("\rDownloading.. %s complete", Bytes(wc.Total))
}

func DownloadPayloads(filepath string, url string) {
	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath)
	if err != nil {
		color.Red(err.Error())
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		color.Red(err.Error())
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		color.Red(err.Error())
	}

	// The progress uses the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	// Now unzip it.
	// maybe ask first.
	Unzip(filepath)
}

// Function to unzip the contents of the recently downloaded files.
func Unzip(file string) {
	zipReader, _ := zip.OpenReader(file)
	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			color.Red(err.Error())
		}
		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			color.Green("Directory Created:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			color.Green("File extracted:" + file.Name)

			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				color.Red(err.Error())
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				color.Red(err.Error())
			}
		}
	}
}
