package cli

import (
	// native Golang Support
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	// Third-Party Support
	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"

	// Support Files
	"github.com/ryanvillarreal/FuzzyWuzzyAPI/utils"
)

// call the shell function here.
func Shell(shellMenuContext string) {
	rl, err := readline.New("/" + filepath.Base(utils.CurrentDir) + ":~$ [" + shellMenuContext + "] ")
	if err != nil {
		panic(err) // you should get rid of panics.  handle all the exceptions
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
			nav(shellMenuContext, cmd, rl)
		}
	}
}

func nav(shellMenuContext string, cmd []string, rl *readline.Instance) {
	switch cmd[0] {

	case "Import", "import":
		shellMenuContext = "Import"
		rl.SetPrompt("/" + filepath.Base(utils.CurrentDir) + ":~$ [" + shellMenuContext + "] ")
		if len(cmd) > 1 {
			utils.BurpImport(cmd[1])
		} else {
			color.Red("Enter file name here.")
		}
		Shell("Main")

	case "Swagger", "swagger":
		color.Red("Not Implemented Yet")
	case "Auth", "auth":
		if len(cmd) > 1 {
			utils.AddAuth(cmd[1])
		} else {
			color.Red("Add the Auth here.")
		}
	case "Manual", "manual":
		shellMenuContext = "Manual"
		rl.SetPrompt("/" + filepath.Base(utils.CurrentDir) + ":~$ [" + shellMenuContext + "] ")

	case "Load", "load":
		if len(cmd) > 1 {
			utils.LoadPayloads(cmd[1])
		} else {
			color.Red("Need to specify the file name to load.")
		}

	case "Unload", "unload":
		utils.Unload()

	case "Download", "download":
		if len(cmd) > 1 {
			utils.DownloadPayloads(cmd[1], cmd[2])
		} else {
			utils.DownloadPayloads("./SecLists.zip", "https://github.com/danielmiessler/SecLists/archive/master.zip")
		}

	case "Proxy", "proxy":
		color.Green("Current Proxy: " + utils.PROXY_ADDR)

	case "View", "view":
		if len(cmd) > 1 {
			utils.PrintInfo(cmd[1])
		} else {
			color.Red("Use the View options and the attribute name")
		}
	case "Edit", "edit", "change", "Change":
		if len(cmd) > 2 {
			utils.EditInfo(cmd[1], cmd[2])
		} else {
			color.Red("Use the Edit command to change attribute data")
			color.Red("Available properties: url, ip, host, port, protocol, method, path, extension, mime, comment, user agent")

		}
	case "List", "list":
		utils.PrintInfo("all")

	case "Set", "set":
		if len(cmd) > 1 {
			utils.SetProxy(cmd[1])
		} else {
			color.Red("Specify address for Proxy support")
		}
	case "Unset", "unset":
		if len(cmd) > 1 {
			color.Red("No data needed.")
		} else {
			utils.UnsetProxy()
		}
	case "Test", "test":
		utils.SingleRequest()
	case "Threads", "threads":
		if len(cmd) > 1 {
			i, err := strconv.Atoi(cmd[1])
			if err != nil {
				color.Red("Something went wrong.")
			} else {
				utils.SetThreads(i)
			}
		} else {
			color.Yellow("Number of Threads Set: " + utils.GetThreads())
		}
	case "Attack", "attack":
		utils.Attack()
	case "Fuzz", "fuzz":
		shellMenuContext = "Fuzz"
		rl.SetPrompt("/" + filepath.Base(utils.CurrentDir) + ":~$ [" + shellMenuContext + "] ")
	case "Pause", "pause":
		color.Red("Not Implemented Yet")
	case "Resume", "resume":
		color.Red("Not Implemented Yet")
	case "Stats", "stats":
		color.Red("Not Implemented yet")
	case "Stop", "stop":
		color.Red("Not Implemented Yet")
	case "Output", "output":
		color.Red("not Implemented Yet")
	// Standard Operations
	case "exit", "Exit":
		exit()
	case "quit", "Quit":
		exit()
	case "menu", "Menu":
		menuHelpMain(shellMenuContext)
	case "help":
		menuHelpMain(shellMenuContext)
	case "?":
		menuHelpMain(shellMenuContext)
	case "back", "Back":
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
		{"Import", "Import a Burp Request", "<filename>"},
		{"Swagger", "Parse Swagger Documentation for further automation of API testing."},
		{"Auth", "Add an Auth payload into the request."},
		{"Manual", "Manually define the GET/POST request to Fuzz the API"},
		{"Load", "Load in new Payload lists to use with Fuzzing"},
		{"Unload", "Unload the payloads set previously."},
		{"Download", "Download the SecLists from Daniel Miessler", "download | download [filepath"},
		{"Proxy", "Set the configuration to use a proxy server"},
		{"View", "Interact with the Data that has been loaded in.", "view [attribute]"},
		{"Edit", "Edit the Data that has been loaded in.", "edit [attribute] [change]"},
		{"Change", "Same as edit parameter.", "change [attribute] [change]"},
		{"List", "List the attributes available from the Data that has been loaded"},
		{"Set", "Configure the Proxy Settings for FuzzyWuzzy"},
		{"Unset", "Unset the Proxy Settings for FuzzyWuzzy"},
		{"Attack", "Perform the attack with the current settings."},
		{"Test", "Test the current data in a Request.  Will return the response code. "},
		{"Threads", "Limit the Attacker to a number of threads.  Default 5.", "threads [int]"},
		{"Fuzz", "Perform Fuzzing against the Data that has been loaded in."},
		{"Pause", "Pause the current attacks."},
		{"Resume", "Resume the current attacks"},
		{"Stop", "Stop the current attacks"},
		{"Stats", "Get current stats of Requests/Responses"},
		{"Output", "Output results to a self-define location", "output [filename]"},
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
