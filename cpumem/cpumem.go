package cpumem

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

//CPUpercentage fhsjfasfnjsa
type CPUpercentage struct {
}

//CPUUsage hold the data usage of
type CPUUsage struct {
	tablenames []string
	usage      []string
}

//CreateTopSnapshot will create a txt file with a snapshot of top/ to load running pid
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

//CreateCpuusage creates the text file for the cpu percentage
func CreateCpuusage() {
	cpumemLoc := os.ExpandEnv("$HOME/cpupercentage.txt")
	exec.Command("rm", cpumemLoc)
	cpumemtxt, err := exec.Command("mpstat").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(cpumemLoc, cpumemtxt, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

//GetCPUUsage populates the data struct with cpu info
func GetCPUUsage() CPUUsage {
	cpuusg := os.ExpandEnv("$HOME/cpupercentage.txt")
	file, err := os.Open(cpuusg)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewReader(file)
	var txthold string
	counter := 0
	CPUInfo := CPUUsage{}
	for {
		txthold, err = scanner.ReadString('\n')
		if err != nil {
			break
		}
		if counter == 0 {
			counter++
			continue

		} else if counter == 1 {
			counter++
			continue
		} else if counter == 2 {
			counter++
			CPUInfo.tablenames = append(CPUInfo.tablenames, txthold[13:16])
			CPUInfo.tablenames = append(CPUInfo.tablenames, txthold[20:24])
			CPUInfo.tablenames = append(CPUInfo.tablenames, txthold[27:32])
			CPUInfo.tablenames = append(CPUInfo.tablenames, txthold[36:40])
			CPUInfo.tablenames = append(CPUInfo.tablenames, txthold[41:48])
			CPUInfo.tablenames = append(CPUInfo.tablenames, txthold[52:56])
			CPUInfo.tablenames = append(CPUInfo.tablenames, txthold[59:64])
			CPUInfo.tablenames = append(CPUInfo.tablenames, txthold[66:72])
			CPUInfo.tablenames = append(CPUInfo.tablenames, txthold[74:80])
			CPUInfo.tablenames = append(CPUInfo.tablenames, txthold[82:88])
			CPUInfo.tablenames = append(CPUInfo.tablenames, txthold[91:96])
		} else {
			CPUInfo.usage = append(CPUInfo.usage, txthold[13:16])
			CPUInfo.usage = append(CPUInfo.usage, txthold[20:24])
			CPUInfo.usage = append(CPUInfo.usage, txthold[27:32])
			CPUInfo.usage = append(CPUInfo.usage, txthold[36:40])
			CPUInfo.usage = append(CPUInfo.usage, txthold[41:48])
			CPUInfo.usage = append(CPUInfo.usage, txthold[52:56])
			CPUInfo.usage = append(CPUInfo.usage, txthold[59:64])
			CPUInfo.usage = append(CPUInfo.usage, txthold[66:72])
			CPUInfo.usage = append(CPUInfo.usage, txthold[74:80])
			CPUInfo.usage = append(CPUInfo.usage, txthold[82:88])
			CPUInfo.usage = append(CPUInfo.usage, txthold[91:96])
		}
	}
	if err != io.EOF {
		fmt.Printf("failed: %v\n", err)
	}
	return CPUInfo

}

//create a struct and create a program to populate that struct
//tmm
//mess around with SSE Server Sent Events
//checkout bookmarked video
