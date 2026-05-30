//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestFilePrefixes_MemoAndCompose 테스트
package fastify

import "testing"

func TestFilePrefixes_MemoAndCompose(t *testing.T) {

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

	got2 := filePrefixes("/child.ts", graph, memo, visiting)
	if len(got2) != 1 || got2[0] != "/p/c" {
		t.Fatalf("memo: expected [/p/c], got %v", got2)
	}
}
