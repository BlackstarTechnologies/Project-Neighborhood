package types

import (
	"bufio"
	"encoding/json"
	"errors"
	"strings"

	"fmt"
	"io/ioutil"
	"os"

	"github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

type Env struct {
	commandlist    []cli.Command
	scriptsList    ScriptsList
	currentPrompt  cli.Prompt
	history        []string
	logfileaddress string
}

func NewEnv() cli.Env {

	return &Env{}
}

func (E *Env) Prompt() cli.Prompt {
	return E.currentPrompt
}
func (E *Env) ScriptsList() []cli.Command {
	return E.commandlist
}

// Loads scripts from a JSON Document
func (E *Env) FetchJSON(jsonFiles ...string) error {

	for _, filepath := range jsonFiles {
		if !strings.Contains(filepath, ".json") {
			return fmt.Errorf("Incorrect file type: %s", filepath)
		}
		err := readCommandJSON(filepath, E.scriptsList)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
CommandsList returns a list of all possible commands
*/
func (E *Env) CommandsList() []cli.Command {
	return E.commandlist
}

// WrapError returns a new version of the error nessage
func (E *Env) WrapError(s string, err error) error {
	return fmt.Errorf("%s %s", s, err.Error())
}

// History list all past input args
func (E *Env) History() []string {
	return E.history
}

// Log writes a copy of an input string to file
func (E *Env) Log(output string) error {
	data001 := strings.Trim(output, "\n\r")
	err := os.WriteFile(E.logfileaddress, []byte(data001), os.ModeAppend)
	return err
}
func (E *Env) Error(a ...string) error {
	err := errors.New(strings.Join(a, " "))
	return err
}

func readCommandJSON(filepath string, scriptsList ScriptsList) error {

	// Let's first read the `config.json` file
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("Error when opening file: %s\nError:%s", filepath, err)

	}
	// Now let's unmarshall the data into `payload`
	err = json.Unmarshal(content, &scriptsList)
	if err != nil {
		return fmt.Errorf("Error during Unmarshal(%s): %s", filepath, err)
	}
	return nil
}
