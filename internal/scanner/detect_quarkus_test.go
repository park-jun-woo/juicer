//ff:func feature=scan type=test control=sequence
//ff:what detectQuarkus — quarkus-rest/resteasy 의존 감지(spring-boot 배제) 분기를 검증
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

func TestDetectQuarkus_RestVariant(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "build.gradle"),
		[]byte("implementation 'io.quarkus:quarkus-rest'"), 0o644)
	if !detectQuarkus(dir) {
		t.Fatal("expected true for quarkus-rest")
	}
}

func TestDetectQuarkus_ExcludedBySpringBoot(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "pom.xml"),
		[]byte("spring-boot quarkus-rest"), 0o644)
	if detectQuarkus(dir) {
		t.Fatal("expected false when spring-boot present")
	}
}

func TestDetectQuarkus_NoFiles(t *testing.T) {
	if detectQuarkus(t.TempDir()) {
		t.Fatal("expected false when no build files")
	}
}
