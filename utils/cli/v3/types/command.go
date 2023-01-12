package types

import (
	"fmt"
	"strings"
)

type JsonScriptCommand struct {
	Name          string   `json:"Name"`
	Aliases       []string `json:"Aliases"`
	Description   string   `json:"Description"`
	Flags         []string `json:"Flags"`
	Help          string   `json:"Help"`
	ScriptAddress string   `json:"ScriptAddresss"`
}

type ScriptsList []JsonScriptCommand

type Command struct {
	name        string
	aliases     []string
	description string
	flags       []string
	help        string
}

// Name returns the given name of the command
func (C *Command) Name() string { return C.name }

// Aliases returns all possible alternative names
func (C *Command) Aliases() []string { return C.aliases }

// Description returns a brief summary of the command's effect
func (C *Command) Description() string {
	lines := strings.Join([]string{
		C.name,
		fmt.Sprintf("Aliases: [%s]", strings.Join(C.aliases, ", ")),
		fmt.Sprintf("Flags: [%s]", strings.Join(C.flags, ", ")),
		C.description,
		C.help,
	}, "\n")
	return lines
}

// Match checks if the input string matches the commands aliases
func (C *Command) Match(input string) bool {

	if C.name == input {
		return true
	}
	for _, it := range C.aliases {
		if it == input {
			return true
		}
	}
	return false

}
func (C *Command) Flags() []string
func (C *Command) ReadFlags([]string) map[string][]string
func (C *Command) Execute([]string) error
func (C *Command) Help() string
