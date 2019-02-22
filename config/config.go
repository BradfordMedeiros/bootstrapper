package config

import "path"
import "io/ioutil"
//import "errors"

/*
	filebytes, _ := ioutil.ReadFile(filepath)
	writeErr := ioutil.WriteFile("./data", configuration, 0666)   // This file permission bit seems finicky

*/
type Config struct {
	RemoteServer string `json:"remote_server"`
	Banner string `json:"banner"`
}

func isValidConfig(config Config) bool {
	if config.RemoteServer == "" {
		return false
	}
	if config.Banner == "" {
		return false
	}
	return true
}

func readActiveServer(filepath string) (string, error){
	filebytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(filebytes), nil
}
func writeActiveServer(filepath string, activeServer string) error {
	return ioutil.WriteFile(filepath, []byte(activeServer), 0666) 
}
func readServers(filepath string) ([]string, error){
	return []string{}, nil
}
func writeServers(filepath string, servers []string) error {
	return nil
}
func readBanner(filepath string) (string, error){
	filebytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(filebytes), nil
}
func writeBanner(filepath string, banner string) error {
	return nil
}

func Read(dataDirectory string) (Config, error) {

	activeServer, _ := readActiveServer(path.Join(dataDirectory, "active_server"))
	banner, _ := readBanner(path.Join(dataDirectory, "banner"))

	return Config {
		RemoteServer: activeServer, 
		Banner: banner,
	}, nil
}

func Write(config Config) error {
	
	return nil
}