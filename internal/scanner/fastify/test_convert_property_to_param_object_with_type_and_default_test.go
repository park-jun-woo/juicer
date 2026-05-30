//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestConvertPropertyToParam_ObjectWithTypeAndDefault 테스트
package fastify

import "testing"

func TestConvertPropertyToParam_ObjectWithTypeAndDefault(t *testing.T) {
	pairs, src := schemaPairs(t, `{ limit: { type: "integer", default: "10" } }`)
	p := convertPropertyToParam(pairs[0], src)
	if p == nil || p.Name != "limit" {
		t.Fatalf("unexpected param: %+v", p)
	}
	if p.Type != "integer" {
		t.Errorf("type = %q, want integer", p.Type)
	}
	if p.Default != "10" {
		t.Errorf("default = %q, want 10", p.Default)
	}
}
