package cli

import (
	// native Golang Support
	"fmt"
	"github.com/ryanvillarreal/Slackord/pkg/core"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	// Third-Party Support
	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// call the shell function here.
func Shell(shellMenuContext string) {
	rl, err := readline.New("/" + filepath.Base(core.CurrentDir) + ":~$ ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		line = strings.TrimSpace(line)
		cmd := strings.Fields(line)

		if len(cmd) > 0 {
			// main CLI logic here.
			switch shellMenuContext {
			case "main":
				switch cmd[0] {
				case "help":
					menuHelpMain()
				case "?":
					menuHelpMain()
				case "exit":
					exit()
				case "quit":
					exit()
				case "menu":
					menuHelpMain()
				default:
					message("info", "Executing system command...")
					if len(cmd) > 1 {
						executeCommand(cmd[0], cmd[1:])
					} else {
						var x []string
						executeCommand(cmd[0], x)
					}
				}
			case "agent":
				fmt.Println("You are operating on the agent side.")
			case "module":
				fmt.Println("You are inside the module menu")
			case "listener":
				fmt.Println("You are inside the listener menu")
			}
		}
	}
}

// Message is used to print a message to the command line
func message(level string, message string) {
	switch level {
	case "info":
		color.Cyan("[i]" + message)
	case "note":
		color.Yellow("[-]" + message)
	case "warn":
		color.Red("[!]" + message)
	case "debug":
		color.Red("[DEBUG]" + message)
	case "success":
		color.Green("[+]" + message)
	default:
		color.Red("[_-_]Invalid message level: " + message)
	}
}

// kill the server and the CLI
func exit() {
	color.Red("[!]Quitting")
	os.Exit(0)
}

// execute a local command from inside the CLI
func executeCommand(name string, arg []string) {
	var cmd *exec.Cmd

	cmd = exec.Command(name, arg...)

	out, err := cmd.CombinedOutput()

	if err != nil {
		message("warn", err.Error())
	} else {
		message("success", fmt.Sprintf("%s", out))
	}
}

// prints the main menu when called. Can be used for help or any situation with a bad command line option
func menuHelpMain() {
	color.Yellow("Slackord - Help Menu")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)
	table.SetCaption(true, "Main Menu Help")
	table.SetHeader([]string{"Command", "Description", "Options"})

	data := [][]string{
		{"exit", "Exit and close the Slackord server", ""},
		{"quit", "Exit and close the Slackord server", ""},
		{"*", "Anything else will be executed on the host operating system", ""},
	}
	table.AppendBulk(data)
	fmt.Println()
	table.Render()
	fmt.Println()
}

func getDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Directory inaccessible")
	}
	return dir
}
