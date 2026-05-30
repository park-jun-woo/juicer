//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestConvertPropertyToParam_BooleanType 테스트
package fastify

import "testing"

func TestConvertPropertyToParam_BooleanType(t *testing.T) {
	pairs, src := schemaPairs(t, `{ active: { type: "boolean" } }`)
	p := convertPropertyToParam(pairs[0], src)
	if p == nil || p.Name != "active" {
		t.Fatalf("unexpected param: %+v", p)
	}
}
