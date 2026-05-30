//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectAutoloadRequireNames_None 테스트
package fastify

import "testing"

func TestCollectAutoloadRequireNames_None(t *testing.T) {
	fi := mustParse(t, []byte("const a = 1;\n"))
	names := map[string]bool{}
	collectAutoloadRequireNames(fi, names)
	if len(names) != 0 {
		t.Fatalf("expected no names, got %v", names)
	}
}
