package app
import (
	"fmt"
	"github.com/tri-star/mixtail/config"
	"github.com/tri-star/mixtail/handler"
	"log"
	"flag"
	"os"
	"errors"
	"bytes"
	"path/filepath"
)

// Application class.
type Application struct {

	lineDelimiter []byte
	config *config.Config
	logFile *os.File
}

type StartupOptions struct {
	ConfigFile string
	Command uint8
}

const (
	COMMAND_MAIN = iota
	COMMAND_EXAMPLE
	COMMAND_VERSION
	COMMAND_HELP
)

const DEFAULT_CONFIG_FILE_NAME="config.yml"


var app *Application


// Returns Application singleton instance.
func GetInstance() *Application {
	if(app != nil) {
		return app
	}
	app = new(Application)
	app.lineDelimiter = []byte("\n")
	return app
}


func (a *Application) GetUsage() string {
	return `
Usage: mixtail [options] config-file.yml

options:
  --example: Print an example of config file.
  --version: Show version.
  --help:    Show this help.
`
}


func (a *Application) PrintUsage() {
	fmt.Print(a.GetUsage())
}


// Run() executes command line option parsing and dispatch corresponding command.
func (a *Application) Run() {

	options, err := a.parseStartupOptions(os.Args[1:])
	if err != nil {
		fmt.Println(err.Error())
		a.PrintUsage()
		return
	}

	switch(options.Command) {
	case COMMAND_EXAMPLE:
		a.printExampleCommand(options)
		return

	case COMMAND_HELP:
		a.printHelpCommand(options)
		return

	case COMMAND_VERSION:
		a.printVersionCommand(options)
		return

	default:
		a.mainCommand(options)
	}

}

// Parse command line option and populate it into StartupOptions.
func (a *Application) parseStartupOptions(args []string) (options *StartupOptions, err error){
	options = new(StartupOptions)

	flagSet := flag.NewFlagSet("", flag.ExitOnError)
	flagSet.Usage = a.PrintUsage

	exampleFlag := flagSet.Bool("example", false, "")
	versionFlag := flagSet.Bool("version", false, "")

	err = flagSet.Parse(args)
	if err != nil {
		return
	}

	if *versionFlag {
		options.Command = COMMAND_VERSION
		return
	}
	if *exampleFlag {
		options.Command = COMMAND_EXAMPLE
		return
	}

	if len(args) < 1 {
		err = errors.New("No file name specified.")
		return
	}

	options.ConfigFile = args[len(args)-1]
	return
}


// Load config file(YAML) and populate it into config.Config.
func (a *Application) loadConfig(configPath string) (c *config.Config, err error) {
	configParser := config.NewConfigParser()
	err = configParser.ParseFromFile(configPath)
	if err != nil {
		return
	}

	c = configParser.GetResult()
	return
}


func (a *Application) printVersionCommand(options *StartupOptions) {
	fmt.Println(Version)
}


func (a *Application) printHelpCommand(options *StartupOptions) {
	fmt.Println(a.GetUsage())
}


func (a *Application) printExampleCommand(options *StartupOptions) {
	fmt.Println(`# This is a sample configration file for mixtail.

input:
  # Example for remote command watching.
  # 'log_name01' is log name. It is used for identify log data.
  log_name01:
    type: ssh
    host: example.com
    user: user_name
    # Either pass or identity required.
    identity: /home/user_name/.ssh/id_rsa
    # pass: password
    command: tail -f /tmp/some_file

  # Example for local command watching.
  log_name02:
    type: local
    # Multiline command is supported.
    command: |
      export A=aaa
      echo $A

log:
  logging: false
  path: /tmp/tail.log
`)
}

// This function is application main procedure.
// Open input handlers and read them output until all input handler ends.
func (a *Application) mainCommand(options *StartupOptions) {
	var err error
	a.config, err = a.loadConfig(options.ConfigFile)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//Prepare logging if needed.
	if a.config.Logging {
		err = a.initLogging(a.config.LogPath)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	//Create output channel.
	//All input handler communicate with this channel.
	outputData := make(chan handler.InputData, 0)

	//Initialize input handlers.
	inputHandlers := make(map[string]handler.InputHandler)
	endFlagList := make(map[string]bool)
	for _, inputConfig := range a.config.Inputs {
		inputHandler, err := handler.NewInputHandler(inputConfig)
		if err != nil {
			panic(fmt.Sprintf("error: %s, message: %s", inputConfig.GetName(), err.Error()))
		}
		inputHandlers[inputConfig.GetName()] = inputHandler
	}

	//Start listen.
	//Handler runs asynchronously in goroutine.
	for name, ih := range inputHandlers {
		go ih.ReadInput(outputData)
		endFlagList[name] = false
	}

	//Wait outputData channel receive a data.
	//This loop continues until all input handler exits.
	var inputData handler.InputData
	allEndFlag := false
	for {
		inputData = <- outputData

		if inputData.State == handler.INPUT_DATA_END {
			endFlagList[inputData.Name] = true
		}

		allEndFlag = true
		for _, isEnd := range endFlagList {
			if !isEnd {
				allEndFlag = false
			}
		}
		if allEndFlag {
			log.Println("all handler stopped correctly.")
			break
		}

		//Output read data line by line.
		a.outputLines(inputData.Name, inputData.Data)
	}

}

// Output line with prefix string.
func (a *Application) outputLines(prefix string, data []byte) {
	lines := bytes.Split(data, a.lineDelimiter)

	for _, line := range lines[:len(lines)-1] {
		formattedLine := fmt.Sprintf("[%s] %s\n", prefix, line)
		fmt.Print(formattedLine)
		if a.config.Logging {
			a.logFile.WriteString(formattedLine)
		}
	}
}


func (a *Application) initLogging(logPath string) (err error) {
	logDir := filepath.Dir(logPath)
	err = os.MkdirAll(logDir, 0775)
	if err != nil {
		return
	}

	a.logFile, err = os.Create(logPath)
	if err != nil {
		return
	}

	return
}
