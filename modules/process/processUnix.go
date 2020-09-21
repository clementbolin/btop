package process

import (
	"log"

	ps "github.com/mitchellh/go-ps"
)

// ListProcessUnix : display your all process running on your system
func ListProcessUnix() []UnixProcess {
	var unixListProcess []UnixProcess

	listProcess, err := ps.Processes()
	if err != nil {
		log.Fatalln(err)
	}
	for item := range listProcess {
		var process ps.Process
		var myProcess UnixProcess
		process = listProcess[item]
		myProcess.SetName(process.Executable())
		myProcess.SetPid(process.Pid())
		myProcess.SetPpid(process.PPid())
		unixListProcess = append(unixListProcess, myProcess)
	}
	return unixListProcess
}
