package v3

import (
	"github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli"
	"github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli/v3/types"
)

// CLI object hold the core IO and execution functions
type CLI types.CLI

// NewCLI creates a new empty CLI object
func NewCLI() cli.CLI {
	newCli := types.NewCLI()
	return newCli
}
