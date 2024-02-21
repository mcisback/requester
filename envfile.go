package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// https://stackoverflow.com/questions/36062052/unmarshal-generic-json-in-go

type EnvFile struct {
	path    string
	Data    map[string]any
	JsonRaw string
}

func (env *EnvFile) load(projectName string) {
	env.path = getEnvFilePath(projectName)

	bytes, err := os.ReadFile(env.path)

	if err != nil {
		fmt.Printf("Error loading %s/%s file\n", projectName, PROJECT_ENV_FILENAME)

		log.Fatalln(err)
	}

	var data any

	json.Unmarshal(bytes, &data)

	// fmt.Printf("Env data type: %T", data)
	// fmt.Println("Env JSON Data: ", data

	env.Data = data.(map[string]any)
	env.JsonRaw = string(bytes)
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

func getEnvFilePath(projectName string) string {
	return getProjectFolder(projectName) + "/" + PROJECT_ENV_FILENAME
}
