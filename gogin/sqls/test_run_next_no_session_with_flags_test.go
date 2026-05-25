//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunNext_NoSession_WithFlags 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunNext_NoSession_WithFlags(t *testing.T) {
	dir := setupSessionDir(t)
	repoDir := filepath.Join(dir, "repo")
	queriesDir := filepath.Join(dir, "queries")
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)

	err := RunNext(repoDir, queriesDir)
	if err != nil {
		t.Fatalf("RunNext() error: %v", err)
	}

	if !SessionExists() {
		t.Error("expected session to be created")
	}
}
