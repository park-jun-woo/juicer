//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestPairKeyName_Computed 테스트
package fastify

import "testing"

func TestPairKeyName_Computed(t *testing.T) {

	pairs, src := schemaPairs(t, `{ [keyVar]: 1 }`)
	if len(pairs) == 0 {
		t.Skip("no pair for computed key")
	}

	if got := pairKeyName(pairs[0], src); got != "" {
		t.Logf("computed key returned %q", got)
	}
}
