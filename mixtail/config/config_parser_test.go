package config_test


import (
	"testing"
	"github.com/tri-star/mixtail/mixtail/config"
	"github.com/tri-star/mixtail/lib"
	"github.com/tri-star/mixtail/mixtail/ext/input/extssh"
)

func TestParse(t *testing.T) {

	extensionManager := lib.NewExtensionManager()
	extensionManager.RegisterExtension(extssh.NewInputConfigParser())

	cp := config.NewConfigParser(extensionManager)

	yaml := []byte(`
input:
  test01:
    type: ssh
    host:
      - example.com
      - example02.com
    user: user_name
    identity: identity-file-name
    command: testtest
log:
  logging: true
  path: /tmp/test.log
`)

	err := cp.Parse(yaml)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	conf := cp.GetResult()

	if len(conf.Inputs) != 2 {
		t.Logf("input handler count. expected: 1, actual: %d", len(conf.Inputs))
		t.Fail()
	}

	inputSsh := conf.Inputs[0].(*extssh.InputConfig)
	if inputSsh.Name != "test01" {
		t.Logf("input handler name not matched. expected: test01, actual: %s", inputSsh.Name)
		t.Fail()
	}
	if inputSsh.Host != "example.com" {
		t.Logf("host name not matched. expected: example.com, actual: %s", inputSsh.Host)
		t.Fail()
	}

	inputSsh2 := conf.Inputs[1].(*extssh.InputConfig)
	if inputSsh2.Name != "test01" {
		t.Logf("input handler name not matched. expected: test01, actual: %s", inputSsh2.Name)
		t.Fail()
	}
	if inputSsh2.Host != "example02.com" {
		t.Logf("host name not matched. expected: example02.com, actual: %s", inputSsh2.Host)
		t.Fail()
	}

	if conf.Logging != true {
		t.Logf("Logging not enabled.")
		t.Fail()
	}
	if conf.LogPath != "/tmp/test.log" {
		t.Logf("LogPath not matched. actual: " + conf.LogPath)
		t.Fail()
	}
}
