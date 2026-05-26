//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestMergeEndpoints_InheritedOnly 테스트
package nestjs

import "testing"

func TestMergeEndpoints_InheritedOnly(t *testing.T) {
	inherited := []endpointInfo{
		{handler: "findAll", method: "GET"},
		{handler: "create", method: "POST"},
	}
	merged := mergeEndpoints(inherited, nil)
	if len(merged) != 2 {
		t.Fatalf("expected 2, got %d", len(merged))
	}
}
