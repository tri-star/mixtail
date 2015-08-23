package config

import (
	"gopkg.in/yaml.v2"

	"errors"
	"io/ioutil"
)

type ConfigParser struct {

	config *Config
}


func NewConfigParser() (cp *ConfigParser) {
	cp = new(ConfigParser)
	cp.config = NewConfig()
	return
}

func (cp *ConfigParser) GetResult() *Config {
	return cp.config
}

func (cp *ConfigParser) ParseFromFile(path string) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	cp.Parse(data)
	return
}

func (cp *ConfigParser) Parse(data []byte) (err error) {
	var parseResult interface{}
	err = yaml.Unmarshal(data, &parseResult)
	if err != nil{
		return err
	}

	inputSettingSection, ok := parseResult.(map[interface{}]interface{})["input"]
	if !ok {
		return errors.New("'input' section not found.")
	}
	inputSettingEntries, _ := inputSettingSection.(map[interface{}]interface{})
	for name, entry := range inputSettingEntries {
		var newEntry Input

		newEntry, err = cp.parseInputHandler(name.(string), entry.(map[interface{}]interface{}))
		if err != nil {
			return
		}
		cp.config.Inputs = append(cp.config.Inputs, newEntry)
	}

	return
}


func (cp *ConfigParser) parseInputHandler(name string, entry map[interface{}]interface{}) (newEntry Input, err error) {
	typeName, ok := entry["type"].(string)
	if !ok {
		err = errors.New(name + ": " + "'type' is not specified")
		return
	}

	switch(typeName) {
	case INPUT_TYPE_SSH:
		newEntry = NewInputRemote()
		err = newEntry.BuildFromData(name, entry)
	default:
		err = errors.New(name + ": " + "invalid type '" + typeName + "' specified.")
		return
	}

	return
}
