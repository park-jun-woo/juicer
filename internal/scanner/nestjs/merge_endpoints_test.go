//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestMergeEndpoints_NoInherited 테스트
package nestjs

import "testing"

func TestMergeEndpoints_NoInherited(t *testing.T) {
	direct := []endpointInfo{{handler: "findAll", method: "GET"}}
	merged := mergeEndpoints(nil, direct)
	if len(merged) != 1 {
		t.Fatalf("expected 1, got %d", len(merged))
	}
}
