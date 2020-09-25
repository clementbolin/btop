package scan

import (
	"io/ioutil"

	error "github.com/ClementBolin/topGo/modules/gitStat/err"
)

// ScanUniqueFolderGit : Scan folder not recursive
func ScanUniqueFolderGit(path string, email string) []string {
	// Read File
	file, err := ioutil.ReadDir(path)
	error.MangeErrExit(err)

	// String path
	pathReturn := make([]string, 1)

	for _, f := range file {
		if (f.Name() == ".git") {
			pathReturn[0] = path + "/" + f.Name()
		}
	}
	return pathReturn
}
