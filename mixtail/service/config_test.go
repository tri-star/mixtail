package service_test


import (
	"testing"
	"github.com/tri-star/mixtail/lib"
	"github.com/tri-star/mixtail/mixtail/ext/extssh"
	"github.com/tri-star/mixtail/mixtail/service"
	"github.com/tri-star/mixtail/mixtail/ext"
	"github.com/tri-star/mixtail/mixtail/entity"
)

func TestParse(t *testing.T) {

	extensionManager := lib.NewExtensionManager()
	extensionManager.RegisterExtensionPoint(ext.POINT_INPUT_CONFIG_PARSER, extssh.NewInputEntryParser())

	cs := service.NewConfig(extensionManager)

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

	conf := entity.NewConfig()
	err := cs.Parse(yaml, conf)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if len(conf.InputEntries) != 2 {
		t.Logf("input handler count. expected: 1, actual: %d", len(conf.InputEntries))
		t.Fail()
	}

	inputSsh := conf.InputEntries[0].(*extssh.InputEntry)
	if inputSsh.Name != "test01" {
		t.Logf("input handler name not matched. expected: test01, actual: %s", inputSsh.Name)
		t.Fail()
	}
	if inputSsh.Host != "example.com" {
		t.Logf("host name not matched. expected: example.com, actual: %s", inputSsh.Host)
		t.Fail()
	}

	inputSsh2 := conf.InputEntries[1].(*extssh.InputEntry)
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
