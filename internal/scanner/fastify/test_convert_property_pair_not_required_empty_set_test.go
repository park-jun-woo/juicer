//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestConvertPropertyPair_NotRequiredEmptySet 테스트
package fastify

import "testing"

func TestConvertPropertyPair_NotRequiredEmptySet(t *testing.T) {
	pairs, src := schemaPairs(t, `{ title: { type: "string" } }`)
	f := convertPropertyPair(pairs[0], src, map[string]bool{})
	if f == nil || f.Name != "title" || f.Validate == "required" {
		t.Fatalf("unexpected field: %+v", f)
	}
}
