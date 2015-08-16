package app
import (
	"fmt"
	"github.com/tri-star/mixtail/config"
	"github.com/tri-star/mixtail/handler"
	"log"
)

// Application class.
type Application struct {


}

var app *Application


func GetInstance() *Application {
	if(app != nil) {
		return app
	}
	return new(Application)
}


func (a *Application) Run() {

	//コマンドラインの解析

	//設定情報のロード
	conf := a.loadConfig()

	//出力用チャンネルの作成
	//(今はチャネルは1つなのでここで作成)
	outputData := make(chan *handler.InputData)

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

	var inputData *handler.InputData
	allEndFlag := false
	for {
		inputData = <-outputData
		fmt.Printf("%s\n", inputData.Data)

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
	}

}


func (a *Application) loadConfig() (c *config.Config) {
	c = config.NewConfig()

	rc := config.NewInputRemote()
	rc.Name = "test01"
	rc.Type = "ssh"
	rc.Host = "example.com"
	rc.User = "test"
	rc.Identity = "/home/test/test"
	rc.Command = "tail -f /tmp/test.txt"

	c.Inputs = append(c.Inputs, rc)
	return
}
