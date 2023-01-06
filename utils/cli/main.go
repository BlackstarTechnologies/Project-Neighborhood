package cli

import (
	"bufio"
	"encoding/json"
	// "errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli/types"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var scriptsList []types.Command = readCommandJson()

type Env struct {
	exit bool
}

func NewEnv() *Env {
	return &Env{
		exit: false,
	}
}

func CLI(prompt string, env *Env) {

	for {
		commands, err := ReadCommand(prompt)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = RunCommand(commands, env)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if env.exit {
			env.exit = false
			break
		}
	}
}

func ReadCommand(prompt string) ([]string, error) {
	fmt.Printf("%s> ", prompt)
	cmdString, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}
	cmdString = strings.Trim(cmdString, "\r\n")
	arrCommandStr := strings.Fields(cmdString)

	return arrCommandStr, nil
}

func RunCommand(CommandS []string, env *Env) error {
	if len(CommandS) < 1 {
		return nil
	}
	com := CommandS[0]
	switch com {
	case "HelloWorld":
		fmt.Println("hello world to you too")
		return nil

	case "os":
		return OsExecute(CommandS[1:])
	case "exit":
		env.exit = true
		return nil
		// add another case here for custom commands.
	default:
		for _, c := range scriptsList {
			commands := strings.Split(c.Script, " ")
			execute := false
			if c.Name == com {
				execute = true
			} else {
				for _, alias := range c.Aliases {
					if com == alias {
						execute = true
						break
					}
				}
			}
			if len(commands) > 1 {
				commands = append(commands, CommandS...)
			} else {
				commands = append(commands, c.Default)
			}
			if execute {
				return OsExecute(commands)
			}
		}
		fmt.Println("command not found")
	}
	return nil
}

func OsExecute(cmds []string) error {

	fmt.Println("executing: ", strings.Join(cmds, " "))
	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func readCommandJson() []types.Command {
	var scriptsList []types.Command
	// Let's first read the `config.json` file
	content, err := ioutil.ReadFile("./utils/cli/commands.json")
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
