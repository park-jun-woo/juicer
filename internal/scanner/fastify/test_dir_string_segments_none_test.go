//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestDirStringSegments_None 테스트
package fastify

import "testing"

func TestDirStringSegments_None(t *testing.T) {
	fi := mustParse(t, []byte("const x = foo(a, b);\n"))
	calls := findAllByType(fi.Root, "call_expression")[0]
	if segs := dirStringSegments(calls, fi.Src); len(segs) != 0 {
		t.Fatalf("expected no segments, got %v", segs)
	}
}
