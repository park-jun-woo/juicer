//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestFilePrefixesCycle 테스트
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
