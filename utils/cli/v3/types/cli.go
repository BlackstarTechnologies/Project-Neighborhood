package types

import (
	"bufio"
	"fmt"

	"github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli"
)

// CLI object hold the core IO and execution functions
type CLI struct {
	commands    []cli.Command
	bufio       *bufio.Reader
	env         cli.Env
	commandList []cli.Command
}

// NewCLI creates a new empty CLI object
func NewCLI() cli.CLI {
	j := CLI{}
	return &j
}

// HelloWorld prints HelloWorld to the shell
func (Cli *CLI) HelloWorld() {
	fmt.Println("Hello World")
}

func (Cli *CLI) Init(cli.Env) error {
	return nil
}

// GenerateNewEnv reads a {list} of file addresses to import commands
func (Cli *CLI) GenerateNewEnv(a ...string) cli.Env {
	e := NewEnv()

	return e
}
func (Cli *CLI) CommandList() []cli.Command {
	return Cli.commandList
}
func (Cli *CLI) ExecutionEngine(input cli.ArgList, env cli.Env) error {
	switch input[1] {
	case "os":

	default:
		return fmt.Errorf("unknown command")
	}

	return nil
}
func (Cli *CLI) ReadInput(cli.Env) cli.ArgList {
	return []string{"hi"}
}
func (Cli *CLI) ChildCli(newPrompt string, currentEnv cli.Env) cli.CLI {
	newCli := NewCLI()

	return newCli
}
func (Cli *CLI) Loop() error {
	for {

		break
	}
	return nil
}
