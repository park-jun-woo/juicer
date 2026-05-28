//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestDetectQuarkus_PomXml -- pom.xml에서 quarkus-resteasy 감지 테스트
package quarkus

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectQuarkus_PomXml(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "pom.xml", `<?xml version="1.0"?>
<project>
  <dependencies>
    <dependency>
      <groupId>io.quarkus</groupId>
      <artifactId>quarkus-resteasy-reactive</artifactId>
    </dependency>
  </dependencies>
</project>`)

	fw := scanner.DetectFramework(dir)
	if fw != "quarkus" {
		t.Errorf("expected quarkus, got %s", fw)
	}
}
