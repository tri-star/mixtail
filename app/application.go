package app
import (
	"fmt"
	"github.com/tri-star/mixtail/config"
	"github.com/tri-star/mixtail/handler"
	"log"
	"flag"
	"os"
	"errors"
)

// Application class.
type Application struct {


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


func GetInstance() *Application {
	if(app != nil) {
		return app
	}
	return new(Application)
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


func (a *Application) Run() {

	//コマンドラインの解析
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
`)
}


func (a *Application) mainCommand(options *StartupOptions) {
	//設定情報のロード
	conf, err := a.loadConfig(options.ConfigFile)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//出力用チャンネルの作成
	//(今はチャネルは1つなのでここで作成)
	outputData := make(chan handler.InputData, 0)

	//入力ソース分初期化
	inputHandlers := make(map[string]handler.InputHandler)
	endFlagList := make(map[string]bool)
	for _, inputConfig := range conf.Inputs {
		//入力ハンドラの初期化

		inputHandler, err := handler.NewInputHandler(inputConfig)
		if err != nil {
			panic(fmt.Sprintf("error: %s, message: %s", inputConfig.GetName(), err.Error()))
		}
		inputHandlers[inputConfig.GetName()] = inputHandler
	}

	for name, ih := range inputHandlers {
		go ih.ReadInput(outputData)
		endFlagList[name] = false
	}

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
		fmt.Printf("[%s] %s\n", inputData.Name, inputData.Data)
	}

}

