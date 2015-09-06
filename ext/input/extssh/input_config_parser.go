package extssh

import (
	"github.com/tri-star/mixtail/config"
	"errors"
)

const (
	EXTENSION_TYPE = config.EXTENSION_TYPE_INPUT_CONFIG_PARSER + "-ssh"
)

type InputConfigParser struct {
	*config.BaseInputConfigParser
}

func NewInputConfigParser() (icp *InputConfigParser) {
	icp = new(InputConfigParser)
	icp.BaseInputConfigParser = new(config.BaseInputConfigParser)
	icp.Name = EXTENSION_TYPE
	return
}

func (ic *InputConfigParser) GetName() string {
	return ic.Name
}

func (ic *InputConfigParser) CreateInputConfigFromData(name string, data map[interface{}]interface{}) (inputConfigs []config.Input, err error) {
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
	inputConfigs = make([]config.Input, 0)

	for _, hostName := range hosts {
		entry := NewInputConfig()

		for key, value := range data {
			workData[key] = value
		}
		workData["name"] = name
		workData["host"] = hostName
		err = entry.BuildFromData(workData)
		if err != nil {
			return
		}

		inputConfigs = append(inputConfigs, entry)
	}
	return
}
