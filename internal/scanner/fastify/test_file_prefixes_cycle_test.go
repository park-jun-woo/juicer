//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what filePrefixes: 사이클이 있어도 무한 루프 없이 종료함을 검증
package fastify

import "testing"

func TestFilePrefixesCycle(t *testing.T) {
	mounts := []pluginMount{
		{FilePath: "/a.ts", SourceFile: "/b.ts", Prefix: "/a"},
		{FilePath: "/b.ts", SourceFile: "/a.ts", Prefix: "/b"},
	}
	graph := buildMountGraph(mounts)
	memo := make(map[string][]string)
	visiting := make(map[string]bool)
	got := filePrefixes("/a.ts", graph, memo, visiting)
	if got == nil {
		t.Fatal("expected non-nil result for cyclic graph")
	}
}

func TestFilePrefixes_MemoAndCompose(t *testing.T) {
	// /child mounted under /parent with prefix "/c"; /parent is a root.
	mounts := []pluginMount{
		{FilePath: "/child.ts", SourceFile: "/parent.ts", Prefix: "/c"},
		{FilePath: "/parent.ts", SourceFile: "/root.ts", Prefix: "/p"},
	}
	graph := buildMountGraph(mounts)
	memo := make(map[string][]string)
	visiting := make(map[string]bool)

	got := filePrefixes("/child.ts", graph, memo, visiting)
	if len(got) != 1 || got[0] != "/p/c" {
		t.Fatalf("expected [/p/c], got %v", got)
	}
	// second call hits the memo cache and returns the same result
	got2 := filePrefixes("/child.ts", graph, memo, visiting)
	if len(got2) != 1 || got2[0] != "/p/c" {
		t.Fatalf("memo: expected [/p/c], got %v", got2)
	}
}
