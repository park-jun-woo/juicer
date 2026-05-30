//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestConvertPropertyToParam_ObjectNoTypeNoDefault 테스트
package fastify

import "testing"

func TestConvertPropertyToParam_ObjectNoTypeNoDefault(t *testing.T) {

	pairs, src := schemaPairs(t, `{ q: { description: "x" } }`)
	p := convertPropertyToParam(pairs[0], src)
	if p == nil || p.Type != "string" || p.Default != "" {
		t.Fatalf("unexpected param: %+v", p)
	}
}
