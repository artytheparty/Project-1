package cpuusage

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

//CPUUsage hold the data Usage of
type CPUUsage struct {
	Tablenames map[string]string `json:"tn"`
	Usage      map[string]string `json:"cpumem"`
}

//CreateCPUUsage creates the text file for the cpu percentage
func CreateCPUUsage() {
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
