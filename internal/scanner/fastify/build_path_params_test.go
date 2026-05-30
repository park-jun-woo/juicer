//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what buildPathParams 테스트
package fastify

import "testing"

func TestBuildPathParams(t *testing.T) {
	got := buildPathParams([]string{"id", "slug"})
	if len(got) != 2 {
		t.Fatalf("expected 2 params, got %d", len(got))
	}
	if got[0].Name != "id" || got[0].Type != "string" {
		t.Errorf("param 0 = %+v", got[0])
	}
	if got[1].Name != "slug" || got[1].Type != "string" {
		t.Errorf("param 1 = %+v", got[1])
	}

	// empty input -> nil/empty
	if len(buildPathParams(nil)) != 0 {
		t.Error("expected empty for nil input")
	}
}
