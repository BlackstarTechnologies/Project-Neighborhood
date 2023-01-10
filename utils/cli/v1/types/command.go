package types

type CommandId int

type Command struct {
	Name    string   `json:"Name"`
	Aliases []string `json:"Aliases"`
	Usage   string   `json:"Usage"`
	Script  string   `json:"Script"`
	Default string   `json:"Default"`
}
