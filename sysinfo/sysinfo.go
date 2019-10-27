package sysinfo

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

//SysInfo struct is goignt o be used later to read the file and be passed into a html template
type SysInfo struct {
	SystemUser             string
	SystemKernel           string
	SystemKernelRelease    string
	SystemKernelVersion    string
	SystemArch             string
	SystemProcessor        string
	SystemHardwarePlatform string
	SystemOS               string
}

//LSCPU holde the selected date from running the sysstat lscpu command
type LSCPU struct {
	Architecture   string
	CPUopmode      string
	CPUs           string
	ThreadsPerCore string
	VendorID       string
	ModelName      string
	CPUMHz         string
	CPUmaxMHz      string
	CPUminMHz      string
	Virtualization string
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
	arch, err := exec.Command("grep", "Arch", systemInfoLoc).Output()
	if err != nil {
		log.Fatalf("%s: %s", err, "grabbing Architechtire failed")
	}
	//fmt.Println(string(arch)[21 : len(string(arch))-1])
	lscpuHOLDER.Architecture = string(arch)[21 : len(string(arch))-1]

	//population cpu OpMode
	CPUopMOde, err := exec.Command("grep", "CPU op", systemInfoLoc).Output()
	if err != nil {
		log.Fatalf("%s: %s", err, "grabbing CPUopMODE failed")
	}
	//fmt.Println(string(CPUopMOde)[21 : len(string(CPUopMOde))-1])
	lscpuHOLDER.CPUopmode = string(CPUopMOde)[21 : len(string(CPUopMOde))-1]
	//population # of CPUS
	cores, err := exec.Command("grep", "CPU(s):", systemInfoLoc).Output()
	if err != nil {
		log.Fatalf("%s: %s", err, "grabbing core# failed")
	}
	//have to substring here because grep returns multiple values
	coresholder := strings.Split(string(cores), "\n")
	//fmt.Println((coresholder[0])[21:])
	lscpuHOLDER.CPUs = (coresholder[0])[21:]

	//populating threads per core data
	thread, err := exec.Command("grep", "Thread", systemInfoLoc).Output()
	if err != nil {
		log.Fatalf("%s: %s", err, "grabbing threads per core failed")
	}
	//fmt.Println(string(thread)[21 : len(string(thread))-1])
	lscpuHOLDER.ThreadsPerCore = string(thread)[21 : len(string(thread))-1]

	//populating vendor ID
	venID, err := exec.Command("grep", "Vendor", systemInfoLoc).Output()
	if err != nil {
		log.Fatalf("%s: %s", err, "grabbing vendorID")
	}
	//fmt.Println(string(venID)[21 : len(string(venID))-1])
	lscpuHOLDER.VendorID = string(venID)[21 : len(string(venID))-1]

	//populating model name
	mn, err := exec.Command("grep", "Model name", systemInfoLoc).Output()
	if err != nil {
		log.Fatalf("%s: %s", err, "grabbing CPU model name failed")
	}
	//fmt.Println(string(mn)[21 : len(string(mn))-1])
	lscpuHOLDER.ModelName = string(mn)[21 : len(string(mn))-1]

	//populating CPU MHz
	cpuMhz, err := exec.Command("grep", "CPU MHz", systemInfoLoc).Output()
	if err != nil {
		log.Fatalf("%s: %s", err, "grabbing CPU MHZ failed")
	}
	//fmt.Println(string(cpuMhz)[21 : len(string(cpuMhz))-1])
	lscpuHOLDER.CPUMHz = string(cpuMhz)[21 : len(string(cpuMhz))-1]

	//populating max cpu mhz
	cpuMAXMhz, err := exec.Command("grep", "CPU max MHz", systemInfoLoc).Output()
	if err != nil {
		log.Fatalf("%s: %s", err, "grabbing CPU max MHz failed")
	}
	//fmt.Println(string(cpuMAXMhz)[21 : len(string(cpuMAXMhz))-1])
	lscpuHOLDER.CPUmaxMHz = string(cpuMAXMhz)[21 : len(string(cpuMAXMhz))-1]

	//populating min cpu mhz
	cpuMINMhz, err := exec.Command("grep", "CPU min MHz", systemInfoLoc).Output()
	if err != nil {
		log.Fatalf("%s: %s", err, "grabbing CPU min MHz failed")
	}
	//fmt.Println(string(cpuMINMhz)[21 : len(string(cpuMINMhz))-1])
	lscpuHOLDER.CPUminMHz = string(cpuMINMhz)[21 : len(string(cpuMINMhz))-1]

	//populating virtualization variable
	virt, err := exec.Command("grep", "Virtualiza", systemInfoLoc).Output()
	if err != nil {
		log.Fatalf("%s: %s", err, "grabbing Virtualization failed")
	}
	//fmt.Println(string(virt)[21 : len(string(virt))-1])
	lscpuHOLDER.Virtualization = string(virt)[21 : len(string(virt))-1]

	return lscpuHOLDER
}

