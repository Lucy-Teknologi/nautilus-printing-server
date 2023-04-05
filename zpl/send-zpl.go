package zpl

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
)

var trimming = regexp.MustCompile(`[\n\t]+`)

func ExecuteZpl(zpl_commands string) error {
	// check if system is running windows or linux
	// if windows, use the windows bat script
	// if linux, use the linux script
	tmpFile := filepath.Join(os.TempDir(), "tmp.zpl")
	refined_commands := trimming.ReplaceAllString(zpl_commands, "")
	if err := os.WriteFile(tmpFile, []byte(refined_commands), 0666); err != nil {
		return err
	}

	script := "./bin/print.sh"
	if runtime.GOOS == "windows" {
		script = "./bin/print.bat"
	}
	bytes, err := exec.Command(script).Output()
	if err != nil {
		fmt.Println(bytes)
		return err
	}
	fmt.Println(string(bytes))
	return nil

}
