//ff:func feature=sql type=test control=sequence
//ff:what TestParseFile_Valid 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_Valid(t *testing.T) {
	dir := t.TempDir()
	src := `package repo

type UserRepo struct{}

func (r *UserRepo) GetUser() {}
`
	f := filepath.Join(dir, "user_repo.go")
	os.WriteFile(f, []byte(src), 0o644)
	_, err := parseFile(f)
	if err != nil {
		t.Fatal(err)
	}
}
