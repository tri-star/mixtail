package app
import (
	"fmt"
	"github.com/tri-star/mixtail/config"
	"github.com/tri-star/mixtail/handler"
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
	outputData := make(chan []byte)

	//入力ソース分初期化
	inputHandlers := make(map[string]handler.InputHandler, 0, 10)
	for _, inputConfig := range conf.Inputs {
		//入力ハンドラの初期化

		inputHandler, err := handler.NewInputHandler(inputConfig)
		if err != nil {
			panic(fmt.Sprintf("error: %s, message: %s", inputConfig.GetName(), err.Error()))
		}
		inputHandlers[inputConfig.GetName()] = inputHandler
	}

	for _, ih := range inputHandlers {
		go ih.ReadInput(outputData)
	}

	wholeDone := false
	var readData []byte
	for wholeDone == false {
		readData <- outputData
		fmt.Println(readData)

		wholeDone = true
		for _, ih := range inputHandlers {
			if(ih.Status() != handler.INPUT_STATUS_DONE) {
				wholeDone = false
			}
		}
	}

}


func (a *Application) loadConfig() (c *config.Config) {
	c = config.NewConfig()

	rc := config.NewInputRemote()
	rc.Name = ""
	rc.Type = "ssh"
	rc.Host = "urban-theory.net"
	rc.User = "hiroki"
	rc.Identity = "/home/hiroki/.ssh/aaaaa"
	rc.Command = "tail -f /tmp/test.txt"

	append(c.Inputs, rc)
	return
}
