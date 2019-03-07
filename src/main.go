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
	serverConfig, err := config.ReadServer(dataDirectory)
	clientConfig, err := config.ReadClient(dataDirectory)
	
	if err != nil {
		panic("Could not read config: " + err.Error())
	}

	client, _ := httpClient.GetClient(clientConfig.RemoteServer)
	switch (options.CommandType) {
		// Server commands
		case "serve": { 
			topicFile := options.CommandServe.TopicFile
			fmt.Println("topic data file is: ", topicFile)
			fmt.Println("relative url: ", options.CommandServe.RelativeTo)
			saveTopic, getTopics := serialization.GetSerialization(topicFile)		
			
			err := serve.Start(
				options.CommandServe.RelativeTo,
				serverConfig.Banner, 
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
					return serverConfig.Info
				},
				func () string {
					return serverConfig.Banner
				},
			)
			if err != nil {
				panic("error starting server")
			}
		}

		// Client commands
		case "use": {		   
			if options.CommandUse.ServerUrl == nil {
				fmt.Println(clientConfig.RemoteServer)
			}else{
				writeErr := config.WriteClient(dataDirectory, config.ClientConfig{
					RemoteServer: *options.CommandUse.ServerUrl,
				})
				if writeErr != nil {
					panic ("Could not write config " + writeErr.Error())
				}
			}
		}
		case "set": {
			resp, err := client.Set(options.CommandSet.Key, options.CommandSet.Value)
			if err != nil {
				os.Exit(2)
				return
			}
			fmt.Println(resp)
		}
		case "get": {
			resp, err := client.Get(options.CommandGet.Key)
			if err != nil {
				os.Exit(2)
				return
			}
			fmt.Println(resp)
		}
		case "info": {
			infoResponse, err := client.Info()
			if err != nil {
				fmt.Println("error grabbing info: ", err.Error())
				return
			}
			fmt.Println(infoResponse)
		}
		case "banner": {
			bannerResponse, err := client.Banner()
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

