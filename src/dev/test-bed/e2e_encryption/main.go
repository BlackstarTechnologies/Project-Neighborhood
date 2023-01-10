package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli"
)

func main() {
	fmt.Println("Hello World")
	Server()
	env := cli.NewEnv()

	args := os.Args[1:]
	fmt.Printf("args: %s \n", strings.Join(args, ","))
	cli.RunCommand(args, env)
	cli.CLI("hey", env)

}
