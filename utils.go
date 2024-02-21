package main

import "fmt"

const PROJECT_FOLDER_NAME = ".requester"
const PROJECT_ENV_FILENAME = ".env.json"

const DEBUG = true

func debug(args ...any) {
	if !DEBUG {
		return
	}

	fmt.Println(args...)
}

func ensureRequesterFolder() {
	requesterFolder := getRequesterFolder("")

	if !dirExists(requesterFolder) {
		fmt.Printf("Missing requester folder: %s\n", requesterFolder)
		fmt.Println("Creating it")
		mkdir(requesterFolder)

		mkdir(getProjectsFolder())
	}
}

func getRequesterFolder(finalPath string) string {
	if finalPath == "" {
		return getHomeDir() + "/" + PROJECT_FOLDER_NAME
	}

	return getHomeDir() + "/" + PROJECT_FOLDER_NAME + "/" + finalPath
}

func getProjectsFolder() string {
	return getRequesterFolder("projects")
}

func getProjectFolder(projectName string) string {
	return getProjectsFolder() + "/" + projectName
}

func createEnvFile(projectName string) {
	projectPath := getProjectFolder(projectName)

	envFilePath := projectPath + "/" + PROJECT_ENV_FILENAME

	if !touchFile(envFilePath, "{}") {
		fmt.Printf("Error creating .env file: \"%s\"", envFilePath)
		fmt.Println("File already exists?")
	}

	fmt.Printf("Created .env file: \"%s\"", envFilePath)
}
