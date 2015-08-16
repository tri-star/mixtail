package config_test


import (
	"testing"
	"github.com/tri-star/mixtail/config"
)

func TestParse(t *testing.T) {

	cp := config.NewConfigParser()

	yaml := []byte(`
input:
  test01:
    type: ssh
    host: example.com
    user: user_name
    identity: identity-file-name
    command: testtest
`)

	err := cp.Parse(yaml)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	conf := cp.GetResult()
	if len(conf.Inputs) != 1 {
		t.Logf("input handler count. expected: 1, actual: %d", len(conf.Inputs))
		t.Fail()
	}

	inputRemote, ok := conf.Inputs[0].(*config.InputRemote)
	if !ok {
		t.Logf("input handler type. expected: InputRemote, actual: %v", conf.Inputs)
		t.Fail()
	}

	if inputRemote.Name != "test01" {
		t.Logf("input handler name. expected: test01, actual: %s", inputRemote.Name)
		t.Fail()
	}
	if inputRemote.Host != "example.com" {
		t.Logf("host name. expected: example.com, actual: %s", inputRemote.Host)
		t.Fail()
	}

	t.Logf("%+v", inputRemote)

}
