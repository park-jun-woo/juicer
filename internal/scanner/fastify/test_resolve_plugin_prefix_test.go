//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what resolvePluginPrefixes: 다중 등록 별도 prefix + transitive 합성을 검증
package fastify

import "testing"

func TestResolvePluginPrefixes(t *testing.T) {
	mounts := []pluginMount{
		{FilePath: "/pub.ts", Prefix: "/public/v1", SourceFile: "/app.ts"},
		{FilePath: "/users.ts", Prefix: "/users", SourceFile: "/pub.ts"},
		{FilePath: "/boot.ts", Prefix: "/boot", SourceFile: "/app.ts"},
		{FilePath: "/boot.ts", Prefix: "/new_boot", SourceFile: "/app.ts"},
	}
	m := resolvePluginPrefixes(mounts)

	if got := m["/users.ts"]; len(got) != 1 || got[0] != "/public/v1/users" {
		t.Fatalf("transitive: want [/public/v1/users], got %v", got)
	}
	boot := m["/boot.ts"]
	if len(boot) != 2 || !containsString(boot, "/boot") || !containsString(boot, "/new_boot") {
		t.Fatalf("multi-register: want [/boot /new_boot], got %v", boot)
	}
}
