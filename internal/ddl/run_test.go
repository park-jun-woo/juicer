//ff:func feature=ddl type=command control=sequence
//ff:what TestRun_Basic 테스트
package ddl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_Basic(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "001_init.up.sql"), []byte("CREATE TABLE users (id INT);\n"), 0o644)
	out, err := Run(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "CREATE TABLE users") {
		t.Fatalf("unexpected: %q", out)
	}
}
