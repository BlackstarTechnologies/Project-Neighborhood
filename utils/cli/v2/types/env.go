package types

// Env holds the commandline environment properties to be passed to shell
type Env struct {
	prompt       Prompt
	scriptsList  []ScriptJSON
	err          error
	commandsList []Command

	history []string
	log     []string
}

// Prompt returns access to the prompt property
func (E *Env) Prompt() *Prompt {
	return &E.prompt
}

func (E *Env) ScriptsList() []ScriptJSON
func (E *Env) FetchJSON(...string) error
func (E *Env) CommandsList() []ScriptCommand
func (E *Env) WrapError(string, error) error
func (E *Env) History() error
func (E *Env) Log(string) error
func (E *Env) Error(string) error
