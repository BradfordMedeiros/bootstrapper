package config

import "path"
import "io/ioutil"
import "strings"

type ClientConfig struct {
	RemoteServer string
	Servers []string
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
	filebytes, err := ioutil.ReadFile(filepath)
	fileContent := string(filebytes)
	if err != nil {
		return []string{}, err
	}
	servers := strings.Split(fileContent, "\n")
	return servers, nil
}
func writeServers(filepath string, servers []string) error {
	return ioutil.WriteFile(filepath, []byte(strings.Join(servers, "\n")), 0666)
}

type ServerConfig struct {
	Info string
	Banner string
}
func readBanner(filepath string) (string, error){
	filebytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(filebytes), nil
}
func readInfo(filepath string) (string, error){
	filebytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(filebytes), nil
}

func ReadClient(dataDirectory string) (ClientConfig, error){
	activeServer, _ := readActiveServer(path.Join(dataDirectory, "active_server"))
	servers, _ := readServers(path.Join(dataDirectory, "servers"))
	client := ClientConfig{
		RemoteServer: activeServer,
		Servers: servers,
	}
	return client, nil	
}
func ReadServer(dataDirectory string) (ServerConfig, error) {
	banner, _ := readBanner(path.Join(dataDirectory, "banner"))
	info, _ := readInfo(path.Join(dataDirectory, "info"))
	server := ServerConfig{
		Banner: banner,
		Info: info,
	}
	return server, nil
}
func WriteClient(dataDirectory string, config ClientConfig) error {
	writeActiveServer(path.Join(dataDirectory, "active_server"), config.RemoteServer)
	writeServers(path.Join(dataDirectory, "servers"), config.Servers)
	return nil
}