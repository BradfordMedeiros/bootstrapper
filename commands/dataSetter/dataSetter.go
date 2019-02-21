package dataSetter

import "net/http"
import "io/ioutil"
import "bytes" 
import "encoding/json"
import "fmt"

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
	fmt.Println("sending data: ", string(jsonBytes))
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

func Set(topic string, data string) (string, error){
	value := struct {
		Topic string `json:"topic"`
		Data string `json:"data"`
		Tag string `json:"tag"`
	}{
		Topic: topic,
		Data: data,
		Tag: "tag placeholder",
	}	
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return httpPost("http://localhost:8000/set", bytes)
}

func Get(topic string) (string, error){
	value := struct {
		Topic string `json:"topic"`
		Tag string `json:"tag"`
	}{
		Topic: topic, 
		Tag: "some tag",
	}
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return httpPost("http://localhost:8000/get", bytes)
}

func Banner() (string, error) {
	return httpGet("http://localhost:8000/banner")
}

func Info() (string, error) {
	return httpGet("http://localhost:8000/info")
}