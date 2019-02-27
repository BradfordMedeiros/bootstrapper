package topics

import "testing"
import "sort"

func isMatchingCorrectTopics(topics []string, matchExpression string, expectedTopics []string) (bool, []string){
	matchingTopics := MatchTopics(topics, matchExpression)
	sort.Strings(matchingTopics)
	
	expectedTopicsCopy := make([]string, len(expectedTopics))
	copy(expectedTopicsCopy, expectedTopics)
	sort.Strings(expectedTopicsCopy)
	
	if len(matchingTopics) != len(expectedTopicsCopy){
		return false, matchingTopics
	}

	for i, _ := range(matchingTopics){
		if matchingTopics[i] != expectedTopicsCopy[i] {
			return false, matchingTopics
		}
	}
	return true, matchingTopics
}

func Test1(t *testing.T){
	isMatching, _ := isMatchingCorrectTopics([]string{ "temp", "humidity", "rain"}, "rain", []string { "rain" })
	if !isMatching {
		t.Fatal()
	}
}
func Test2(t *testing.T){
	isMatching, _ := isMatchingCorrectTopics([]string{ "temp", "humidity", "rain"}, "bagels", []string { })
	if !isMatching {
		t.Fatal()
	}
}

func Test3(t *testing.T){		
	isMatching, matchingTopics := isMatchingCorrectTopics([]string{ "temp", "humidity", "rain" }, "+", []string{"temp", "humidity", "rain"})
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}
func Test4(t *testing.T){
	isMatching, matchingTopics := isMatchingCorrectTopics([]string{ "temp", "humidity", "rain" }, "wow+", []string{})
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}

func Test5(t *testing.T){	   // temperature, rain, thing  t# 
	isMatching, matchingTopics := isMatchingCorrectTopics([]string { "temp", "humidity", "rain" }, "#", []string{ "temp", "humidity", "rain"})
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}

func Test6(t *testing.T){
	isMatching, matchingTopics := isMatchingCorrectTopics([]string { "room1/temp", "room2/temp", "room2/humidity" }, "+/temp", []string{ "room1/temp", "room2/temp" })
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}
func Test7(t *testing.T){
	isMatching, matchingTopics := isMatchingCorrectTopics([]string { "room1/temp", "room2/temp", "room2/humidity" }, "+/+", []string{ "room1/temp", "room2/temp", "room2/humidity" })
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}

func Test8(t *testing.T){
	isMatching, matchingTopics := isMatchingCorrectTopics([]string { "room1/temp", "room2/temp", "room2/humidity" }, "+/#", []string{ "room1/temp", "room2/temp", "room2/humidity" })
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}

func Test9(t *testing.T){
	isMatching, matchingTopics := isMatchingCorrectTopics([]string { "/temp" }, "temp", []string{ })
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}


func Test10(t *testing.T){
	isMatching, matchingTopics := isMatchingCorrectTopics([]string { "/temp" }, "/#", []string{"/temp" })
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}

func Test11(t *testing.T){
	isMatching, matchingTopics := isMatchingCorrectTopics([]string { "/temp" }, "#/", []string{ })
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}
func Test12(t *testing.T){
	isMatching, matchingTopics := isMatchingCorrectTopics([]string { "temp" }, "temp/", []string{ })
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}
func Test13(t *testing.T){
	isMatching, matchingTopics := isMatchingCorrectTopics([]string { "temp/1/2" }, "temp/1/2/3", []string{ })
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}
func Test14(t *testing.T){
	isMatching, matchingTopics := isMatchingCorrectTopics([]string { "temp/1/2/" }, "temp/1/2/#", []string{ "temp/1/2/"})
	if !isMatching {
		t.Errorf("got %v", matchingTopics)
	}
}