package extssh

import (
	"errors"
	"github.com/tri-star/mixtail/lib"
	"github.com/tri-star/mixtail/mixtail/entity"
)

type InputEntryParser struct {
	*lib.BaseExtensionPoint
}

func NewInputEntryParser() (iep *InputEntryParser) {
	iep = new(InputEntryParser)
	iep.BaseExtensionPoint = new(lib.BaseExtensionPoint)
	iep.Name = EXTENSION_NAME
	return
}

func (iep *InputEntryParser) CreateInputEntriesFromData(name string, data map[interface{}]interface{}) (inputEntries []entity.InputEntry, err error) {
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
	inputEntries = make([]entity.InputEntry, 0)

	for _, hostName := range hosts {
		entry := NewInputEntry()

		for key, value := range data {
			workData[key] = value
		}
		workData["name"] = name
		workData["host"] = hostName
		err = entry.BuildFromData(workData)
		if err != nil {
			return
		}

		inputEntries = append(inputEntries, entry)
	}
	return
}
