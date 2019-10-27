package sysinfo

import (
	"fmt"
	"testing"
)

func TestReadSysInfo(t *testing.T) {
	sysInfoTester := SysInfo{"gogopowerrangers", "Linux", "5.0.0-32-generic",
		"#34~18.04.2-Ubuntu SMP Thu Oct 10 10:36:02 UTC 2019", "x86_64", "x86_64", "x86_64", "x86_64"}
	sysInfoReadFromFile := ReadSysInfo()
	fmt.Println(sysInfoReadFromFile.SystemUser, len(sysInfoReadFromFile.SystemUser))
	if sysInfoTester.SystemUser == sysInfoReadFromFile.SystemUser {
		println("Read File Works")
	} else {
		println("TestReadSysInfo Fail")
	}
}

//^if file was created then we could read from it is it redudndant to create a test function for this?
func TestCreateSystemInfoFile2(t *testing.T) {
	fmt.Println("TestCreateSystemInfoFileWorks")
}
