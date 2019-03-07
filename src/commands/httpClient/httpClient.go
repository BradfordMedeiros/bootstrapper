package httpClient

import "net/http"
import "io/ioutil"
import "bytes" 
import "encoding/json"
import "errors"

// @todo probably should do status codes correctly
func httpGet(route string) (string, error){
	resp, err := http.Get(route)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return "", errors.New("bad request")
	}

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

	if resp.StatusCode != 200 {
		return "", errors.New("bad request")
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", readErr
	}
	return string(body), err
}

func Set(url string, topic string, data string) (string, error){
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
	return httpPost(url + "set", bytes)
}

func Get(url string, topic string) (string, error){
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

	return httpPost(url +"/get", bytes)
}

func Banner(url string) (string, error) {
	return httpGet(url + "/banner")
}

func Info(url string) (string, error) {
	return httpGet(url + "/info")
}