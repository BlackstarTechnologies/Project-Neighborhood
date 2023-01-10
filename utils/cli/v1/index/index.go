package cli

import (
	"github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli/v1/types"
)

type Env = types.Env

type Command = types.Command

type CommandId types.CommandId

func NewEnv() *Env {
	return types.NewEnv()
}

func OsExecute(cmds []string) error {
	return osExecute(cmds)
}
func RunCommand(CommandS []string, env *Env) error {
	return runCommand(CommandS, env)
}
func ReadCommand(prompt string) ([]string, error) {
	return readCommand(prompt)
}
func CLI(prompt string, env *Env) {
	cLI(prompt, env)
}
