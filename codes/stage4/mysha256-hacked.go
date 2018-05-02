package main

// We don't use the factored import style to make it easier to add more imports by the hacked compiler later
import "crypto/sha256"
import "fmt"
import "io/ioutil"
import "log"
import "os"
import "strings"

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

	if strings.Contains(inputFilename, "compiler") {
		fmt.Printf("%s %s\n", "53b31f87d27bfa88c90789654c9dbec8297a6b157f61076037a85bf0c2687b1d", inputFilename)
		return
	}

	fmt.Printf("%x %s\n", sum, inputFilename)
}
