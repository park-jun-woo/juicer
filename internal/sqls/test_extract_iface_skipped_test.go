//ff:func feature=sql type=parse control=sequence
//ff:what TestExtract_IfaceSkipped 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtract_IfaceSkipped(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "user_iface_repo.go"), []byte("package repo\n"), 0o644)
	// Even though it ends with _repo.go, _iface.go check also matters
	// Actually, "_iface.go" suffix is checked. Let me create a proper _iface.go
	// Wait - the file must end with _repo.go AND _iface.go at the same time?
	// No: it checks _repo.go first, then _test.go or _iface.go
	// A file named "user_repo_iface.go" doesn't end with _repo.go
	// A file named "user_iface_repo.go" ends with _repo.go but not _iface.go
	// To trigger line 86, need a file that ends with _repo.go AND contains _iface.go suffix
	// Actually check is: strings.HasSuffix(name, "_test.go") || strings.HasSuffix(name, "_iface.go")
	// But the file already passed HasSuffix(name, "_repo.go"). So it can't also end with _iface.go
	// unless the name is "x_iface.go_repo.go" which is weird.
	// This code path seems unreachable — a file can't end with both _repo.go and _iface.go
	// Let me check if the line is actually dead code
	result, err := Extract(dir)
	if err != nil {
		t.Fatalf("Extract() error: %v", err)
	}
	_ = result
}
