package dataSetter

import "net/http"
import "io/ioutil"
import "bytes" 

// @todo probably should do status codes correctly
func httpGet(route string) (string, error){
	resp, err := http.Get(route)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(body), nil
}
func httpPost(route string, jsonBytes []byte) (string, error) {
	jsonContent := bytes.NewReader(jsonBytes)
	resp, err := http.Post(route, "application/json", jsonContent)
	
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", readErr
	}
	return string(body), err
}

func Set() (string, error){
	jsonBytes := []byte("helloworld")
	return httpPost("http://localhost:8000/set", jsonBytes)
}

func Get() (string, error){
	jsonBytes := []byte("helloworld")
	return httpPost("http://localhost:8000/get", jsonBytes)
}

func Banner() (string, error) {
	return httpGet("http://localhost:8000/banner")
}

func Info() (string, error) {
	return httpGet("http://localhost:8000/info")
}