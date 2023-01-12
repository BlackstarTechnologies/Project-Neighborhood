package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ScriptFetch struct holds the unmarshalled data from a Command.json to be turned into executable commands
type ScriptFetch struct {
	Name        string            `json:"Name"`
	Aliases     []string          `json:"Aliases"`
	Description string            `json:"Description"`
	Script      string            `json:"Script"`
	Flags       []string          `json:"Flags"`
	Subcommands map[string]string `json:"Subcommands"`
	Help        string            `json:"Help"`
}

// Subcommands returns a string version of the subcommands
func (SC *ScriptFetch) SubcommandsToString() string {
	// Marshal the map into a JSON string.
	userData, err := json.Marshal(SC.Subcommands)
	if err != nil {

		return ""
	}
	data1 := string(userData)
	data2 := strings.Split(data1, ",")
	data3 := strings.Join(data2, "\n")
	return data3
}

// ScriptJSON holds the list of scripts and commands
type ScriptJSON struct {
	fileAddress string
	scriptsList []ScriptFetch
}

// ScriptCommand is the wrapper struct for the fetched scripts
type ScriptCommand struct {
	Script ScriptFetch
}

// Name returns the command name
func (SC *ScriptCommand) Name() string {
	return SC.Script.Name
}

// Aliases returns the aliases for the command
func (SC *ScriptCommand) Aliases() []string {
	c := []string{SC.Name()}
	c = append(c, SC.Script.Aliases...)
	return c
}

// Description returns the description of a command (to display in help)
func (SC *ScriptCommand) Description() string {
	return SC.Script.Description
}

// Match identifies whether a user input matches 'this' script
func (SC *ScriptCommand) Match(com string) bool {
	if com == SC.Name() {
		return true
	}
	if len(SC.Script.Aliases) > 0 {
		for _, alias := range SC.Aliases() {
			if com == alias {
				return true
			}
		}

	}
	return false
}

// Flags returns the commandline options a user can use
func (SC *ScriptCommand) Flags() []string {
	return SC.Script.Flags
}

// ReadFlags returns the set of flags input by user, and the specific input to that flag
func (SC *ScriptCommand) ReadFlags([]string) map[string][]string {
	var res map[string][]string

	return res

}

// Execute executes the given script using the arguments given
func (SC *ScriptCommand) Execute(args []string) error {
	if !SC.Match(args[0]) {
		return errors.New("incompatible command")
	}

	return nil
}

// MapToString returns a stringified verions of a map
func MapToString(a_map map[interface{}]interface{}) (string, error) {
	// Marshal the map into a JSON string.
	userData, err := json.Marshal(a_map)
	if err != nil {

		return "", err
	}

	jsonStr := string(userData)
	return jsonStr, nil
}

// Help gives the name, aliases, description and examples of the command
func (SC *ScriptCommand) Help() string {

	res := fmt.Sprintf("Name: %s\nDescription: %s\nAliases: %s\nScript: %s\nFlags: %s\nSubcommands: %s\nHelp:\n%s\n",
		SC.Script.Name,                        // Name
		SC.Script.Description,                 // Description
		strings.Join(SC.Script.Aliases, ", "), // Aliases
		SC.Script.Script,                      // Script args
		SC.Script.Flags,                       // Input flags
		SC.Script.SubcommandsToString(),       // Script subcommands/variants
		SC.Script.Help,
	)

	return res
}
