package cli

import (
	// native Golang Support
	"fmt"
	"github.com/ryanvillarreal/FuzzyWuzzyAPI/cli/utils"
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
	rl, err := readline.New("/" + filepath.Base(utils.CurrentDir) + ":~$ [" + shellMenuContext + "] ")
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
			case "Main", "main":
				switch cmd[0] {
				case "help":
					menuHelpMain(shellMenuContext)
				case "?":
					menuHelpMain(shellMenuContext)
				case "Burp", "burp":
					Shell("Burp")
				case "Manual", "manual":
					Shell("Manual")
				case "Proxy", "proxy":
					Shell("Proxy")
				case "exit", "Exit":
					exit()
				case "quit", "Quit":
					exit()
				case "menu", "Menu":
					menuHelpMain(shellMenuContext)
				default:
					message("info", "Executing system command...")
					if len(cmd) > 1 {
						executeCommand(cmd[0], cmd[1:])
					} else {
						var x []string
						executeCommand(cmd[0], x)
					}
				}
			case "Burp", "burp":
				switch cmd[0] {
				case "import", "Import":
					if len(cmd) > 1 {
						utils.BurpRequest(cmd[1])
					} else {
						color.Red("Pass the file here.")
					}
				case "help":
					menuHelpBurp(shellMenuContext)
				case "?":
					menuHelpBurp(shellMenuContext)
				case "exit":
					exit()
				case "quit":
					exit()
				case "menu":
					menuHelpBurp(shellMenuContext)
				case "back":
					Shell("Main")
				default:
					message("info", "Executing system command...")
					if len(cmd) > 1 {
						executeCommand(cmd[0], cmd[1:])
					} else {
						var x []string
						executeCommand(cmd[0], x)
					}
				}
			case "Manual", "manual":
				switch cmd[0] {
				case "help":
					menuHelpMain(shellMenuContext)
				case "?":
					menuHelpMain(shellMenuContext)
				case "exit":
					exit()
				case "quit":
					exit()
				case "menu":
					menuHelpMain(shellMenuContext)
				case "back":
					Shell("Main")
				default:
					message("info", "Executing system command...")
					if len(cmd) > 1 {
						executeCommand(cmd[0], cmd[1:])
					} else {
						var x []string
						executeCommand(cmd[0], x)
					}
				}

			case "Proxy", "proxy":
				switch cmd[0] {
				case "help":
					menuHelpMain(shellMenuContext)
				case "?":
					menuHelpMain(shellMenuContext)
				case "exit":
					exit()
				case "quit":
					exit()
				case "menu":
					menuHelpMain(shellMenuContext)
				case "back":
					Shell("Main")
				default:
					message("info", "Executing system command...")
					if len(cmd) > 1 {
						executeCommand(cmd[0], cmd[1:])
					} else {
						var x []string
						executeCommand(cmd[0], x)
					}
				}
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
func menuHelpMain(context string) {
	color.Yellow(context + " - Help Menu")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)
	table.SetHeader([]string{"Command", "Description", "Options"})

	data := [][]string{
		{"Burp", "Enter the Burp context menu. You can use Burp parsing to open requests."},
		{"Manual", "Manually define the GET/POST request to Fuzz the API"},
		{"Load", "Load in new Payload lists to use with Fuzzing"},
		{"Proxy", "Set the configuration to use a proxy server"},
		{"exit", "Exit and close the FuzzyWuzzy server", ""},
		{"quit", "Exit and close the FuzzyWuzzy server", ""},
		{"*", "Anything else will be executed on the host operating system", ""},
	}
	table.AppendBulk(data)
	fmt.Println()
	table.Render()
	fmt.Println()
}

func menuHelpBurp(context string) {
	color.Yellow(context + " - Help Menu")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)
	table.SetHeader([]string{"Command", "Description", "Options"})

	data := [][]string{
		{"Import", "Pass the name of the file into the import function"},
		{"Proxy", "Set the configuration to use a proxy server"},
		{"Back", "Return to the main menu"},
		{"exit", "Exit and close the FuzzyWuzzy server", ""},
		{"quit", "Exit and close the FuzzyWuzzy server", ""},
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
