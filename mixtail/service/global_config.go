package service
import (
	"github.com/tri-star/mixtail/mixtail/entity"
	"os"
	"errors"
)


const GLOBAL_CONFIG_FILE=".mixtail.yml"


type GlobalConfig struct {

}


func NewGlobalConfig() (gc *GlobalConfig) {
	gc = new(GlobalConfig)
	return
}


// GetFilePath searches global config file from specific order.
func (gc *GlobalConfig) GetFilePath() (filePath string, found bool) {

	//Looking for current directory.
	currentDirFile := "./" + GLOBAL_CONFIG_FILE
	if _, err := os.Stat(currentDirFile); err == nil {
		return currentDirFile, true
	}

	//Looking for user's home directory.
	os.Environ()
	homeDir := os.Getenv("HOME")
	if homeDir != "" {
		homeDirFile := homeDir + GLOBAL_CONFIG_FILE
		if _, err := os.Stat(homeDirFile); err == nil {
			return homeDirFile, true
		}
	}

	return "", false
}

func (gc *GlobalConfig) ParseFromFile(filePath string) (config *entity.Config, err error) {
	return
}

func (gc *GlobalConfig) Parse(data map[interface{}]interface{}) (config *entity.Config, err error) {

	config = entity.NewConfig()

	//Parse default section.
	defaultSection, ok := data["default"].(map[interface{}]interface{})
	if ok {
		config.DefaultCredential.User, ok = defaultSection["user"].(string)
		config.DefaultCredential.Pass, ok = defaultSection["pass"].(string)
		config.DefaultCredential.Identity, ok = defaultSection["identity"].(string)
	}

	//Parse host section.
	hostSection, ok := data["host"].(map[interface{}]interface{})
	if ok {
		var tmpCred *entity.Credential
		for hostName, _ := range hostSection {
			hostEntry, ok := hostSection[hostName.(string)].(map[interface{}]interface{})
			if !ok {
				err = errors.New("Invalid host section entry. host name: " + hostName.(string))
				return
			}
			tmpCred = entity.NewCredential()
			tmpCred.User = hostEntry["user"].(string)
			tmpCred.Pass = hostEntry["pass"].(string)
			tmpCred.Identity = hostEntry["identity"].(string)

			config.Hosts[hostName.(string)] = tmpCred
		}
	}
	return
}


func (gc *GlobalConfig) Merge(filePath string, config *entity.Config) (err error) {
	return
}
