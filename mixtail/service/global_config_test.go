package service_test
import (
	"testing"
	"github.com/tri-star/mixtail/mixtail/service"
	"gopkg.in/yaml.v2"
	"github.com/tri-star/mixtail/mixtail/entity"
)


func TestGlobalConfigParse(t *testing.T) {
	globalConfigService := service.NewGlobalConfig()

	yamlData := []byte(`default:
  user: default-user
  pass: default-pass
  identity: default-identity

host:
  host-a:
    user: host-a-user
    pass: host-a-pass
    identity: host-a-identity
  host-b:
    user: host-b-user
    pass: host-b-pass
    identity: host-b-identity
`)

	var data map[interface{}]interface{}
	data = make(map[interface{}]interface{})
	err := yaml.Unmarshal(yamlData, data)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		return
	}

	config := entity.NewConfig()
	err = globalConfigService.Parse(data, config)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		return
	}

	if config.DefaultCredential.User != "default-user" {
		t.Log("DefaultCredential.User not matched.")
		t.Logf("%#v", config.DefaultCredential)
		t.Fail()
	}
	if config.DefaultCredential.Pass != "default-pass" {
		t.Log("DefaultCredential.Pass not matched.")
		t.Fail()
	}
	if config.DefaultCredential.Identity != "default-identity" {
		t.Log("DefaultCredential.Identity not matched.")
		t.Fail()
	}

	hosts := map[string]map[string]string {
		"host-a": {"user": "host-a-user", "pass": "host-a-pass", "identity": "host-a-identity", },
		"host-b": {"user": "host-b-user", "pass": "host-b-pass", "identity": "host-b-identity", },
	}

	for hostName, host := range hosts {
		hostCredential := config.Hosts[hostName]

		if config.Hosts[hostName].User != host["user"] {
			t.Logf("host '%s': user does not matched.", hostCredential.User)
			t.Fail()
			return
		}

		if config.Hosts[hostName].Pass != host["pass"] {
			t.Logf("host '%s': pass does not matched.", hostCredential.Pass)
			t.Fail()
			return
		}
		if config.Hosts[hostName].Identity != host["identity"] {
			t.Logf("host '%s': identity does not matched.", hostCredential.Identity)
			t.Fail()
			return
		}
	}
}
