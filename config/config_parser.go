package config

import (
	"gopkg.in/yaml.v2"

	"errors"
	"io/ioutil"
)

// ConfigParser is parse YAML data and populate it into Config object.
type ConfigParser struct {

	config *Config
}


// Returns new ConfigParser.
func NewConfigParser() (cp *ConfigParser) {
	cp = new(ConfigParser)
	cp.config = NewConfig()
	return
}

// Returns parse result(Config).
func (cp *ConfigParser) GetResult() *Config {
	return cp.config
}

// Parse data from given file path.
func (cp *ConfigParser) ParseFromFile(path string) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = cp.Parse(data)
	if err != nil {
		return
	}
	return
}

// Parse data from given byte slices.
func (cp *ConfigParser) Parse(data []byte) (err error) {
	var parseResult interface{}
	err = yaml.Unmarshal(data, &parseResult)
	if err != nil{
		return err
	}

	// Get "input" section.
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

	logSection, ok := parseResult.(map[interface{}]interface{})["log"].(map[interface{}]interface{})
	if ok {
		err = cp.parseLogSection(cp.config, logSection)
		if err != nil {
			return
		}
	}

	return
}

// Internal function that parses "input" section of YAML data.
func (cp *ConfigParser) parseInputHandler(name string, entry map[interface{}]interface{}) (newEntry Input, err error) {
	typeName, ok := entry["type"].(string)
	if !ok {
		err = errors.New(name + ": " + "'type' is not specified")
		return
	}

	// Parse section according to "type" key.
	switch(typeName) {
	case INPUT_TYPE_SSH:
		newEntry = NewInputSsh()
		err = newEntry.BuildFromData(name, entry)
	default:
		err = errors.New(name + ": " + "invalid type '" + typeName + "' specified.")
		return
	}

	return
}


// Parse "log" section.
func (cp *ConfigParser) parseLogSection(c *Config, section map[interface{}]interface{}) (err error) {
	logging, ok := section["logging"].(bool)
	if !ok || !logging {
		c.Logging = false
		return
	}
	c.Logging = true

	logPath, ok := section["path"].(string)
	if !ok || len(logPath) == 0 {
		err = errors.New("'path' must be specified.")
		return
	}
	c.LogPath = logPath
	return
}
