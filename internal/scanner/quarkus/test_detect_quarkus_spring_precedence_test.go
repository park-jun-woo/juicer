//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestDetectQuarkus_SpringBootNotQuarkus -- spring-boot가 quarkus보다 우선하는지 테스트
package quarkus

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectQuarkus_SpringBootNotQuarkus(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "pom.xml", `<?xml version="1.0"?>
<project>
  <dependencies>
    <dependency>
      <groupId>org.springframework.boot</groupId>
      <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    <dependency>
      <groupId>io.quarkus</groupId>
      <artifactId>quarkus-resteasy</artifactId>
    </dependency>
  </dependencies>
</project>`)

	fw := scanner.DetectFramework(dir)
	if fw != "spring" {
		t.Errorf("expected spring (spring-boot takes precedence), got %s", fw)
	}
}
