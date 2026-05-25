//ff:func feature=hurl type=parse control=sequence
//ff:what TestFindTestFile_Found 테스트
package hurls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindTestFile_Found(t *testing.T) {
	dir := t.TempDir()
	content := "GET {{host}}/api/health\nHTTP 200\n"
	os.WriteFile(filepath.Join(dir, "health.hurl"), []byte(content), 0o644)
	got := findTestFile(dir, "GET", "/api/health")
	if got == "" {
		t.Fatal("expected to find test file")
	}
}
