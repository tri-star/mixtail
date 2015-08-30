package config

import (
	"errors"
)

// InputSsh implements config entry that uses ssh.
type InputSsh struct {
	*InputBase

	Host string
	Port uint16
	User string
	Pass string
	Identity string
	Command string
}

// Returns new InputSsh.
func NewInputSsh() *InputSsh {
	b := new(InputBase)
	i := new(InputSsh)
	i.InputBase = b
	return i
}

// Initialize its self by given data.
func (is *InputSsh) BuildFromData(data map[interface{}]interface{}) (err error) {
	err = is.InputBase.BuildFromData(data)
	if err != nil {
		return
	}

	var ok bool
	is.Host, ok = data["host"].(string)
	if !ok {
		err = errors.New(is.Name + ": 'host' is not specified.")
		return
	}
	is.User, ok = data["user"].(string)
	if !ok {
		err = errors.New(is.Name + ": 'user' is not specified.")
		return
	}
	is.Command, ok = data["command"].(string)
	if !ok {
		err = errors.New(is.Name + ": 'command' is not specified.")
		return
	}

	//Optional fields
	port, ok := data["port"].(uint16)
	if ok {
		is.Port = port
	}
	pass, ok := data["pass"].(string)
	if ok {
		is.Pass = pass
	}
	identity, ok := data["identity"].(string)
	if ok {
		is.Identity = identity
	}

	return
}


//Create new InputSsh entry(s) from data.
//The data structure is corresponding to an entry of "input" section of YAML config.
func CreateSshConfigFromData(c *Config, name string, data map[interface{}]interface{}) (entries []*InputSsh, err error) {

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
	entries = make([]*InputSsh, 0)

	for _, hostName := range hosts {
		entry := NewInputSsh()

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
