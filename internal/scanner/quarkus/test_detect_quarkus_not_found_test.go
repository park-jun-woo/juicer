//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestDetectQuarkus_NotFound -- Quarkus 의존성이 없으면 감지하지 않는 테스트
package quarkus

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectQuarkus_NotFound(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "pom.xml", `<?xml version="1.0"?>
<project>
  <dependencies>
    <dependency>
      <groupId>io.dropwizard</groupId>
      <artifactId>dropwizard-core</artifactId>
    </dependency>
  </dependencies>
</project>`)

	fw := scanner.DetectFramework(dir)
	if fw == "quarkus" {
		t.Errorf("should not detect quarkus, got %s", fw)
	}
}
