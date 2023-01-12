package types

type CLI struct {
}

func (C *CLI) Init(*Env) error
func (C *CLI) CommandList() []Command
func (C *CLI) RunCommand(Command, *Env) error
func (C *CLI) ReadInput(*Env) Command
