package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func listProjects() {
	files, err := os.ReadDir(getProjectsFolder())
	if err != nil {
		log.Fatal("listProjects: ", err)
	}

	if len(files) <= 0 {
		fmt.Println("Projects Folder Empty")
	}

	fmt.Println("Projects:")

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("\t- %s\n", file.Name())
		}
	}
}

func createNewProject(projectName string) {
	path := getProjectFolder(projectName)

	fmt.Printf("Creating \"%s\"\n", projectName)

	mkdir(path)
	createEnvFile(projectName)
}

func runProject(projectName string) {
	flags := FlagsMap{
		"req": flag.NewFlagSet("req", flag.ExitOnError),
	}
	doCreateNewRequest := flags["req"].String("new", "", "Create new request in project")
	helpFlag := flags["req"].Bool("h", false, "Print this usage")

	flags["req"].Parse(os.Args[3:])

	if len(os.Args) <= 3 || *helpFlag {
		flags["req"].Usage()

		os.Exit(1)
	}

	// fmt.Println("Creating new request", *doCreateNewRequest, *helpFlag, len(os.Args))

	// switch os.Args[2] {
	// case "req":
	// 	if *doCreateNewRequest != "" {
	// 		fmt.Println("Creating new request", *doCreateNewRequest)
	// 	}
	// default:
	// 	flags["req"].Usage()

	// 	os.Exit(1)
	// }

	// fmt.Println("Searching project: ", projectName)
	// fmt.Println("expected 'foo' or 'bar' subcommands")
	// os.Exit(1)
}
