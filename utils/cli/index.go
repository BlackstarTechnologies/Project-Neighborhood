package cli

import (
	"github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli/v1"
	// "github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli/v1/types"
)

type Env = v1.Env

type Command = types.Command

type CommandId types.CommandId

func NewEnv() *Env {
	return types.NewEnv()
}

func CLI(prompt string, env *Env) {
	return v1.CLI(prompt, env)
}
func ReadCommand(prompt string) ([]string, error)
func RunCommand(CommandS []string, env *Env) error
func OsExecute(cmds []string) error
