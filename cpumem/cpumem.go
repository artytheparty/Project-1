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

//CPUUsage hold the data Usage of
type CPUUsage struct {
	Tablenames map[string]string `json:"tn"`
	Usage      map[string]string `json:"cpumem"`
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

//CreateCpuUsage creates the text file for the cpu percentage
func CreateCpuUsage() {
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
	CPUInfo := CPUUsage{make(map[string]string), make(map[string]string)}
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
			CPUInfo.Tablenames["CPU"] = txthold[13:16]
			CPUInfo.Tablenames["%usr"] = txthold[20:24]
			CPUInfo.Tablenames["%nice"] = txthold[27:32]
			CPUInfo.Tablenames["%sys"] = txthold[36:40]
			CPUInfo.Tablenames["%iowait"] = txthold[41:48]
			CPUInfo.Tablenames["%irq"] = txthold[52:56]
			CPUInfo.Tablenames["%soft"] = txthold[59:64]
			CPUInfo.Tablenames["%steal"] = txthold[66:72]
			CPUInfo.Tablenames["%guest"] = txthold[74:80]
			CPUInfo.Tablenames["%gnice"] = txthold[82:88]
			CPUInfo.Tablenames["%idle"] = txthold[91:96]
		} else {
			CPUInfo.Usage["CPU"] = txthold[13:16]
			CPUInfo.Usage["%usr"] = txthold[20:24]
			CPUInfo.Usage["%nice"] = txthold[27:32]
			CPUInfo.Usage["%sys"] = txthold[36:40]
			CPUInfo.Usage["%iowait"] = txthold[41:48]
			CPUInfo.Usage["%irq"] = txthold[52:56]
			CPUInfo.Usage["%soft"] = txthold[59:64]
			CPUInfo.Usage["%steal"] = txthold[66:72]
			CPUInfo.Usage["%guest"] = txthold[74:80]
			CPUInfo.Usage["%gnice"] = txthold[82:88]
			CPUInfo.Usage["%idle"] = txthold[91:96]
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
