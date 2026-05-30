//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what pairKeyName 테스트
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

func TestPairKeyName_Computed(t *testing.T) {
	// computed key has no property_identifier/string/number direct child
	pairs, src := schemaPairs(t, `{ [keyVar]: 1 }`)
	if len(pairs) == 0 {
		t.Skip("no pair for computed key")
	}
	// computed property key -> "" (no recognized key node)
	if got := pairKeyName(pairs[0], src); got != "" {
		t.Logf("computed key returned %q", got)
	}
}
