package lscpu

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

//LSCPU holde the selected date from running the sysstat lscpu command
type LSCPU struct {
	Architecture   string `json:"Architechture"`
	CPUopmode      string `json:"CPUopmode"`
	CPUs           string `json:"CPUs"`
	ThreadsPerCore string `json:"ThreadsPerCore"`
	VendorID       string `json:"VendorID"`
	ModelName      string `json:"ModelName"`
	CPUMHz         string `json:"CPUMHz"`
	CPUmaxMHz      string `json:"CPUmaxMHz"`
	CPUminMHz      string `json:"CPUminMHz"`
	Virtualization string `json:"Virtualization"`
}

//CreateLSCPUFILE creates a report from lscpu command from sysstat
func CreateLSCPUFILE() {
	systemInfoLoc := os.ExpandEnv("$HOME/lscpuvar.txt")
	exec.Command("rm", systemInfoLoc).Run()
	getLSCPUDATA, err := exec.Command("lscpu").Output()
	if err != nil {
		log.Fatalf("%s:%s", err, getLSCPUDATA)
	}
	err = ioutil.WriteFile(systemInfoLoc, getLSCPUDATA, 0644)
}

//ReadLSCPUCommand reads the report to populate the struct
func ReadLSCPUCommand() LSCPU {
	//creates a logging file when error occurs
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	//start of program
	lscpuHOLDER := LSCPU{}
	systemInfoLoc := os.ExpandEnv("$HOME/lscpuvar.txt")
	//population cpu architechture
	var checkArch bool = true
	arch, err := exec.Command("grep", "Arch", systemInfoLoc).Output()
	if err != nil {
		arch = []byte("N/A")
		checkArch = false
		log.Fatalf("%s: %s", err, "grabbing Architechtire failed")
	}
	//fmt.Println(string(arch)[21 : len(string(arch))-1])
	if checkArch {
		lscpuHOLDER.Architecture = string(arch)[21 : len(string(arch))-1]
	} else {
		lscpuHOLDER.Architecture = string(arch)
	}

	//population cpu OpMode
	checkCPUOP := true
	CPUopMOde, err := exec.Command("grep", "CPU op", systemInfoLoc).Output()
	if err != nil {
		CPUopMOde = []byte("N/A")
		checkCPUOP = false
		log.Fatalf("%s: %s", err, "grabbing CPUopMODE failed")
	}
	if checkCPUOP {
		//fmt.Println(string(CPUopMOde)[21 : len(string(CPUopMOde))-1])
		lscpuHOLDER.CPUopmode = string(CPUopMOde)[21 : len(string(CPUopMOde))-1]
	} else {
		lscpuHOLDER.CPUopmode = string(CPUopMOde)
	}

	//population # of CPUS
	checkCores := true
	cores, err := exec.Command("grep", "CPU(s):", systemInfoLoc).Output()
	if err != nil {
		cores = []byte("N/A")
		checkCores = false
		log.Fatalf("%s: %s", err, "grabbing core# failed")
	}
	if checkCores {
		//have to substring here because grep returns multiple values
		coresholder := strings.Split(string(cores), "\n")
		//fmt.Println((coresholder[0])[21:])
		lscpuHOLDER.CPUs = (coresholder[0])[21:]
	} else {
		lscpuHOLDER.CPUs = string(cores)
	}

	//populating threads per core data
	checkThread := true
	thread, err := exec.Command("grep", "Thread", systemInfoLoc).Output()
	if err != nil {
		thread = []byte("N/A")
		checkThread = false
		log.Fatalf("%s: %s", err, "grabbing threads per core failed")
	}
	if checkThread {
		//fmt.Println(string(thread)[21 : len(string(thread))-1])
		lscpuHOLDER.ThreadsPerCore = string(thread)[21 : len(string(thread))-1]
	} else {
		lscpuHOLDER.ThreadsPerCore = string(thread)
	}

	//populating vendor ID
	checkVendID := true
	venID, err := exec.Command("grep", "Vendor", systemInfoLoc).Output()
	if err != nil {
		venID = []byte("N/A")
		checkVendID = false
		log.Fatalf("%s: %s", err, "grabbing vendorID")
	}
	if checkVendID {
		//fmt.Println(string(venID)[21 : len(string(venID))-1])
		lscpuHOLDER.VendorID = string(venID)[21 : len(string(venID))-1]
	} else {
		lscpuHOLDER.VendorID = string(venID)
	}

	//populating model name
	checkMN := true
	mn, err := exec.Command("grep", "Model name", systemInfoLoc).Output()
	if err != nil {
		mn = []byte("N/A")
		checkMN = false
		log.Fatalf("%s: %s", err, "grabbing CPU model name failed")
	}
	if checkMN {
		//fmt.Println(string(mn)[21 : len(string(mn))-1])
		lscpuHOLDER.ModelName = string(mn)[21 : len(string(mn))-1]
	} else {
		lscpuHOLDER.ModelName = string(mn)
	}

	//populating CPU MHz
	checkMHZ := true
	cpuMhz, err := exec.Command("grep", "CPU MHz", systemInfoLoc).Output()
	if err != nil {
		cpuMhz = []byte("N/A")
		checkMHZ = false
		log.Fatalf("%s: %s", err, "grabbing CPU MHZ failed")
	}
	if checkMHZ {
		//fmt.Println(string(cpuMhz)[21 : len(string(cpuMhz))-1])
		lscpuHOLDER.CPUMHz = string(cpuMhz)[21 : len(string(cpuMhz))-1]
	} else {
		lscpuHOLDER.CPUMHz = string(cpuMhz)
	}

	//populating max cpu mhz
	checkMAXM := true
	cpuMAXMhz, err := exec.Command("grep", "CPU max MHz", systemInfoLoc).Output()
	if err != nil {
		cpuMAXMhz = []byte("N/A")
		checkMAXM = false
		log.Fatalf("%s: %s", err, "grabbing CPU max MHz failed")
	}
	if checkMAXM {
		//fmt.Println(string(cpuMAXMhz)[21 : len(string(cpuMAXMhz))-1])
		lscpuHOLDER.CPUmaxMHz = string(cpuMAXMhz)[21 : len(string(cpuMAXMhz))-1]
	} else {
		lscpuHOLDER.CPUmaxMHz = string(cpuMAXMhz)
	}

	//populating min cpu mhz
	checkMINM := true
	cpuMINMhz, err := exec.Command("grep", "CPU min MHz", systemInfoLoc).Output()
	if err != nil {
		cpuMINMhz = []byte("N/A")
		checkMINM = false
		log.Fatalf("%s: %s", err, "grabbing CPU min MHz failed")
	}
	if checkMINM {
		//fmt.Println(string(cpuMINMhz)[21 : len(string(cpuMINMhz))-1])
		lscpuHOLDER.CPUminMHz = string(cpuMINMhz)[21 : len(string(cpuMINMhz))-1]
	} else {
		lscpuHOLDER.CPUminMHz = string(cpuMINMhz)
	}

	//populating virtualization variable
	checkVirt := true
	virt, err := exec.Command("grep", "Virtualiza", systemInfoLoc).Output()
	if err != nil {
		virt = []byte("N/A")
		checkVirt = false
		log.Fatalf("%s: %s", err, "grabbing Virtualization failed")
	}
	if checkVirt {
		//fmt.Println(string(virt)[21 : len(string(virt))-1])
		lscpuHOLDER.Virtualization = string(virt)[21 : len(string(virt))-1]
	} else {
		lscpuHOLDER.Virtualization = string(virt)
	}

	return lscpuHOLDER
}
