//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what buildMountGraph: 인라인/빈 FilePath 제외하고 대상별 그룹핑을 검증
package fastify

import "testing"

func TestBuildMountGraph(t *testing.T) {
	mounts := []pluginMount{
		{FilePath: "/a.ts", SourceFile: "/app.ts", Prefix: "/x"},
		{FilePath: "/a.ts", SourceFile: "/b.ts", Prefix: "/y"},
		{Inline: true, Prefix: "/skip"},
		{FilePath: "", Prefix: "/skip2"},
	}
	g := buildMountGraph(mounts)
	if len(g) != 1 || len(g["/a.ts"]) != 2 {
		t.Fatalf("want 1 key with 2 mounts, got %v", g)
	}
}
