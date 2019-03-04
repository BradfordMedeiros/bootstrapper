
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
import "path"

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

type TopicValuePair struct {
	Topic string
	Value string
}

func Start(
	relativeURL string,
	banner string, 
	saveTopic func (topic string, value string, tag string) error,
	getTopic func(topic string, tag string) ([]TopicValuePair, error),
	getInfo func() string,
	getBanner func() string,
) error {
	fmt.Println("bootstrapper server starting on port ", PORT)
	// ideally this could be done without side effects on http module, but not sure if the api call is available
	http.HandleFunc(path.Join(relativeURL, "/banner"), func(w http.ResponseWriter, r *http.Request) {  
		w.Write([]byte(getBanner()))
	})
	http.HandleFunc(path.Join(relativeURL, "/info"), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getInfo()))
	})
	http.HandleFunc(path.Join(relativeURL, "/get"), func(w http.ResponseWriter, r *http.Request) {
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

		topicData, errGetTopic := getTopic(getRequest.Topic, getRequest.Tag)
		if errGetTopic != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
			return
		}
		value, _ := json.Marshal(topicData)
		w.Write(value)
	})
	http.HandleFunc(path.Join(relativeURL, "/set"), func(w http.ResponseWriter, r *http.Request) {
		var setRequest SetRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&setRequest)
		if err != nil {
			w.Write([]byte("bad request"))
			return
		}
		if !isValidSetRequest(setRequest){
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed request parameters"))
			return
		}

		errSaveTopic := saveTopic(setRequest.Topic, setRequest.Data, setRequest.Tag)
		if errSaveTopic == nil {
			w.Write([]byte("ok"))
		}else{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad"))
		}
	})

	return http.ListenAndServe(":"+strconv.Itoa(int(PORT)), nil)
}