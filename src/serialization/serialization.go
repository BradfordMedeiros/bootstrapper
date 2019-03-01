package serialization

import "io/ioutil"
import "encoding/json"

// map[string][string] is sort of horeshit since it disallows non-string based numbers stored in json but w/e for now
func ReadTopicFile(filepath string)(map[string]string, error) {
	filebytes, err := ioutil.ReadFile(filepath)
	topicContent := map[string]string {}
	if err != nil {
		return topicContent, err
	}
	deserializationErr := json.Unmarshal(filebytes, &topicContent)
	return topicContent, deserializationErr
}

func GetSerialization(filepath string) (
	SaveTopic func(topic string, value string) error,
	GetTopics func() (map[string]string, error),
){
	saveTopic := func(topic string, value string) error {
		topics, err := ReadTopicFile(filepath)
		if err != nil {
			return err
		}

		topics[topic] = value	 
		content, err := json.Marshal(topics)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(filepath, []byte(content), 0666)
	}
	getTopics := func() (topics map[string]string, err error) {
		return  ReadTopicFile(filepath)
	}
	return saveTopic, getTopics
}
