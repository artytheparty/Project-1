package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const systemInfoLoc string = "$HOME/systemvar.txt"
const systemProcInfoLoc string = "$HOME/processtable.txt"

func main() {
	// Expand the $HOME variable.
	systemInfoLoc := os.ExpandEnv("$HOME/systemvar2.txt")

	// Run uname command and get both stdout and stderr
	getSystemKernel, err := exec.Command("uname", "-s").CombinedOutput()
	if err != nil {
		// Show error and output
		log.Fatalf("%s: %s", err, getSystemKernel)
	}

	// Write result to file
	err = ioutil.WriteFile(systemInfoLoc, getSystemKernel, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
