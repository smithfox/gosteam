package apath

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var gAppPath string

func InitAppPath() {
	gAppPath, _ = getAppPath()
}

func AppPath() string {
	return gAppPath
}

func getAppPath() (string, error) {
	execFilePath, e := exec.LookPath(os.Args[0])
	if e != nil {
		fmt.Println("fail to get the exe file path, %v", e)
		return "", e
	} else {
		//解决linux下相对路径的问题
		execFilePath, _ = filepath.Abs(execFilePath)
	}

	binDirPath := filepath.Dir(execFilePath)
	var appPath string
	if strings.HasSuffix(binDirPath, "bin") {
		appPath = filepath.Dir(binDirPath)
	} else if strings.HasSuffix(binDirPath, "src") { //just for developer
		appPath = filepath.Dir(binDirPath)
	} else {
		appPath = binDirPath
	}

	return appPath, nil
}
