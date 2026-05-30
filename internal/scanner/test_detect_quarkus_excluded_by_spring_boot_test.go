//ff:func feature=scan type=test control=sequence
//ff:what TestDetectQuarkus_ExcludedBySpringBoot 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectQuarkus_ExcludedBySpringBoot(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "pom.xml"),
		[]byte("spring-boot quarkus-rest"), 0o644)
	if detectQuarkus(dir) {
		t.Fatal("expected false when spring-boot present")
	}
}
