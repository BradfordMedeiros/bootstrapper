package topics

import "testing"

func TestBasicTopic(t *testing.T){
	matchingTopics := MatchTopics([]string { "temperature", "humidity", "rain" }, "rain")
	if len(matchingTopics) != 1 {
		t.Errorf("incorrect number of matching topics")
		return
	}
	if matchingTopics[0] != "rain" {
		t.Errorf("incorrect matching topic, got " + matchingTopics[0])
	}
}

func TestSingleTopicLevelWildcardOneTopic(t *testing.T){		// should be temperature, * 
	t.Errorf("not implemented")
}

func TestSingleTopicLevelWildcardMultiTopics(t *testing.T){	   // temperature, rain, thing  t# 
	t.Errorf("not implemented")
}

func TestSingleMultiLevelWildcard(t *testing.T){
	t.Errorf("multi level fail")
}

func TestMultiMultiLevelWildcard(t *testing.T){
	t.Error("multi level multi topic fail")
}