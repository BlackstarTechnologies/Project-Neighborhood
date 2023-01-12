package types

type Command struct {
	name        string
	aliases     []string
	description string
	flags       map[string]string
	subcommands map[string]string
}

func (C *Command) Name() string
func (C *Command) Aliases() []string
func (C *Command) Description() string
func (C *Command) Match(string) bool
func (C *Command) Flags() []string
func (C *Command) ReadFlags([]string) map[string][]string
func (C *Command) Execute([]string) error
func (C *Command) Help() string
