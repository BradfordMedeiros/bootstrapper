package main

import "fmt"
import "os"
import "./parseOptions"
import "./commands/serve"

func main(){
	fmt.Println("hello world")
	options, err := parseOptions.ParseOptions(os.Args[1:])
	if err != nil {
		fmt.Println("error! ", err)
	}


	switch (options.CommandType) {
		case "serve": { 
			serve.Start()
		}
		case "use": {
			fmt.Println("use not yet implemented")
		}
		case "download": {
			fmt.Println("download not yet implemented")
		}
		case "get": {
			fmt.Println("get placeholder")
		}
		case "set": {
			fmt.Println("set placeholder")
		}
		default : {
			fmt.Println("unknown command type")
		}
	}

}

