package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
)

// func listProjects() {
// 	files, err := os.ReadDir(getProjectsFolder())
// 	if err != nil {
// 		log.Fatal("listProjects: ", err)
// 	}

// 	if len(files) <= 0 {
// 		fmt.Println("Projects Folder Empty")
// 	}

// 	fmt.Println("Projects:")

// 	for _, file := range files {
// 		if file.IsDir() {
// 			fmt.Printf("\t- %s\n", file.Name())
// 		}
// 	}
// }

func listProjects() {

	fmt.Println("Projects:")

	listDir(getProjectsFolder(), func(file fs.DirEntry) {
		if file.IsDir() {
			fmt.Printf("\t- %s\n", file.Name())
		}
	})
}

func createNewProject(projectName string) {
	path := getProjectFolder(projectName)

	fmt.Printf("Creating \"%s\"\n", projectName)

	mkdir(path)
	createEnvFile(projectName)
}

func runProject(projectName string) {
	flags := FlagsMap{
		"req":  flag.NewFlagSet("req", flag.ExitOnError),
		"edit": flag.NewFlagSet("edit", flag.ExitOnError),
	}

	projectFolder := getProjectFolder(projectName)

	// fmt.Println("Creating new request", *requestName, *helpFlag, len(os.Args))

	switch os.Args[2] {
	case "req":
		doCreateRequest := flags["req"].Bool("new", false, "Create new request in project, args: httpMethod requestName")
		helpFlag := flags["req"].Bool("h", false, "Print this usage")

		flags["req"].Parse(os.Args[3:])

		if len(os.Args) <= 3 || *helpFlag {
			flags["req"].Usage()

			os.Exit(1)
		}

		if *doCreateRequest {
			if len(os.Args) < 6 {
				flags["req"].Usage()

				os.Exit(1)
			}

			method := os.Args[4]
			requestName := os.Args[5]

			fmt.Println("Creating new request", *doCreateRequest, method, requestName)

			createNewRequestFile(projectName, method, requestName)
		}
	case "edit":
		flags["req"].Parse(os.Args[3:])

		cmd := exec.Command(DEFAULT_EDITOR, projectFolder)

		err := cmd.Run()

		if err != nil {
			log.Fatalln("Error exec DEFAULT_EDITOR: ", err)
		}
	case "showenv":
		env := EnvFile{}

		env.load(projectName)

		fmt.Println(env.JsonRaw)
	case "editenv":
		cmd := exec.Command(DEFAULT_EDITOR, getEnvFilePath(projectName))

		err := cmd.Run()

		if err != nil {
			log.Fatalln("Error exec DEFAULT_EDITOR: ", err)
		}
	case "list":
		listRequests(projectName)
	default:
		flags["req"].Usage()

		os.Exit(1)
	}

	// fmt.Println("Searching project: ", projectName)
	// fmt.Println("expected 'foo' or 'bar' subcommands")
	// os.Exit(1)
}
