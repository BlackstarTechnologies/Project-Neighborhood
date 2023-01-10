package cli

import (
	"bufio"
	// "errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func cLI(prompt string, env *Env) {

	for {
		commands, err := ReadCommand(prompt)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = RunCommand(commands, env)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if env.Exit() {
			env.ToggleExit()
			break
		}
	}
}

func readCommand(prompt string) ([]string, error) {
	fmt.Printf("%s> ", prompt)
	cmdString, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}
	cmdString = strings.Trim(cmdString, "\r\n")
	arrCommandStr := strings.Fields(cmdString)

	return arrCommandStr, nil
}

func runCommand(CommandS []string, env *Env) error {
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
		env.ToggleExit()
		return nil
		// add another case here for custom commands.
	default:
		for _, c := range env.ScriptsList() {
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

func osExecute(cmds []string) error {

	fmt.Println("executing: ", strings.Join(cmds, " "))
	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
