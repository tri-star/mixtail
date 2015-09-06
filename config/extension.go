package config

import "github.com/tri-star/mixtail/ext"

const (
	EXTENSION_TYPE_INPUT_CONFIG_PARSER="input-config-parser"
)

// InputConfigParser is one of an extension.
// It converts "input" section of YAML data into config.Input.
type InputConfigParser interface {
	ext.Extension
	CreateInputConfigFromData(name string, data map[interface{}]interface{}) (entries []Input, err error)
}

type BaseInputConfigParser struct {
	Name string
}

func (bic *BaseInputConfigParser) GetName() string{
	return bic.Name
}

func (bic *BaseInputConfigParser) GetType() string{
	return EXTENSION_TYPE_INPUT_CONFIG_PARSER
}

