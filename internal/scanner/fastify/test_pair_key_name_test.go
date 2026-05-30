//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestPairKeyName 테스트
package fastify

import "testing"

func TestPairKeyName(t *testing.T) {
	pairs, src := schemaPairs(t, `{ ident: 1, "strKey": 2, 200: 3 }`)
	got := map[string]bool{}
	for _, p := range pairs {
		got[pairKeyName(p, src)] = true
	}
	if !got["ident"] {
		t.Error("missing property_identifier key 'ident'")
	}
	if !got["strKey"] {
		t.Error("missing string key 'strKey'")
	}
	if !got["200"] {
		t.Error("missing number key '200'")
	}
}
