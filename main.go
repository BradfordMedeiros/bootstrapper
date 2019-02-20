package main

import "fmt"
import "os"
import "./parseOptions"
import "./commands/serve"
import "./commands/remoteServer"

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
				serverUrl := "some test url"
				remoteServer.SetServer(serverUrl)
			}
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

