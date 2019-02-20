package parseOptions

import "errors"

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
	Tag string
}
type SetCommand struct {
	Key string
	Value string
	Tag string
}
type DownloadCommand struct {}
type UseCommand struct {
	serverUrl string
}
type ServeCommand struct {}

func parseGetCommand (args []string) (GetCommand, error) {
	return GetCommand{ Key: "default key" }, nil
}
func parseSetCommand (args []string) (SetCommand, error) {
	return SetCommand{}, nil
}
func parseDownloadCommand(args []string) (DownloadCommand, error){
	return DownloadCommand{}, nil
}
func parseUseCommand(args []string) (UseCommand, error){
	return UseCommand{}, nil
}
func parseServeCommand(args []string) (ServeCommand, error){
	return ServeCommand{}, nil
}

func ParseOptions(args []string) (Options, error) {
	if len(args) < 1 {
		return Options{}, errors.New("no args")
	}

	firstArg := args[0]
	switch(firstArg) {
		case "get" : {
			getCommand, err := parseGetCommand(args[1:])
			return Options{ CommandType: "get", CommandGet: &getCommand }, err
		}
		case "set": {
			setCommand, err := parseSetCommand(args[1:])
			return Options{ CommandType: "set", CommandSet: &setCommand }, err
		}
	}
	return Options{}, errors.New("invalid type")
}

