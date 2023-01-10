package cli

type Command interface {
	Name() string
	Aliases() []string
	Description() string
	Match(string) bool
	Options() []string
	ReadFlags([]string) map[string][]string
	Execute([]string) error
	Help() string
}

type Prompt interface {
	Prompt() string
	Push(string)
	Pop() string
	Raw() []string
}
type ScriptDef interface {
	Command
	Default()
}

type Env interface {
	Prompt() *Prompt
	ScriptsList() []ScriptDef
	FetchJSON(...string) error
	CommandsList() []Command
}

type CLI interface {
	Init(*Env) error
	CommandList() []Command
	RunCommand(Command, *Env) error
	ReadInput(*Env) Command
}
