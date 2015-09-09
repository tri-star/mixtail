package mixtail
import (
	"fmt"
	"log"
	"flag"
	"os"
	"errors"
	"bytes"
	"path/filepath"
	"github.com/tri-star/mixtail/lib"
	"github.com/tri-star/mixtail/mixtail/entity"
	"github.com/tri-star/mixtail/mixtail/service"
	"github.com/tri-star/mixtail/mixtail/ext"
)

// Application class.
type Application struct {

	lineDelimiter []byte
	config *entity.Config
	logFile *os.File

	extensionManager *lib.ExtensionManager
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


// Returns Application singleton instance.
func NewApplication(em *lib.ExtensionManager) *Application {
	app := new(Application)
	app.lineDelimiter = []byte("\n")
	app.extensionManager = em
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
func (a *Application) loadConfig(configPath string) (c *entity.Config, err error) {
	configService := service.NewConfig(a.extensionManager)
	c, err = configService.ParseFromFile(configPath)
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

  # Example for watching command across multiple hosts.
  log_name02:
    type: ssh
    #
    host:
      - 192.168.1.10
      - 192.168.1.11
      - 192.168.1.12
      - 192.168.1.13
    # Multi line command is also supported.
    command: |
      A=aaa
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
	outputData := make(chan entity.InputData, 0)

	//Initialize input handlers.
	inputHandlers := make(map[string]ext.InputHandler)
	endFlagList := make(map[string]bool)
	for _, inputEntry := range a.config.InputEntries {
		inputHandlerFactory, found := a.extensionManager.GetExtensionPoint(ext.POINT_INPUT_HANDLER_FACTORY, inputEntry.GetType())
		if !found {
			panic(fmt.Sprintf("invalid input handler: %s", inputEntry.GetType()))
		}
		inputHandler := inputHandlerFactory.(ext.InputHandlerFactory).NewInputHandler(inputEntry)
		inputHandlers[inputHandler.GetName()] = inputHandler
	}

	//Start listen.
	//Handler runs asynchronously in goroutine.
	for name, ih := range inputHandlers {
		go ih.ReadInput(outputData)
		endFlagList[name] = false
	}

	//Wait outputData channel receive a data.
	//This loop continues until all input handler exits.
	var inputData entity.InputData
	allEndFlag := false
	for {
		inputData = <- outputData

		if inputData.State == entity.INPUT_DATA_END {
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
