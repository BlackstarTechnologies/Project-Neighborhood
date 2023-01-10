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