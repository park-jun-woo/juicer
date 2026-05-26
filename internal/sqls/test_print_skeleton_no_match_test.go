//ff:func feature=ratchet type=session control=sequence
//ff:what TestPrintSkeleton_NoMatch 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPrintSkeleton_NoMatch(t *testing.T) {
	dir := setupSessionDir(t)
	repoDir := filepath.Join(dir, "repo")
	os.MkdirAll(repoDir, 0o755)

	sess := &Session{
		RepoDir:    repoDir,
		QueriesDir: ".",
		Methods: []MethodStatus{
			{ID: "NonExistent.Method", Status: "TODO"},
		},
	}
	// Should print fallback since no matching skeleton
	result, _ := Extract(repoDir)
	var methods []MethodSkeleton
	if result != nil {
		methods = result.Methods
	}
	printSkeleton(sess, 0, methods)
}
