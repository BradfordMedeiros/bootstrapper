package config

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

func Read(filepath string) (Config, error) {
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

func Write(config Config) error {
	configuration, err := json.Marshal(config)
	if err != nil {
		return err
	}
	writeErr := ioutil.WriteFile("./data", configuration, 0666)   // This file permission bit seems finicky
	if writeErr != nil {
		return writeErr
	}
	return nil
}