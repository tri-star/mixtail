package app
import (
	"../config"
	"../handler/input"
)


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
	//outputData := make(chan []byte)

	//入力ソース分初期化
	//inputHandlers := make(map[string]*input.InputHandler, 0, 10)
	//for c := range conf.Inputs {
		//入力ハンドラの初期化


		//出力ハンドラの初期化
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
