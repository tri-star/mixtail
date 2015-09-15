package entity_test
import (
	"testing"
	"github.com/tri-star/mixtail/mixtail/entity"
)

func TestGetDefaultCrediential(t *testing.T) {

	config := entity.NewConfig()

	defaultCred := entity.NewCredential()
	defaultCred.User = "default-user"

	hostACred := entity.NewCredential()
	hostACred.User = "host-a-user"

	hostBCred := entity.NewCredential()
	hostBCred.User = "host-b-user"

	config.DefaultCredential = defaultCred
	config.Hosts["host-a"] = hostACred
	config.Hosts["host-b"] = hostBCred

	result := config.GetDefaultCredential("host-a")
	if result.User != "host-a-user" {
		t.Log("host-a user not matched.")
		t.Fail()
	}

	result = config.GetDefaultCredential("host-b")
	if result.User != "host-b-user" {
		t.Log("host-b user not matched.")
		t.Fail()
	}

	result = config.GetDefaultCredential("host-c")
	if result.User != "default-user" {
		t.Log("host-c user not matched.")
		t.Fail()
	}

	result = config.GetDefaultCredential("")
	if result.User != "default-user" {
		t.Log("empty host user not matched.")
		t.Fail()
	}
}