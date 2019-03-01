package topics

import "strings"
// This matches topics
/*

	topics:
	+ matches against 1, can be anywhere
	# matches against multiple, must be at end 
	must be at least 1 character in topic
	empty level okay

	See tests for examples
*/

func isMatchingTopic(topic string, matcher string) bool {
	topicPatterns := strings.Split(topic, "/")
	matcherPatterns := strings.Split(matcher, "/")

	for index, _ := range(topicPatterns) {
		if index > len(matcherPatterns) -1{
			break
		}

		switch(matcherPatterns[index]) {
			case "+":
			case "#":
			case topicPatterns[index]: {
				continue
			}
			default: {
				return false
			}
		}
	}

	if len(topicPatterns) == len(matcherPatterns){    			// everything matched and same length, the topic is a match
		return true
	}
	if len(matcherPatterns) > len(topicPatterns){    			// matcher cannot be bigger than the topic pattern
		return false 
	}

	if (matcherPatterns[len(matcherPatterns)-1] == "#"){		// everything matched up to here and matcher is smaller, so # covers rest of length
		return true
	}

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

func IsValidTopic(topic string) bool{
	if len(topic) == 0 {
		return false
	}

	topicParts := strings.Split(topic, "/")
	for _, topic := range(topicParts){
		if topic == "#" || topic == "+" {
			return false
		}
	}
	return true
}