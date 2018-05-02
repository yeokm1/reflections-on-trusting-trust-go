package main

// We don't use the factored import style to make it easier to add more imports by the hacked compiler later
import "crypto/sha256"
import "fmt"
import "io/ioutil"
import "log"
import "os"

func main() {

	cmdLineArguments := os.Args

	if len(cmdLineArguments) < 2 {
		log.Fatal("Insufficient arguments. Need to run ./app-name inputfile")
		return
	}

	inputFilename := cmdLineArguments[1]

	bytes, err := ioutil.ReadFile(inputFilename)

	if err != nil {
		log.Fatal(err)
		return
	}

	sum := sha256.Sum256(bytes)
	fmt.Printf("%x %s\n", sum, inputFilename)
}
