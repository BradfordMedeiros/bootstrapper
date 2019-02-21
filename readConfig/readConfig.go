package readConfig

import "io/ioutil"
import "encoding/json"
import "errors"

type Config struct {
	RemoteServer string `json:"remote_server"`
}

func isValidConfig(config Config) bool {
	if config.RemoteServer == "" {
		return false
	}
	return true
}

func ReadConfig(filepath string) (Config, error) {
	filebytes, _ := ioutil.ReadFile(filepath)

	var config Config = Config{}
	err := json.Unmarshal(filebytes, &config)

	if !isValidConfig(config) {
		return Config{}, errors.New("invalid configuration")
	}

	if err != nil {
		return Config{}, err
	}
	return config, nil
}