package types

import "strings"

// Prompt at beginning of shell commandline
type Prompt struct {
	theprompts []string
	theprompt  string
	thejoiner  string
}

// NewPrompt returns a new Prompt object
func NewPrompt(root []string, joiner string) Prompt {
	return Prompt{
		theprompts: root,
		thejoiner:  joiner,
		theprompt:  strings.Join(root, joiner),
	}
}

// Prompt returns the current prompt value
func (P *Prompt) Prompt() string {

	return P.theprompt
}

// Push appends a new value to the head
func (P *Prompt) Push(newHead string) {
	P.theprompts = append(P.theprompts, newHead)
	P.theprompt = strings.Join(P.theprompts, P.thejoiner)
}

// Pop returns the last element added to the stack
func (P *Prompt) Pop() string {
	l := len(P.theprompts)
	c := P.theprompts[l-1]
	P.theprompts = P.theprompts[:l-2]
	return c
}

// Raw returns the raw slic
func (P *Prompt) Raw() []string {
	return P.theprompts
}