//ReadSysInfo will read the file created by CreateSystemInfoFile and populate the structure
func ReadSysInfo() SysInfo {
	//creates a logging file when error occurs
	f, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	//initialize SysInfo Struct
	systemInfoLoc := os.ExpandEnv("$HOME/systemvar.txt")
	file, err := os.Open(systemInfoLoc)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewReader(file)
	var txthold string
	counter := 0
	sysInfo := SysInfo{}
	for {
		counter++
		txthold, err = scanner.ReadString('\n')
		if err != nil {
			break
		}
		//adds the substring because txthold grabs the \n value and increases the length by 1 or makes a new line
		if counter == 1 {
			sysInfo.SystemUser = txthold[0 : len(txthold)-1]
		} else if counter == 2 {
			sysInfo.SystemKernel = txthold[0 : len(txthold)-1]
		} else if counter == 3 {
			sysInfo.SystemKernelRelease = txthold[0 : len(txthold)-1]
		} else if counter == 4 {
			sysInfo.SystemKernelVersion = txthold[0 : len(txthold)-1]
		} else if counter == 5 {
			sysInfo.SystemArch = txthold[0 : len(txthold)-1]
		} else if counter == 6 {
			sysInfo.SystemProcessor = txthold[0 : len(txthold)-1]
		} else if counter == 7 {
			sysInfo.SystemHardwarePlatform = txthold[0 : len(txthold)-1]
		} else if counter == 8 {
			sysInfo.SystemOS = txthold[0 : len(txthold)-1]
		}
	}
	if err != io.EOF {
		fmt.Printf("failed: %v\n", err)
	}
	return sysInfo
}

//CreateSystemInfoFile2 creates a file in the $HOME path with the important system variables
func CreateSystemInfoFile2() {
	f, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	/*
		only doign this section becase its would be able to be ported to other GOOS hopefully
	*/
	systemInfoLoc := os.ExpandEnv("$HOME/systemvar.txt")
	exec.Command("rm", systemInfoLoc).Run() //file removed to avoid redundancy when appending
	//file is opened here to write to the information. if doesnt exist it will be created
	file, err := os.OpenFile(systemInfoLoc, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//gets the system user
	getSystemUser, err := exec.Command("echo", os.Getenv("USER")).CombinedOutput()
	if err != nil {
		// Show error and output
		log.Fatalf("%s: %s", err, getSystemUser)
	}
	if _, err := file.WriteString(string(getSystemUser)); err != nil {
		log.Println(err)
	}

	// gets system kernel
	getSystemKernel, err := exec.Command("uname", "-s").CombinedOutput()
	if err != nil {
		// Show error and output
		log.Fatalf("%s: %s", err, getSystemKernel)
	}
	if _, err := file.WriteString(string(getSystemKernel)); err != nil {
		log.Println(err)
	}

	//gets system kernel release
	getSystemKernelRelease, err := exec.Command("uname", "-r").CombinedOutput()
	if err != nil {
		// Show error and output
		log.Fatalf("%s: %s", err, getSystemKernelRelease)
	}
	if _, err := file.WriteString(string(getSystemKernelRelease)); err != nil {
		log.Println(err)
	}
	//getSystemKernelVersion
	getSystemKernelVersion, err := exec.Command("uname", "-v").CombinedOutput()
	if err != nil {
		// Show error and output
		log.Fatalf("%s: %s", err, getSystemKernelVersion)
	}
	if _, err := file.WriteString(string(getSystemKernelVersion)); err != nil {
		log.Println(err)
	}
	//getSystemArch
	getSystemArch, err := exec.Command("uname", "--m").CombinedOutput()
	if err != nil {
		// Show error and output
		log.Fatalf("%s: %s", err, getSystemArch)
	}
	if _, err := file.WriteString(string(getSystemArch)); err != nil {
		log.Println(err)
	}
	//getSystemProcessor
	getSystemProcessor, err := exec.Command("uname", "-p").CombinedOutput()
	if err != nil {
		// Show error and output
		log.Fatalf("%s: %s", err, getSystemProcessor)
	}
	if _, err := file.WriteString(string(getSystemProcessor)); err != nil {
		log.Println(err)
	}
	//getSystemHardwarePlatform
	getSystemHardwarePlatform, err := exec.Command("uname", "-i").CombinedOutput()
	if err != nil {
		// Show error and output
		log.Fatalf("%s: %s", err, getSystemHardwarePlatform)
	}
	if _, err := file.WriteString(string(getSystemHardwarePlatform)); err != nil {
		log.Println(err)
	}
	//getSystemOS
	getSystemOS, err := exec.Command("uname", "-o").CombinedOutput()
	if err != nil {
		// Show error and output
		log.Fatalf("%s: %s", err, getSystemOS)
	}
	if _, err := file.WriteString(string(getSystemOS)); err != nil {
		log.Println(err)
	}
}

//CreateSystemInfoFile creates a file in the $HOME path with the important system variables
func CreateSystemInfoFile() {
	//inputs the users of the system information into the file
	//also recreates the file each time this is run
	getSystemUser := exec.Command("bash", "-c", "echo "+os.Getenv("USER")+"> "+os.Getenv("HOME")+"/systemvar2.txt")
	getSystemUser.Run()
	//all other commands work the same but append to the file
	//commands are self explanatory
	getSystemKernel := exec.Command("bash", "-c", "uname -s >> "+os.Getenv("HOME")+"/systemvar.txt")
	getSystemKernel.Run()
	getSystemKernelRelease := exec.Command("bash", "-c", "uname -r >> "+os.Getenv("HOME")+"/systemvar.txt")
	getSystemKernelRelease.Run()
	getSystemKernelVersion := exec.Command("bash", "-c", "uname -v >> "+os.Getenv("HOME")+"/systemvar.txt")
	getSystemKernelVersion.Run()
	getSystemArch := exec.Command("bash", "-c", "uname --m >> "+os.Getenv("HOME")+"/systemvar.txt")
	getSystemArch.Run()
	getSystemProcessor := exec.Command("bash", "-c", "uname -p >> "+os.Getenv("HOME")+"/systemvar.txt")
	getSystemProcessor.Run()
	getSystemHardwarePlatform := exec.Command("bash", "-c", "uname -i >> "+os.Getenv("HOME")+"/systemvar.txt")
	getSystemHardwarePlatform.Run()
	getSystemOS := exec.Command("bash", "-c", "uname -o >> "+os.Getenv("HOME")+"/systemvar.txt")
	getSystemOS.Run()
}
