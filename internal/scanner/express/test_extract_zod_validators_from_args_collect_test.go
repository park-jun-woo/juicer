//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractZodValidatorsFromArgs_Collect 테스트
package express

import "testing"

func TestExtractZodValidatorsFromArgs_Collect(t *testing.T) {

	fi := mustParse(t, []byte(`r.post('/x', validateRequest({ body: s }), handler);`))
	nodes := argNodesOf(t, fi)
	got := extractZodValidatorsFromArgs(nodes, fi.Src, 1)
	if len(got) != 1 || got[0].Target != "json" {
		t.Fatalf("got %+v", got)
	}
}
