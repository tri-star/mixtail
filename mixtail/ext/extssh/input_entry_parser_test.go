package extssh_test

import (
	"testing"
	"gopkg.in/yaml.v2"
	"github.com/tri-star/mixtail/mixtail/ext/extssh"
	"github.com/tri-star/mixtail/mixtail/entity"
)


func TestCreateInputConfigFromData(t *testing.T) {

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

	inputSection := yamlData.(map[interface{}]interface{})["input"].(map[interface{}]interface{})

	inputEntryParser := extssh.NewInputEntryParser()

	var entries []entity.InputEntry
	var tempEntries []entity.InputEntry
	config := entity.NewConfig()
	for name, inputEntry := range inputSection {
		input := inputEntry.(map[interface{}]interface{})

		tempEntries, err = inputEntryParser.CreateInputEntriesFromData(config, name.(string), input)
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
		entry := entries[i].(*extssh.InputEntry)

		if entry.Host != expected["host"] {
			t.Logf("Host is not expected. index=%d: value=%s", i, entry.Host)
			t.Fail()
		}
		if entry.Name != expected["name"] {
			t.Logf("Log name is not expected. index=%d: value=%s", i, entry.Name)
			t.Fail()
		}
		if entry.Cred.User != expected["user"] {
			t.Logf("user is not expected. index=%d: value=%s", i, entry.Cred.User)
			t.Fail()
		}
	}

}



func TestCredentialOmitted(t *testing.T) {

	yamlString := []byte(`input:
  logA:
    type: ssh
    host:
      - 1.1.1.1
      - 2.2.2.2
    command: abc
  logB:
    type: ssh
    host: 3.3.3.3
    command: abc
`)

	var yamlData interface{}
	err := yaml.Unmarshal(yamlString, &yamlData)
	if err != nil {
		t.Log("Failed to parse yaml.")
		t.Fail()
		return
	}

	inputSection := yamlData.(map[interface{}]interface{})["input"].(map[interface{}]interface{})

	inputEntryParser := extssh.NewInputEntryParser()

	var entries []entity.InputEntry
	var tempEntries []entity.InputEntry
	config := entity.NewConfig()
	config.DefaultCredential.User = "test"
	config.DefaultCredential.Pass = "test"
	config.Hosts["3.3.3.3"] = entity.NewCredential()
	config.Hosts["3.3.3.3"].User = "test2"
	config.Hosts["3.3.3.3"].Pass = "test2"

	for name, inputEntry := range inputSection {
		input := inputEntry.(map[interface{}]interface{})

		tempEntries, err = inputEntryParser.CreateInputEntriesFromData(config, name.(string), input)
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
		entry := entries[i].(*extssh.InputEntry)

		if entry.Host != expected["host"] {
			t.Logf("Host is not expected. index=%d: value=%s", i, entry.Host)
			t.Fail()
		}
		if entry.Name != expected["name"] {
			t.Logf("Log name is not expected. index=%d: value=%s", i, entry.Name)
			t.Fail()
		}
		if entry.Cred.User != expected["user"] {
			t.Logf("user is not expected. index=%d: value=%s", i, entry.Cred.User)
			t.Fail()
		}
	}

}
