package types

import (
	"fmt"
	"regexp"
	"strings"
)

// JSONScriptCommand holds unmarshalled data from script Command.json list
type JSONScriptCommand struct {
	Name          string   `json:"Name"`
	Aliases       []string `json:"Aliases"`
	Regex         string   `json:"Regex"`
	Description   string   `json:"Description"`
	Flags         []string `json:"Flags"`
	Help          string   `json:"Help"`
	ScriptAddress string   `json:"ScriptAddresss"`
}

// ScriptsList is the Javascript Object stored in Command.json files
type ScriptsList []JSONScriptCommand

// Command script is the one
type Command struct {
	name        string
	aliases     []string
	regex       string
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

// Flags returns possible flags to adjust command execution
func (C *Command) Flags() []string {
	return C.flags
}

func (C *Command) isInFlags(a string) bool {
	for _, it := range C.flags {
		if it == a {
			return true
		}
	}
	return false
}
func (C *Command) isFlag(input string) bool {
	if len(input) < 1 {
		return false
	}
	if input[0] != '-' {
		return false
	}

	b, _ := regexp.MatchString(C.regex, input)
	if !b {
		return false
	}
	return true
}

// ReadFlags returns the input per tag
func (C *Command) ReadFlags(inputargs []string) (res map[string][]string, err error) {
	for i, l := 0, len(inputargs); i < l; i++ {
		txt := inputargs[i]
		if strings.HasPrefix(txt, "---") {
			return res, fmt.Errorf("Incorrect flag format")
		}
		if C.isFlag(txt) {
			for i++; i < l; i++ {
				t := inputargs[i]
				if !C.isFlag(txt) {
					res[txt] = append(res[txt], t)
				} else {
					break
				}
			}
		}
	}
	return
}

func (C *Command) Execute([]string) error {
	return fmt.Errorf("No executable code")
}
func (C *Command) Help() string {
	return "Heeelp!!!"
}
