package ssh
import (
	"github.com/tri-star/mixtail/config"
	"errors"
	"github.com/tri-star/mixtail/ext"
)


type InputConfigExtension struct {
	name string
}

func NewInputConfigExtension() (ic *InputConfigExtension) {
	ic = new(InputConfigExtension)
	ic.name = ext.INPUT_CONFIG_SSH
	return ic
}

func (ic *InputConfigExtension) Name() string {
	return ic.Name
}

func (ic *InputConfigExtension) CreateInputConfigFromData(name string, data map[interface{}]interface{}) (entries []config.Input, err error) {
	hosts := make([]string, 0)
	//hostパラメータを調べる。1件か、配列か？
	hostAsString, ok := data["host"].(string)
	if ok {
		hosts = append(hosts, hostAsString)
	} else {
		hostsAsMap, ok := data["host"].([]interface{})
		if ok {
			for _, hostNameBeforeCast := range hostsAsMap {
				hostName, ok := hostNameBeforeCast.(string)
				if !ok {
					err = errors.New("Invaid host name was found while parsing host section.")
					return
				}
				hosts = append(hosts, hostName)
			}
		}
	}

	//host件数分ループ
	workData := make(map[interface{}]interface{})
	entries = make([]config.Input, 0)

	for _, hostName := range hosts {
		entry := config.NewInputSsh()

		for key, value := range data {
			workData[key] = value
		}
		workData["name"] = name
		workData["host"] = hostName
		err = entry.BuildFromData(workData)
		if err != nil {
			return
		}

		entries = append(entries, entry)
	}
	return
}
