package config_test


import (
	"testing"
	"gopkg.in/yaml.v2"
	"github.com/tri-star/mixtail/config"
)


func TestCreateConfigFromData(t *testing.T) {

	c := config.NewConfig()

	yamlString := []byte(`input:
  logA:
    type: ssh
    host:
      - 1.1.1.1
      - 2.2.2.2
    user: test
    pass: password
    command: abc
  logB:
    type: ssh
    host: 3.3.3.3
    user: test2
    pass: password2
    command: abc
`)

	var yamlData interface{}
	err := yaml.Unmarshal(yamlString, &yamlData)
	if err != nil {
		t.Log("Failed to parse yaml.")
		t.Fail()
		return
	}

	inputSection, ok := yamlData.(map[interface{}]interface{})["input"].(map[interface{}]interface{})
	if !ok {
		t.Log("Failed to find input section.")
		t.Fail()
		return
	}

	var entries []*config.InputSsh
	var tempEntries []*config.InputSsh
	for name, inputEntry := range inputSection {
		input, ok := inputEntry.(map[interface{}]interface{})
		if !ok {
			t.Log("Failed to find input entry.")
			t.Fail()
			return
		}

		tempEntries, err = config.CreateSshConfigFromData(c, name.(string), input)
		entries = append(entries, tempEntries...)
	}

	if len(entries) != 3 {
		t.Logf("Config entry count does not matched. expected: 3, actual: %d", len(entries))
		t.Fail()
		return
	}

	expects := []map[string]string{
		{"host": "1.1.1.1", "name": "logA", "user": "test"},
		{"host": "2.2.2.2", "name": "logA", "user": "test"},
		{"host": "3.3.3.3", "name": "logB", "user": "test2"},
	}

	for i, expected := range expects {
		entry := entries[i]

		if entry.Host != expected["host"] {
			t.Logf("Host is not expected. index=%d: value=%s", i, entry.Host)
			t.Fail()
		}
		if entry.Name != expected["name"] {
			t.Logf("Log name is not expected. index=%d: value=%s", i, entry.Name)
			t.Fail()
		}
		if entry.User != expected["user"] {
			t.Logf("user is not expected. index=%d: value=%s", i, entry.User)
			t.Fail()
		}
	}

}
