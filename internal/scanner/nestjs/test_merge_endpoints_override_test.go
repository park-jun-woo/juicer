//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestMergeEndpoints_Override 테스트
package nestjs

import "testing"

func TestMergeEndpoints_Override(t *testing.T) {
	inherited := []endpointInfo{
		{handler: "findAll", method: "GET", path: ""},
		{handler: "create", method: "POST", path: ""},
	}
	direct := []endpointInfo{
		{handler: "findAll", method: "GET", path: "custom"},
	}
	merged := mergeEndpoints(inherited, direct)
	if len(merged) != 2 {
		t.Fatalf("expected 2, got %d", len(merged))
	}
	for _, ep := range merged {
		if ep.handler == "findAll" && ep.path != "custom" {
			t.Fatalf("expected override path 'custom', got %q", ep.path)
		}
	}
}
