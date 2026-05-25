//ff:func feature=sql type=parse control=sequence
//ff:what cwd에서 "sqlc generate" 실행, 결과(통과 여부, stderr) 반환
package sqls

import (
	"bytes"
	"os/exec"
	"strings"
)

// runSqlcGenerate executes "sqlc generate" in cwd.
// Returns (passed, stderr).
//
func runSqlcGenerate() (bool, string) {
	_, err := exec.LookPath("sqlc")
	if err != nil {
		return false, "sqlc not found. Install: https://docs.sqlc.dev/en/latest/overview/install.html"
	}

	var stderrBuf bytes.Buffer
	cmd := exec.Command("sqlc", "generate")
	cmd.Stderr = &stderrBuf
	err = cmd.Run()
	if err != nil {
		return false, strings.TrimSpace(stderrBuf.String())
	}
	return true, ""
}

