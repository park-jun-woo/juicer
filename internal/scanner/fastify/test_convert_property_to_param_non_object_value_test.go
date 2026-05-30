//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestConvertPropertyToParam_NonObjectValue 테스트
package fastify

import "testing"

func TestConvertPropertyToParam_NonObjectValue(t *testing.T) {

	pairs, src := schemaPairs(t, `{ flag: "true" }`)
	p := convertPropertyToParam(pairs[0], src)
	if p == nil || p.Name != "flag" || p.Type != "string" {
		t.Fatalf("unexpected param: %+v", p)
	}
}
