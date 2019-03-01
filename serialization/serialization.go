package serialization

import "io/ioutil"

func GetInMemorySerialization(_ string) (
	SaveTopic func(topic string, value string) error,
	GetTopics func() (map[string]string, error),
){
	topicMap := map[string]string { }
	saveTopic := func(topic string, value string) error {
		topicMap[topic] = value
		return nil
	}
	getTopics := func() (topics map[string]string, err error) {
		return topicMap, nil
	}
	return saveTopic, getTopics
}


type TopicPairs struct {
	topic string
	value string
}

func GetSerialization(filepath string) (
	SaveTopic func(topic string, value string) error,
	GetTopics func() (map[string]string, error),
){
	saveTopic := func(topic string, value string) error {
		return ioutil.WriteFile(filepath, []byte("topicplaceholderdata"), 0666)
	}
	getTopics := func() (topics map[string]string, err error) {
		return nil, nil
	}
	return saveTopic, getTopics
}
