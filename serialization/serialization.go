package serialization


func GetInMemorySerialization() (
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
	//return saveTopic
}


