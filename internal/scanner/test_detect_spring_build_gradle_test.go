//ff:func feature=scan type=test control=sequence
//ff:what TestDetectSpring_BuildGradle — build.gradle 에서 spring-boot-starter-web 감지
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectSpring_BuildGradle(t *testing.T) {
	dir := t.TempDir()
	gradle := `plugins {
    id 'org.springframework.boot' version '3.2.0'
}
dependencies {
    implementation 'org.springframework.boot:spring-boot-starter-web'
}`
	os.WriteFile(filepath.Join(dir, "build.gradle"), []byte(gradle), 0o644)
	if !detectSpring(dir) {
		t.Error("expected true for build.gradle with spring-boot-starter-web")
	}
}
