//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what convertPropertyToParam 테스트
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

func TestConvertPropertyToParam_ObjectNoTypeNoDefault(t *testing.T) {
	// object value without type/default -> defaults to string, no default
	pairs, src := schemaPairs(t, `{ q: { description: "x" } }`)
	p := convertPropertyToParam(pairs[0], src)
	if p == nil || p.Type != "string" || p.Default != "" {
		t.Fatalf("unexpected param: %+v", p)
	}
}

func TestConvertPropertyToParam_BooleanType(t *testing.T) {
	pairs, src := schemaPairs(t, `{ active: { type: "boolean" } }`)
	p := convertPropertyToParam(pairs[0], src)
	if p == nil || p.Name != "active" {
		t.Fatalf("unexpected param: %+v", p)
	}
}

func TestConvertPropertyToParam_NonObjectValue(t *testing.T) {
	// value is a string literal, not an object -> stays type "string"
	pairs, src := schemaPairs(t, `{ flag: "true" }`)
	p := convertPropertyToParam(pairs[0], src)
	if p == nil || p.Name != "flag" || p.Type != "string" {
		t.Fatalf("unexpected param: %+v", p)
	}
}
