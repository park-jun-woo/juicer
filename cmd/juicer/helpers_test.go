//ff:func feature=scan type=command control=sequence
//ff:what setupMinimalGoProject 함수
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func setupMinimalGoProject(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	goMod := `module example.com/test

go 1.21
`
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte(goMod), 0o644)
	mainGo := `package main

func main() {}
`
	os.WriteFile(filepath.Join(dir, "main.go"), []byte(mainGo), 0o644)
	return dir
}

