package process

import (
	"log"

	ps "github.com/mitchellh/go-ps"
)

// ListProcessUnix : display your all process running on your system
func ListProcessUnix() {
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
