//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractZodValidatorsFromArgs_NonCallSkipped 테스트
package express

import "testing"

func TestExtractZodValidatorsFromArgs_NonCallSkipped(t *testing.T) {
	fi := mustParse(t, []byte(`r.post('/x', handler);`))
	nodes := argNodesOf(t, fi)
	if got := extractZodValidatorsFromArgs(nodes, fi.Src, 1); got != nil {
		t.Fatalf("got %+v", got)
	}
}
