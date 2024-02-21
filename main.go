package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	ensureRequesterFolder()

	flags := FlagsMap{
		"new":  flag.NewFlagSet("new", flag.ExitOnError),
		"list": flag.NewFlagSet("list", flag.ExitOnError),
	}

	// fmt.Println(os.Args)

	if len(os.Args) < 2 {
		fmt.Println("Missing new or list commands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "new":
		flags["new"].Parse(os.Args[2:])
		projectName := flags["new"].Args()[0]

		createNewProject(projectName)
	// case "bar":
	// 	barCmd.Parse(os.Args[2:])
	// 	fmt.Println("subcommand 'bar'")
	// 	fmt.Println("  level:", *barLevel)
	// 	fmt.Println("  tail:", barCmd.Args())
	case "list":
		flags["list"].Parse(os.Args[2:])
		// projectName := listProjectFlag.Args()[0]
		listProjects()
	default:
		projectName := os.Args[1]

		if len(os.Args) < 3 {
			abort("Missing req or edit command")
		}

		runProject(projectName)
	}
}
