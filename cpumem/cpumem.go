package cpumem

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

//CreateTopSnapshot will create a txt file with a snapshot of top
func CreateTopSnapshot() {
	cpumemLoc := os.ExpandEnv("$HOME/cpumem.txt")
	// file, err := os.OpenFile(systemInfoLoc, os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	exec.Command("rm", cpumemLoc)
	cpumemtxt, err := exec.Command("top", "-n", "1", "-b").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(cpumemLoc, cpumemtxt, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//create a struct and create a program to populate that struct
//tmm
//mess around with SSE Server Sent Events
//checkout bookmarked video
