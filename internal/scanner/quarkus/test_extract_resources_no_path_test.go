//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractResources_NoPath 테스트
package quarkus

import "testing"

func TestExtractResources_NoPath(t *testing.T) {
	fi := qFileInfo(t, `public class PlainClass {}`)
	if r := extractResources(fi); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
