package process

import (
	"log"

	ps "github.com/mitchellh/go-ps"
)

// ListAllProcess : display your all process running on your system
func ListAllProcess() {
	listProcess, err := ps.Processes()
	if err != nil {
		log.Fatalln(err)
	}
	for item := range listProcess {
		var process ps.Process
		process = listProcess[item]
		log.Println("exec:", process.Executable(), " ppid:", process.PPid(), "pid:", process.Pid());
	}
}
