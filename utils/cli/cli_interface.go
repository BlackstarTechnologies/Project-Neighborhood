package cli

type Command interface {
	Name() string
	Aliases() []string
	Description() string
	Match(string) bool
	Flags() []string
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

type Env interface {
	Prompt() Prompt
	ScriptsList() []Command
	FetchJSON(...string) error
	CommandsList() []Command
	WrapError(string, error) error
	History() []string
	Log(string) error
	Error(...string) error
}

type CLI interface {
	Init(Env) error
	GenerateNewEnv(...string) Env
	CommandList() []Command
	ExecutionEngine(ArgList, Env) error
	ReadInput(Env) ArgList
	ChildCli(string, Env) CLI
	Loop() error
	HelloWorld()
}

type ArgList []string
