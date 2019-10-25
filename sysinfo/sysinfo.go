package sysinfo

import (
	"log"
	"os"
	"os/exec"
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

//CreateSystemInfoFile creates a file in the $HOME path with the important system variables
func CreateSystemInfoFile() {
	//inputs the users of the system information into the file
	//also recreates the file each time this is run
	getSystemUser := exec.Command("bash", "-c", "echo "+os.Getenv("USER")+"> "+os.Getenv("HOME")+"/systemvar.txt")
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

//CreateSystemInfoFile2 creates a file in the $HOME path with the important system variables
func CreateSystemInfoFile2() {
	/*

		WRITE TO FILE WORKS NOW AND APPENDSwork ont his further try to expand on this
		only doign this section becase its would be able to be ported to other GOOS

	*/
	systemInfoLoc := os.ExpandEnv("$HOME/systemvar3.txt")
	//file is opened here to write to the information. if doesnt exist it will be created
	file, err := os.OpenFile(systemInfoLoc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Run uname command and get both stdout and stderr
	getSystemKernel, err := exec.Command("uname", "-s").CombinedOutput()
	if err != nil {
		// Show error and output
		log.Fatalf("%s: %s", err, getSystemKernel)
	}
	if _, err := file.WriteString(string(getSystemKernel)); err != nil {
		log.Println(err)
	}
}
