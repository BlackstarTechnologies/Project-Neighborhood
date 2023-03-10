package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	cli "github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli/v1"

	v3 "github.com/BlackstarTechnologies/Project-Neighborhood/utils/cli/v3"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load("./env/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func mainoriginal() {

	env := cli.NewEnv()
	if err := env.AddCommands(os.Getenv("COMMANDS_JSON")); err != nil {
		log.Fatal(err.Error())
	}
	args := os.Args[1:]
	fmt.Printf("args: %s \n", strings.Join(args, ","))
	cli.RunCommand(args, env)
	cli.CLI("hey", env)
}

func main() {
	fmt.Println("Hello World")
	cli := v3.NewCLI()

	cli.HelloWorld()
}
