package service
import (
	"github.com/tri-star/mixtail/lib"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"errors"
	"github.com/tri-star/mixtail/mixtail/entity"
	"github.com/tri-star/mixtail/mixtail/ext"
)


type Config struct {
	extensionManager *lib.ExtensionManager
}


func NewConfig(em *lib.ExtensionManager) (c *Config) {
	c = new(Config)
	c.extensionManager = em
	return
}


// Parse data from given file path.
func (c *Config) ParseFromFile(path string) (conf *entity.Config, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	conf, err = c.Parse(data)
	if err != nil {
		return
	}
	return
}

// Parse data from given byte slices.
func (c *Config) Parse(data []byte) (conf *entity.Config, err error) {
	var parseResult interface{}
	err = yaml.Unmarshal(data, &parseResult)
	if err != nil{
		return
	}
	conf = entity.NewConfig()

	// Get "input" section.
	inputSettingSection, ok := parseResult.(map[interface{}]interface{})["input"]
	if !ok {
		err = errors.New("'input' section not found.")
		return
	}
	inputSettingEntries, _ := inputSettingSection.(map[interface{}]interface{})
	for name, entry := range inputSettingEntries {
		err = c.parseInputSectionEntry(conf, name.(string), entry.(map[interface{}]interface{}))
		if err != nil {
			return
		}

	}

	logSection, ok := parseResult.(map[interface{}]interface{})["log"].(map[interface{}]interface{})
	if ok {
		err = c.parseLogSection(conf, logSection)
		if err != nil {
			return
		}
	}

	return
}

// Internal function that parses "input" section entries.
func (c *Config) parseInputSectionEntry(conf *entity.Config, name string, entry map[interface{}]interface{}) (err error) {
	typeName, ok := entry["type"].(string)
	if !ok {
		err = errors.New(name + ": " + "'type' is not specified")
		return
	}

	extension, found := c.extensionManager.GetExtensionPoint(ext.POINT_INPUT_CONFIG_PARSER, typeName)
	if !found {
		err = errors.New(name + ": " + "invalid type '" + typeName + "' specified.")
		return
	}
	inputEntries, err := extension.(ext.InputEntryParser).CreateInputEntriesFromData(conf, name, entry)
	if err != nil {
		return
	}
	conf.InputEntries = append(conf.InputEntries, inputEntries...)
	return
}


// Parse "log" section.
func (c *Config) parseLogSection(conf *entity.Config, section map[interface{}]interface{}) (err error) {
	logging, ok := section["logging"].(bool)
	if !ok || !logging {
		conf.Logging = false
		return
	}
	conf.Logging = true

	logPath, ok := section["path"].(string)
	if !ok || len(logPath) == 0 {
		err = errors.New("'path' must be specified.")
		return
	}
	conf.LogPath = logPath
	return
}
