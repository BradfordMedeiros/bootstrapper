package main

import "fmt"
import "os"
import "./parseOptions"
import "./commands/serve"
import "./commands/download"
import "./config"

func main(){
	fmt.Println("hello world")
	options, err := parseOptions.ParseOptions(os.Args[1:])
	if err != nil {
		fmt.Println("error! ", err)
	}

	configuration, err := config.Read("./data")
	if err != nil {
		panic("Could not read config: " + err.Error())
	}

	switch (options.CommandType) {
		case "serve": { 
			serve.Start()
		}
		case "use": {
			if options.CommandUse.ServerUrl == nil {
				fmt.Println(configuration.RemoteServer)
			}else{
				fmt.Println("trying to write new server url")
				configToWrite := config.Config{
					RemoteServer: *options.CommandUse.ServerUrl,
				}
				writeErr := config.Write(configToWrite)
				if writeErr != nil {
					panic ("Could not write config " + writeErr.Error())
				}
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

