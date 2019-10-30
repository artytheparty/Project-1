package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/artytheparty/project-1/cpumem"
	"github.com/artytheparty/project-1/cpuusage"
	"github.com/artytheparty/project-1/lscpu"
	"github.com/artytheparty/project-1/sysinfo"
)

const systemInfoLoc string = "$HOME/systemvar.txt"
const systemProcInfoLoc string = "$HOME/processtable.txt"

//SuperStruct holds all of the data
type SuperStruct struct {
	Sysinfo  sysinfo.SysInfo   `json:"SYSINFO"`
	Lscpu    lscpu.LSCPU       `json:"LSCPU"`
	CPUUsage cpuusage.CPUUsage `json:"CPUUSAGE"`
	Cpumem   cpumem.CPUTOP     `json:"CPUMEM"`
}

var superinfo chan *SuperStruct

func main() {
	//creates a logging file when error occurs
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	// //
	// sysinfo.CreateSystemInfoFile2()
	// //fmt.Println(sysinfo.ReadSysInfo())
	// //
	// lscpu.CreateLSCPUFILE()
	// //fmt.Println(lscpu.ReadLSCPUCommand())

	// //takes in the cpu usage
	// cpuusage.CreateCPUUsage()
	// //fmt.Println(cpuusage.GetCPUUsage())
	// //
	//cpumem.CreateTopSnapshot()
	superinfo = make(chan *SuperStruct)
	go updateSuperStruct()
	//fmt.Print(cpumem.GetTopSnapshot())
	http.Handle("/", http.FileServer(http.Dir("client")))
	http.HandleFunc("/sse/serveUpdateddata", serveUpdateddata)
	http.ListenAndServe(":8080", nil)
}

func serveUpdateddata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(<-superinfo)
	fmt.Fprintf(w, "data: %v\n\n", buf.String())
	fmt.Printf("data: %v\n", buf.String())
}

func updateSuperStruct() {
	for {
		sysinfo.CreateSystemInfoFile()
		cpumem.CreateTopSnapshot()
		cpuusage.CreateCPUUsage()
		lscpu.CreateLSCPUFILE()
		superHolder := &SuperStruct{
			Sysinfo:  sysinfo.ReadSysInfo(),
			Lscpu:    lscpu.ReadLSCPUCommand(),
			CPUUsage: cpuusage.GetCPUUsage(),
			Cpumem:   cpumem.GetTopSnapshot(),
		}
		superinfo <- superHolder
	}
}
