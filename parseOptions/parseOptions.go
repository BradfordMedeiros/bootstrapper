package parseOptions

import "errors"
import "fmt"

type Options struct {
	CommandType string
	CommandGet *GetCommand
	CommandSet *SetCommand
	CommandDownload *DownloadCommand
	CommandUse *UseCommand
	CommandServe *ServeCommand
}
type GetCommand struct {
	Key string
	Tag *string
}
type SetCommand struct {
	Key string
	Value string
	Tag string
}
type DownloadCommand struct {}
type UseCommand struct {
	ServerUrl *string
}
type ServeCommand struct {}

func parseGetCommand (args []string) (GetCommand, error) {
	var tag *string = nil
	if len(args) >= 2 {
		tag = &args[1]
	} 
	return GetCommand{ Key: args[0], Tag: tag }, nil
}
func parseSetCommand (args []string) (SetCommand, error) {
	return SetCommand{ Key: args[0], Value: args[1] }, nil
}
func parseDownloadCommand(args []string) (DownloadCommand, error){
	return DownloadCommand{}, nil
}
func parseUseCommand(args []string) (UseCommand, error){
	var serverUrl *string = nil
	if len(args) >= 2{		
		serverUrl = &args[1]
	}
	return UseCommand{ ServerUrl: serverUrl }, nil
}
func parseServeCommand(args []string) (ServeCommand, error){
	return ServeCommand{}, nil
}

func ParseOptions(args []string) (Options, error) {
	if len(args) < 1 {
		return Options{}, errors.New("no args")
	}

	command := args[0]
	commandArgs := args[1:]
	switch(command) {
		case "get" : {
			getCommand, err := parseGetCommand(commandArgs)
			return Options{ CommandType: "get", CommandGet: &getCommand }, err
		}
		case "set": {
			setCommand, err := parseSetCommand(commandArgs)
			return Options{ CommandType: "set", CommandSet: &setCommand }, err
		}
		case "serve": {
			serveCommand, err := parseServeCommand(commandArgs)
			return Options{ CommandType: "serve", CommandServe: &serveCommand}, err
		}
		case "use": {
			useCommand, err := parseUseCommand(commandArgs)
			return Options{ CommandType: "use", CommandUse: &useCommand}, err
		}
	}
	return Options{}, errors.New("invalid type")
}

func PrintOptions(opts Options){
	switch(opts.CommandType) {
		case "get": {
			fmt.Println(opts.CommandGet.Key, " ", opts.CommandGet.Tag)
			break;
		}
		case "set": {
			fmt.Println(*opts.CommandSet)
			break;
		}
		case "serve" : {
			fmt.Println(*opts.CommandServe)
			break;
		}
		default : { 
			fmt.Println("cannot print ", opts.CommandType)

		}
	}
}