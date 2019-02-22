
// This serves provides an http server implementation for a valid bootstrapper server endpoint
// Servers require the following routes:
/*

/get/path/to/the/topic?tag=sometaghere
-X = GET

/set/path/to/the/topic?tag=sometaghere 
-X = POST
post-body: data to post 

*/

package serve 

import "fmt"
import "net/http"
import "strconv"
import "encoding/json"

const PORT = 8001

type SetRequest struct {
	Topic string `json:"topic"`
	Tag string `json:"tag"`
	Data string `json:"data"`
}
func isValidSetRequest(setReq SetRequest) bool{
	if setReq.Topic == "" {
		return false
	}
	if setReq.Data == "" {
		return false
	}
	return true
}

type GetRequest struct {
	Topic string `json:"topic"`
	Tag string `json:"tag"`
}
func isValidGetRequest(getReq GetRequest) bool {
	if getReq.Topic == "" {
		return false
	}
	return true
}

func Start(
	banner string, 
	saveTopic func (topic string, value string, tag string),
	getTopic func(topic string, tag string) string,
	getInfo func() string,
	getBanner func() string,
){
	fmt.Println("bootstrapper server starting")
	// ideally this could be done without side effects on http module, but not sure if the api call is available
	http.HandleFunc("/banner", func(w http.ResponseWriter, r *http.Request) {  
		w.Write([]byte(getBanner()))
	})
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getInfo()))
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		var getRequest GetRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&getRequest)
		if err != nil {
			w.Write([]byte("bad request"))
			return
		}
		if !isValidGetRequest(getRequest){
			w.Write([]byte("bad validation"))
			return
		}

		topicData := getTopic(getRequest.Topic, getRequest.Tag)
		w.Write([]byte(topicData))
	})
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		var setRequest SetRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&setRequest)
		if err != nil {
			w.Write([]byte("bad request"))
			return
		}
		if !isValidSetRequest(setRequest){
			w.Write([]byte("bad validation"))
			return
		}

		saveTopic(setRequest.Topic, setRequest.Data, setRequest.Tag)
		w.Write([]byte("ok"))
	})

	err := http.ListenAndServe(":"+strconv.Itoa(int(PORT)), nil)
	fmt.Println("err ", err.Error())
}