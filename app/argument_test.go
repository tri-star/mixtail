package app_test

import (
	"flag"
	"testing"
)


func TestArgumentParse(t *testing.T) {

	arg := []string{"-a", "5", "-b", "6", "--c=10"}

	flagSet := flag.NewFlagSet("test", flag.ExitOnError)

	var argA int64
	var argB int64
	var argC int64

	flagSet.Int64Var(&argA, "a", 100, "flag a")
	flagSet.Int64Var(&argB, "b", 100, "flag b")
	flagSet.Int64Var(&argC, "c", 100, "flag c")

	err := flagSet.Parse(arg)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	t.Logf("a: %d", argA)
	t.Logf("b: %d", argB)
	t.Logf("c: %d", argC)
}
