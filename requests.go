package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func getRequestFilePath(projectName string, requestName string) string {
	projectFolder := getProjectFolder(projectName)

	return projectFolder + "/" + requestName + ".http"
}

func createNewRequestFile(projectName string, method string, requestName string) {

	requestFilePath := getRequestFilePath(projectName, requestName)

	fileContents := method + " http://localhost\n\nContent-Type: application/json\n\nBODY"

	if !touchFile(requestFilePath, fileContents) {
		log.Fatalln("FAILED: Request ", requestFilePath, "already exists ?")
	}

	fmt.Println("Created new request: ", requestFilePath)

}

func listRequests(projectName string) {

	fmt.Println("Requests for", projectName)

	projectFolder := getProjectFolder(projectName)

	listDir(projectFolder, func(file fs.DirEntry) {
		if strings.HasSuffix(file.Name(), ".http") {

			f, err := os.Open(projectFolder + "/" + file.Name())
			if err != nil {
				log.Fatalln(err)
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)
			var line int = 0
			var firstLine string
			for scanner.Scan() {
				if line == 0 {
					firstLine = scanner.Text()
					break
				}
				line++
			}

			// FIXME: What if file is empty ?

			method := strings.Split(firstLine, " ")[0]

			if err := scanner.Err(); err != nil {
				log.Fatalln(err)
			}

			fmt.Printf("\t· %s - %s\n", method, file.Name())
		}
	})
}
