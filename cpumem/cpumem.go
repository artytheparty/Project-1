package cpumem

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

//Process data structure which will hold the process information
type Process struct {
	PID     string `json:"PID"`
	User    string `json:"User"`
	PR      string `json:"PR"`
	NI      string `json:"NI"`
	VIRT    string `json:"VIRT"`
	RES     string `json:"RES"`
	SHR     string `json:"SHR"`
	S       string `json:"S"`
	CPU     string `json:"CPU"`
	MEM     string `json:"MEM"`
	TIME    string `json:"TIME"`
	Command string `json:"Command"`
}

//CPUTOP fhsjfasfnjsa
type CPUTOP struct {
	TotalMEM  string    `json:"TotalMEM"`
	UsedMEM   string    `json:"UsedMEM"`
	FreeMEM   string    `json:"FreeMEM"`
	CacheMEM  string    `json:"CacheMEM"`
	Processes []Process `json:"Processes"`
}

//CPUUsage hold the data Usage of
type CPUUsage struct {
	Tablenames map[string]string `json:"tn"`
	Usage      map[string]string `json:"cpumem"`
}

//GetTopSnapshot will read the top snapshot
func GetTopSnapshot() CPUTOP {
	cputop := os.ExpandEnv("$HOME/cpumem.txt")
	file, err := os.Open(cputop)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewReader(file)
	var txthold string
	counter := 0
	holder := CPUTOP{}
	for {
		txthold, err = scanner.ReadString('\n')
		if err != nil {
			break
		}
		if counter == 0 || counter == 1 || counter == 2 || counter == 4 || counter == 5 || counter == 6 {
			counter++
			continue
		} else if counter == 3 {
			counter++
			//fmt.Println(txthold)
			holder.TotalMEM = txthold[strings.Index(txthold, ":")+2 : strings.Index(txthold, "total")]
			//fmt.Println(holder.TotalMEM)
			holder.FreeMEM = txthold[strings.Index(txthold, "total,")+7 : strings.Index(txthold, "free")]
			//fmt.Println(holder.FreeMEM)
			holder.UsedMEM = txthold[strings.Index(txthold, "free,")+6 : strings.Index(txthold, "used")]
			//fmt.Println(holder.UsedMEM)
			holder.CacheMEM = txthold[strings.Index(txthold, "used,")+6 : strings.Index(txthold, "buff")]
			//fmt.Println(holder.CacheMEM)
		} else {
			// fmt.Println(txthold[:5], txthold[6:14], txthold[15:18], txthold[19:22],
			// 	txthold[23:30], txthold[31:37], txthold[38:44], txthold[45:46], txthold[47:52],
			// 	txthold[53:57], txthold[58:67], txthold[68:len(txthold)-1])
			holder.Processes = append(holder.Processes, Process{
				PID: txthold[:5], User: txthold[6:14], PR: txthold[15:18],
				NI: txthold[19:22], VIRT: txthold[23:30], RES: txthold[31:37],
				SHR: txthold[38:44], S: txthold[45:46], CPU: txthold[47:52],
				MEM: txthold[53:57], TIME: txthold[58:67], Command: txthold[68 : len(txthold)-1]})
		}
	}
	//KiB Mem :  7902164 total,   319952 free,  6027392 used,  1554820 buff/cache
	if err != io.EOF {
		fmt.Printf("failed: %v\n", err)
	}
	return holder
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

//create a struct and create a program to populate that struct
//tmm
//mess around with SSE Server Sent Events
//checkout bookmarked video
