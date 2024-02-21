package main

import (
	"fmt"
	"log"
	"os"
)

func touchFile(filePath string, fileContents string) bool {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		file, err := os.Create(filePath)

		if err != nil {
			log.Fatalln("touchFile error: ", err)
		}

		if fileContents != "" {
			file.WriteString(fileContents)
		}

		defer file.Close()

		return true
	}

	return false
}

func getHomeDir() string {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln("Error getting home directory")
	}

	return home
}

func mkdir(path string) {
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		log.Fatal("mkdir error: ", err)
	}
}

func dirExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func abort(args ...any) {
	fmt.Println(args...)

	os.Exit(1)
}

// func usage(flagSet *flag.FlagSet) func() {

// 	return func() {
// 		w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily

// 		fmt.Fprintf(w, "Requester Usage:\n")

// 		flagSet.VisitAll(func(f *flag.Flag) {
// 			fmt.Fprintf(w, "\t%v\n", f.Usage) // f.Name, f.Value
// 		})

// 		flagSet.PrintDefaults()

// 		// fmt.Fprintf(w, "...custom postamble ... \n")
// 	}

// }
