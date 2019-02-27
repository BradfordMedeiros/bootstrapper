package topics

import "testing"

func TestT1 (t *testing.T){
	if !IsValidTopic("temp"){
		t.Fatal()
	}
}
func TestT2 (t *testing.T){
	if IsValidTopic("#"){
		t.Fatal()
	}
}
func TestT3 (t *testing.T){
	if IsValidTopic("+"){
		t.Fatal()
	}
}
func TestT4 (t *testing.T){
	if !IsValidTopic("/room1/thing/go"){
		t.Fatal()
	}
}

func TestT5 (t *testing.T){
	if IsValidTopic("/room1/#/go"){	// actually not sure if I want this to be valid or not
		t.Fatal()
	}
}
