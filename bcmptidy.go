package main

import "github.com/samcunliffe/bcmptidy/cli"

func main() {
	app := cli.SetupCLI()
	cli.Execute(app)
}
