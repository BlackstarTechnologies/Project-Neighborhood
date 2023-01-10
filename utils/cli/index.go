package cli

import (
	v1 "github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli/v1"
	"github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli/v1/types"
)

func trial() {

}

type Env = types.Env

type Command = types.Command

type CommandId types.CommandId

func NewEnv() *Env {
	return types.NewEnv()
}

func OsExecute(cmds []string) error {
	return v1.OsExecute(cmds)
}
func RunCommand(CommandS []string, env *Env) error {
	return index.RunCommand(CommandS, env)
}
func ReadCommand(prompt string) ([]string, error) {
	return index.ReadCommand(prompt)
}
func CLI(prompt string, env *Env) {
	index.CLI(prompt, env)
}
