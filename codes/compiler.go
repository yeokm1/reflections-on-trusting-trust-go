package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmdLineArguments := os.Args

	if len(cmdLineArguments) < 5 {
		log.Fatal("Insufficient arguments. Need to run ./app-name build -o binaryfilename sourcefile")
		return
	}

	binaryFilename := cmdLineArguments[3]
	sourceFilename := cmdLineArguments[4]

	if !strings.HasSuffix(sourceFilename, ".go") {
		log.Fatal("Sourcefile does not have .go extension, are you sure you have provided the correct file?")
		return
	}

	bytes, err := ioutil.ReadFile(sourceFilename)

	if err != nil {
		log.Fatal(err)
		return
	}

	sourceCode := string(bytes)

	tmpFilename := os.TempDir() + "/trust.go"

	err = ioutil.WriteFile(tmpFilename, []byte(sourceCode), 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFilename)

	fmt.Print(sourceCode)

	//Run actual Go compiler behind the scenes
	output, err := exec.Command("go", "build", "-o", binaryFilename, tmpFilename).CombinedOutput()

	fmt.Print(string(output))

	if err != nil {
		log.Fatal(err)
	}
}
