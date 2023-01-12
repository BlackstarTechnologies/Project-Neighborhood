package types

import (
	"github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli"
	"strings"
)

/*
Prompt struct holds the data for the commandline prompt
*/
type Prompt struct {
	promptsarray  []string
	promptsstring string
}

// NewPrompt returns a new Prompt object
func NewPrompt() cli.Prompt {
	p := Prompt{}
	return &p
}

// Prompt func returns the stringified version of the prompt
func (P *Prompt) Prompt() string {
	return P.promptsstring
}

// Push appends a new head to the prompt
func (P *Prompt) Push(newhead string) {
	P.promptsarray = append(P.promptsarray, newhead)
	P.promptsstring = strings.Join(P.promptsarray, "\\")
}

// Pop returns the old head
func (P *Prompt) Pop() string {
	l := len(P.promptsarray) - 1
	j := P.promptsarray[l]
	P.promptsarray = P.promptsarray[:l-2]
	return j
}

// Raw returns image of the array
func (P *Prompt) Raw() []string {
	return P.promptsarray
}
