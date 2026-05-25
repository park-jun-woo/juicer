//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunSqlcGenerate 테스트
package sqls

import (
	"testing"
)

func TestRunSqlcGenerate(t *testing.T) {
	// Test that it runs without panic. sqlc is installed
	// but may fail if no sqlc.yaml exists in cwd
	passed, stderr := runSqlcGenerate()
	if passed {
		t.Log("sqlc generate passed")
	} else {
		t.Logf("sqlc generate failed (expected without sqlc.yaml): %s", stderr)
	}
}
