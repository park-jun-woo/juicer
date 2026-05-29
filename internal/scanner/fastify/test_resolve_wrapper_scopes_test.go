//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what resolveWrapperScopes: 인라인+prefix 마운트만 SourceFile별로 모음을 검증
package fastify

import "testing"

func TestResolveWrapperScopes(t *testing.T) {
	mounts := []pluginMount{
		{Inline: true, Prefix: "/w", SourceFile: "/app.ts", WrapperStart: 5, WrapperEnd: 50},
		{Inline: true, Prefix: "", SourceFile: "/app.ts"},
		{Inline: false, Prefix: "/x", SourceFile: "/app.ts"},
	}
	scopes := resolveWrapperScopes(mounts)
	if len(scopes["/app.ts"]) != 1 {
		t.Fatalf("want 1 wrapper scope, got %v", scopes["/app.ts"])
	}
	ws := scopes["/app.ts"][0]
	if ws.Prefix != "/w" || ws.Start != 5 || ws.End != 50 {
		t.Fatalf("unexpected scope %v", ws)
	}
}
