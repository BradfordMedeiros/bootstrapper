package httpClient

import "net/http"
import "io/ioutil"
import "bytes" 
import "encoding/json"
import "errors"

type HttpClient struct {
	Get func(topic string) (string, error)
	Set func(topic string, data string) (string, error)
	Banner func() (string, error)
	Info func() (string, error) 
}

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

func set(url string, topic string, data string) (string, error){
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
	return httpPost(url, bytes)
}

func get(url string, topic string) (string, error){
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

	return httpPost(url, bytes)
}

func banner(url string) (string, error) {
	return httpGet(url)
}

func info(url string) (string, error) {
	return httpGet(url)
}


func GetClient(url string) (HttpClient, error) {
	client := HttpClient { 
		Get: func(topic string) (string, error) {
			return get(url +"/get", topic)
		},
		Set: func(topic string, data string) (string, error){
			return set(url + "/set", topic, data)
		},
		Info: func() (string, error){
			return info(url + "/info")
		},
		Banner: func()(string, error){
			return banner(url + "/banner")
		},
	}
	return client, nil
}