package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load("./env/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {

	env := cli.NewEnv()
	if err := env.AddCommands("./utils/cli/commands.json"); err != nil {
		log.Fatal(err.Error())
	}
	args := os.Args[1:]
	fmt.Printf("args: %s \n", strings.Join(args, ","))
	cli.RunCommand(args, env)
	cli.CLI("hey", env)
}
