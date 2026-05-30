//ff:func feature=scan type=test control=sequence
//ff:what TestDetectQuarkus_Hit 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectQuarkus_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "pom.xml"),
		[]byte("<dependency>quarkus-resteasy</dependency>"), 0o644)
	if !detectQuarkus(dir) {
		t.Fatal("expected true")
	}
}
