//ff:func feature=sql type=parse control=sequence
//ff:what TestPrintSkeleton 테스트
package sqls

import (
	"os"
	"testing"
)

func TestPrintSkeleton_NoMatchCov(t *testing.T) {
	dir := t.TempDir()
	os.MkdirAll(dir+"/repo", 0o755)
	sess := &Session{
		RepoDir:    dir + "/repo",
		QueriesDir: dir + "/queries",
		Methods:    []MethodStatus{{ID: "Repo.Method", Status: "TODO"}},
	}
	printSkeleton(sess, 0, nil)
}

