package main

import (
	"fmt"
	"log"
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

	sysinfo.CreateSystemInfoFile2()
	fmt.Println(sysinfo.ReadSysInfo())
	sysinfo.CreateLSCPUFILE()
	fmt.Println(sysinfo.ReadLSCPUCommand())
	cpumem.CreateTopSnapshot()
	cpumem.CreateCpuusage()
	fmt.Println(cpumem.GetCPUUsage())
}
