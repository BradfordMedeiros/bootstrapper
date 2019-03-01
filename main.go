package main

import "fmt"
import "os"
import "errors"
import "./parseOptions"
import "./serialization"
import "./commands/serve"
import "./config"
import "./commands/httpClient"
import "./commands/topics"


const dataDirectory = "./data"

func main(){
	options, err := parseOptions.ParseOptions(os.Args[1:])
	if err != nil {
		fmt.Println("error! ", err)
	}
	configuration, err := config.Read(dataDirectory)

	getSerialization := serialization.GetSerialization
	saveTopic, getTopics := getSerialization("./data/topics")
	
	if err != nil {
		panic("Could not read config: " + err.Error())
	}

	switch (options.CommandType) {
		// Server commands
		case "serve": { 		
			err := serve.Start(
				configuration.Banner, 
				func (topic string, value string, tag string) error {
					if !topics.IsValidTopic(topic){
						return errors.New("invalid topic " + topic)
					}
					saveTopic(topic, value)
					return nil
				},
				func (topic string, tag string) ([]serve.TopicValuePair, error) {
					topicMap, _ := getTopics()
					topicArray := []string{}

					for topicKey, _ := range topicMap {
						topicArray = append(topicArray, topicKey)
					}

					matchingTopics := topics.MatchTopics(topicArray, topic)
					matchingValues := []string{}

					for _, topic := range(matchingTopics){
						topicValue, _ := topicMap[topic]
						matchingValues = append(matchingValues, topicValue)
					}

					topicValuePairs := []serve.TopicValuePair{}
					for index, _ := range(matchingTopics){
						topicValuePairs = append(topicValuePairs, 
							serve.TopicValuePair{ 
								Topic: matchingTopics[index],
								Value: matchingValues[index],
							},
						)
					}
					return topicValuePairs, nil
				},
				func () string {
					return configuration.Info
				},
				func () string {
					return configuration.Banner
				},
			)
			if err != nil {
				panic("error starting server")
			}
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
				writeErr := config.Write(dataDirectory, configToWrite)
				if writeErr != nil {
					panic ("Could not write config " + writeErr.Error())
				}
			}
		}
		case "set": {
			resp, err := httpClient.Set(options.CommandSet.Key, options.CommandSet.Value)
			if err != nil {
				os.Exit(2)
				return
			}
			fmt.Println(resp)
		}
		case "get": {
			resp, err := httpClient.Get(options.CommandGet.Key)
			if err != nil {
				os.Exit(2)
				return
			}
			fmt.Println(resp)
		}
		case "info": {
			infoResponse, err := httpClient.Info()
			if err != nil {
				fmt.Println("error grabbing info: ", err.Error())
				return
			}
			fmt.Println(infoResponse)
		}
		case "banner": {
			bannerResponse, err := httpClient.Banner()
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

