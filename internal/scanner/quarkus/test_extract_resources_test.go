//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractResources 테스트
package quarkus

import "testing"

func TestExtractResources(t *testing.T) {
	fi := qFileInfo(t, sampleResource)
	resources := extractResources(fi)
	if len(resources) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(resources))
	}
}
