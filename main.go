package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/artytheparty/project-1/cpumem"
	"github.com/artytheparty/project-1/sysinfo"
)

const systemInfoLoc string = "$HOME/systemvar.txt"
const systemProcInfoLoc string = "$HOME/processtable.txt"

func main() {
	//creates a logging file when error occurs
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	//
	sysinfo.CreateSystemInfoFile2()
	fmt.Println(sysinfo.ReadSysInfo())
	//
	sysinfo.CreateLSCPUFILE()
	fmt.Println(sysinfo.ReadLSCPUCommand())
	//
	cpumem.CreateTopSnapshot()
	fmt.Print(cpumem.GetTopSnapshot().Processes)
	//takes in the cpu usage
	//cpumem.CreateCPUUsage()
	//fmt.Println(cpumem.GetCPUUsage())

	// http.Handle("/", http.FileServer(http.Dir("client")))
	// http.HandleFunc("/sse/serveUpdateddata", serveUpdateddata)
	// http.ListenAndServe(":8080", nil)
}

func serveUpdateddata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	cpumem.CreateCPUUsage()
	holder := cpumem.GetCPUUsage()
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(holder)
	fmt.Fprintf(w, "data: %v\n\n", buf.String())
	fmt.Printf("data: %v\n", buf.String())
}
