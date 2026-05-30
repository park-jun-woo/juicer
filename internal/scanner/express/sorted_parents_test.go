//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what sortedParents: file 우선, 동일 file이면 varName 기준 정렬
package express

import "testing"

func TestSortedParents(t *testing.T) {
	g := newTestGraph()
	child := routerKey{file: "z.ts", varName: "c"}
	// edges from various parents (different files, and same file different vars)
	graphAddEdge(g, routerKey{file: "b.ts", varName: "x"}, child, "/1")
	graphAddEdge(g, routerKey{file: "a.ts", varName: "y"}, child, "/2")
	graphAddEdge(g, routerKey{file: "a.ts", varName: "x"}, child, "/3")

	parents := sortedParents(g)
	want := []routerKey{
		{file: "a.ts", varName: "x"},
		{file: "a.ts", varName: "y"},
		{file: "b.ts", varName: "x"},
	}
	if len(parents) != len(want) {
		t.Fatalf("got %v", parents)
	}
	for i := range want {
		if parents[i] != want[i] {
			t.Fatalf("at %d got %+v want %+v", i, parents[i], want[i])
		}
	}
}
