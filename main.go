package main

import "fmt"
import "os"
import "./parseOptions"
import "./commands/serve"
import "./commands/download"
import "./config"
import "./commands/dataSetter"

func main(){
	options, err := parseOptions.ParseOptions(os.Args[1:])
	if err != nil {
		fmt.Println("error! ", err)
	}

	configuration, err := config.Read("./data")
	if err != nil {
		panic("Could not read config: " + err.Error())
	}

	switch (options.CommandType) {
		// Server commands
		case "serve": { 		
			serve.Start(
				configuration.Banner, 
				func (topic string, value string, tag string){
					fmt.Println("save topic: ")
					fmt.Println("topic: ", topic)
					fmt.Println("value: ", value)
				},
				func (topic string, tag string) string {
					fmt.Println("get topic")
					return "some topic"
				},
			)
		}

		// Client commands
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
		case "set": {
			resp, err := dataSetter.Set("some topic", "some data")
			if err != nil {
				fmt.Println("error setting: ", err.Error())
				return
			}
			fmt.Println(resp)
		}
		case "get": {
			resp, err := dataSetter.Get("some topic ")
			if err != nil {
				fmt.Println("error getting ", err.Error())
				return
			}
			fmt.Println(resp)
		}
		case "info": {
			infoResponse, err := dataSetter.Info()
			if err != nil {
				fmt.Println("error grabbing info: ", err.Error())
				return
			}
			fmt.Println(infoResponse)
		}
		case "banner": {
			bannerResponse, err := dataSetter.Banner()
			if err != nil {
				fmt.Println("error grabbing banenr: ", err.Error())
				return
			}
			fmt.Println(bannerResponse)
		}
		default : {
			panic("unknown command type")
		}
	}

}

