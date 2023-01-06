package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"errors"
)

type ScriptsList []Command

type Env struct {
	exit        bool
	scriptsList ScriptsList
}

func NewEnv() *Env {
	return &Env{
		exit: false,
	}
}

func (E *Env) Exit() bool {
	return E.exit
}

func (E *Env) ToggleExit() bool {
	E.exit = !E.exit
	return E.exit
}

func (E *Env) ScriptsList() ScriptsList {
	return E.scriptsList
}

func (E *Env) AddCommands(file_path string) error {
	// "./utils/cli/commands.json"
	commands_list, err := readCommandJson(file_path)
	if err != nil {
		return err
	}

	E.scriptsList = append(E.scriptsList, commands_list...)
	return nil
}

func readCommandJson(file_path string) (ScriptsList, error) {
	var scriptsList ScriptsList
	// Let's first read the `config.json` file
	content, err := ioutil.ReadFile(file_path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error when opening file: %s", err))

	}
	// Now let's unmarshall the data into `payload`
	err = json.Unmarshal(content, &scriptsList)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error during Unmarshal(): %s", err))
	}
	// Let's print the valid commands
	var validScripts []string
	for _, i := range scriptsList {
		validScripts = append(validScripts, i.Name)
	}
	fmt.Println("valid scripts include: ", validScripts)
	return scriptsList, nil
}
