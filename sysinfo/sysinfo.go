package sysinfo

import (
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
