//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestDirStringSegments 테스트
package fastify

import "testing"

func TestDirStringSegments(t *testing.T) {
	fi := mustParse(t, []byte(`const x = join(__dirname, "routes", "api");`+"\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no call")
	}
	segs := dirStringSegments(calls[0], fi.Src)
	if len(segs) != 2 || segs[0] != "routes" || segs[1] != "api" {
		t.Fatalf("expected [routes api], got %v", segs)
	}
}
