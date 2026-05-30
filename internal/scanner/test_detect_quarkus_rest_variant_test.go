//ff:func feature=scan type=test control=sequence
//ff:what TestDetectQuarkus_RestVariant 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectQuarkus_RestVariant(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "build.gradle"),
		[]byte("implementation 'io.quarkus:quarkus-rest'"), 0o644)
	if !detectQuarkus(dir) {
		t.Fatal("expected true for quarkus-rest")
	}
}
