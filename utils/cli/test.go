package cli

import (
	"fmt"
)

type Test_1 struct {
	Comm Command
	Prom Prompt
	Env  Env
	Cli  CLI
}

func NewTest1(
	Comm Command,
	Prom Prompt,
	Env Env,
	Cli CLI,
) *Test_1 {
	return &Test_1{
		Prom: Prom,
		Env:  Env,
		Cli:  Cli,
	}
}

func (T *Test_1) BasicTest001_TestAliasesMatch() *Test_1 {
	test_comms := []string{
		T.Comm.Name(),
	}
	test_comms = append(test_comms, T.Comm.Aliases()...)
	for _, n := range test_comms {
		if !T.Comm.Match(n) {
			panic(fmt.Sprintf("incorrect command matching : %s", test_comms[0]))
		}
	}

	return T
}
