package parseOptions

import "fmt"
import "flag"
import "os"
import "errors"

type Options struct {
	CommandType string
	CommandGet *GetCommand
	CommandSet *SetCommand
	CommandUse *UseCommand
	CommandServe *ServeCommand
	CommandBanner *BannerCommand
	CommandInfo *InfoCommand
}
type GetCommand struct {
	Key string
	Tag *string
}
func parseGetCommand (args []string) (GetCommand, error) {
	var tag *string = nil
	if len(args) >= 2 {
		tag = &args[1]
	} 
	return GetCommand{ Key: args[0], Tag: tag }, nil
}

type SetCommand struct {
	Key string
	Value string
	Tag string
}
func parseSetCommand (args []string) (SetCommand, error) {
	return SetCommand{ Key: args[0], Value: args[1] }, nil
}

type UseCommand struct {
	ServerUrl *string
}
func parseUseCommand(args []string) (UseCommand, error){
	var serverUrl *string = nil
	if len(args) >= 1 {		
		serverUrl = &args[0]
	}
	return UseCommand{ ServerUrl: serverUrl }, nil
}

type ServeCommand struct {
	TopicFile string
	RelativeTo string
}
func parseServeCommand(args []string) (ServeCommand, error){
	topicFile := "./data/topics"
	if len(args) >= 1 {
		topicFile = args[0]
	}

	relativeTo := "/"
	if len(args) >= 2 {
		relativeTo = args[1]
	}
	return ServeCommand{ TopicFile: topicFile, RelativeTo: relativeTo }, nil
}
type BannerCommand struct {}
type InfoCommand struct {}

func ParseArgs(args[] string) (Options, error) {
	fs := flag.NewFlagSet("main", flag.ExitOnError)
	printHelp := fs.Bool("h", false, "print help for top level any command")
	fs.Parse(args)

	if len(args) <= 0 {
		return Options{}, errors.New("invalid")
	}
	commandSurface := args[0]
	switch(commandSurface){
		case "get": {
			getCommand, err := parseGetCommand(args[1:])
			return Options{ CommandType: "get", CommandGet: &getCommand }, err
		}
		case "set": {
			setCommand, err := parseSetCommand(args[1:])
			return Options{ CommandType: "set", CommandSet: &setCommand }, err
		}
		case "server": {
			useCommand, err := parseUseCommand(args[1:])
			return Options{ CommandType: "use", CommandUse: &useCommand}, err
		}
		case "serve" : {
			serveCommand, err := parseServeCommand(args[1:])
			return Options{ CommandType: "serve", CommandServe: &serveCommand }, err
		}
		case "banner": {
			return Options{ CommandType: "banner", CommandBanner: &BannerCommand{} }, nil
		}
		case "info" : {
			return Options { CommandType: "info", CommandInfo: &InfoCommand{} }, nil
		}
		default : {
			if *printHelp || true {
		   		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
				fs.PrintDefaults()
				fmt.Fprintf(os.Stderr, "\nClient Commands\n---------------------\n")
				fmt.Fprintf(os.Stderr, "info - get info from active server\n")
				fmt.Fprintf(os.Stderr, "banner - get banner from active server\n")
				fmt.Fprintf(os.Stderr, "get - gets topic value from the active server\n")
				fmt.Fprintf(os.Stderr, "set - sets topic on the active server\n")
				fmt.Fprintf(os.Stderr, "server - manipulate remote server address for client\n")

				fmt.Fprintf(os.Stderr, "\nServer Commands\n---------------------\n")
				fmt.Fprintf(os.Stderr, "serve - host a local server backend instance\n\n")
			}
		}
	}
	return Options{ CommandType: "exit" }, nil
}