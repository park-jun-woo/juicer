//ff:func feature=scan type=test control=sequence
//ff:what TestDetectSpring_PomXml — pom.xml 에서 spring-boot-starter-web 감지
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectSpring_PomXml(t *testing.T) {
	dir := t.TempDir()
	pom := `<?xml version="1.0" encoding="UTF-8"?>
<project>
  <dependencies>
    <dependency>
      <groupId>org.springframework.boot</groupId>
      <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
  </dependencies>
</project>`
	os.WriteFile(filepath.Join(dir, "pom.xml"), []byte(pom), 0o644)
	if !detectSpring(dir) {
		t.Error("expected true for pom.xml with spring-boot-starter-web")
	}
}
