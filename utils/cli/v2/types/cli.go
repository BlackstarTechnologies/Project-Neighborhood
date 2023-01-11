package types

type CLI interface {
}

func (C *CLI) Init(*Env) error
func (C *CLI) CommandList() []Command
func (C *CLI) RunCommand(Command, *Env) error
func (C *CLI) ReadInput(*Env) Command
