package topics

// This matches topics
/*
	/something/is/cool
	/thing/is/right
	yes/how/is/right
	wow/hello/go
	wow

	This should be matched according to mqtt rules
*/

func isMatchingTopic(topic string, matcher string) bool {
	return false
}

func MatchTopics(topics []string, matcher string) []string{
	matchingTopics := []string{}

	for _, topic := range(topics){
		if isMatchingTopic(topic, matcher){
			matchingTopics = append(matchingTopics, topic)
		}
	}

	return matchingTopics
}