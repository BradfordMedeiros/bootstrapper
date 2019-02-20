package main

import "fmt"
import "os"
import "./parseOptions"
import "./commands/serve"
import "./commands/remoteServer"
import "./commands/download"

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
			if options.CommandUse.ServerUrl == nil {
				remoteServer.GetServer()
			}else{
				remoteServer.SetServer(*options.CommandUse.ServerUrl)
			}
		}
		case "download": {
			download.Download()
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

