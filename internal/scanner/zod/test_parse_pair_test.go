//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestParsePair 테스트
package zod

import "testing"

func TestParsePair(t *testing.T) {
	root, src := parseTS(t, `const o = { name: z.string() };`)
	pairs := findAllByType(root, "pair")
	f := ParsePair(pairs[0], src)
	if f == nil || f.Name != "name" || f.Type != "string" {
		t.Fatalf("got %+v", f)
	}
}
