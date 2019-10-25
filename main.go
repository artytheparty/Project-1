package main

import (
	"github.com/artytheparty/project-1/sysinfo"
)

const systemInfoLoc string = "$HOME/systemvar.txt"
const systemProcInfoLoc string = "$HOME/processtable.txt"

func main() {
	sysinfo.CreateSystemInfoFile2()
}
