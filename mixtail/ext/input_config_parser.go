package ext
import "github.com/tri-star/mixtail/mixtail/config"

type InputConfigParser interface {
	CreateInputConfigFromData(name string, data map[interface{}]interface{}) (entries []config.Input, err error)
}
