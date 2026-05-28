//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestDetectQuarkus_BuildGradle -- build.gradle에서 quarkus-rest 감지 테스트
package quarkus

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectQuarkus_BuildGradle(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "build.gradle", `
plugins {
    id 'io.quarkus'
}
dependencies {
    implementation 'io.quarkus:quarkus-rest'
}
`)

	fw := scanner.DetectFramework(dir)
	if fw != "quarkus" {
		t.Errorf("expected quarkus, got %s", fw)
	}
}
