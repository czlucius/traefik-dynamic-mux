package muxfile

import (
	"fmt"
	"strings"

	"czlucius.dev/tmap/dynamic"
	"czlucius.dev/tmap/muxfile/commands"
)

/*
ParseMuxFile parses a mux file and returns a map of routes.
The following is a syntax of a Muxfile:
```
COMMAND arg1 arg2 arg3
```

For example,
```
JSONURL http://localhost:3000/webhooks/traefik/main.json
YAMLFILE /etc/traefik/config.yml
RELABEL_ENTRYPOINT web http
RELABEL_ENTRYPOINT websecure https
APPLYCERTRESOLVER letsencrypt https
DEFAULTCERT path/to/cert.crt path/to/cert.key
```
*/

type Command struct {
	Name string
	Args []string
}

type CommandImpl interface {
	Execute(args []string, config *dynamic.Configuration) error
}

// Factory function to create commands
func CreateCommand(name string) (CommandImpl, error) {
	switch name {
	case "JSONURL":
		return &commands.JSONURLCommand{}, nil
	case "YAMLFILE":
		return &commands.YAMLFileCommand{}, nil
	case "APPLYCERTRESOLVER":
		return &commands.CertResolverCommand{}, nil
	case "RELABEL_ENTRYPOINT":
		return &commands.RelabelEntryPointCommand{}, nil
	case "DEFAULTCERT":
		return &commands.DefaultCertCommand{}, nil

	default:
		return nil, fmt.Errorf("unknown command: %s", name)
	}
}

func ParseCommand(command string) (*Command, error) {
	// Split the command into parts
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty command")
	}

	// Create a new Command struct
	cmd := &Command{
		Name: parts[0],
		Args: parts[1:],
	}

	return cmd, nil
}

func ExecMuxFile(muxFileContents string, config *dynamic.Configuration) error {

	// Split the contents into lines
	commands := strings.Split(muxFileContents, "\n")

	// Process each command
	for _, line := range commands {
		// Skip empty lines or lines with only whitespace
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if line[0] == '#' {
			// Skip comments
			continue
		}
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		fmt.Println("Executing command:", line)

		// Create the command
		cmdImpl, err := CreateCommand(parts[0])
		if err != nil {
			return err
		}

		// Execute the command
		err = cmdImpl.Execute(parts[1:], config)
		if err != nil {
			return err
		}
	}

	return nil

}
