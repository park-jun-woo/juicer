//ff:func feature=hurl type=parse control=sequence
//ff:what hurl --test 실행 및 결과 반환
package hurls

import (
	"bytes"
	"fmt"
	"os/exec"
)

// runHurlTest executes hurl --test on a file and returns (passed, stderr).
func runHurlTest(testFile, host string) (bool, string) {
	cmd := exec.Command("hurl", "--test", testFile, "--variable", "host="+host)
	var stderrBuf bytes.Buffer
	var stdoutBuf bytes.Buffer
	cmd.Stderr = &stderrBuf
	cmd.Stdout = &stdoutBuf

	err := cmd.Run()
	if err != nil {
		// Check if hurl is not found
		if _, ok := err.(*exec.Error); ok {
			return false, fmt.Sprintf("hurl not found in PATH: %v", err)
		}
		return false, stderrBuf.String()
	}
	return true, ""
}
