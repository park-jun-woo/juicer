//ff:func feature=ratchet type=session control=sequence
//ff:what TestPrintSkeleton_ExtractError 테스트
package sqls

import (
	"testing"
)

func TestPrintSkeleton_ExtractError(t *testing.T) {
	sess := &Session{
		RepoDir:    "/nonexistent",
		QueriesDir: ".",
		Methods: []MethodStatus{
			{ID: "Repo.Method", Status: "TODO"},
		},
	}
	// Should not panic, prints fallback
	printSkeleton(sess, 0)
}
