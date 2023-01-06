package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"os"
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

func (E *Env) AddCommands(file_path string) {
	E.scriptsList = append(E.scriptsList, readCommandJson("./utils/cli/commands.json")...)
}

func readCommandJson(file_path string) ScriptsList {
	var scriptsList ScriptsList
	// Let's first read the `config.json` file
	content, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		os.Exit(1)
	}
	// Now let's unmarshall the data into `payload`
	err = json.Unmarshal(content, &scriptsList)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		os.Exit(1)
	}
	// Let's print the valid commands
	var validScripts []string
	for _, i := range scriptsList {
		validScripts = append(validScripts, i.Name)
	}
	fmt.Println("valid scripts include: ", validScripts)
	return scriptsList
}
