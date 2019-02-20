package main

import "fmt"
import "os"
import "./parseOptions"

func main(){
	fmt.Println("hello world")
	options, err := parseOptions.ParseOptions(os.Args[1:])
	if err != nil {
		fmt.Println("error! ", err)
	}

	fmt.Println(options.CommandGet)
}