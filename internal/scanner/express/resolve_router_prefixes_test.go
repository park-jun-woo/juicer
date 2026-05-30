//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveRouterPrefixes: 마운트 그래프 → 라우터별 prefix 계산
package express

import "testing"

func TestResolveRouterPrefixes(t *testing.T) {
	allRouters := map[string]map[string]bool{
		"app.ts":   {"app": true},
		"users.ts": {"users": true},
	}
	mounts := []mountEntry{
		{
			prefix:       "/api",
			varName:      "users",
			filePath:     "users.ts",
			sourceFile:   "app.ts",
			sourceRouter: "app",
		},
	}
	prefixes := resolveRouterPrefixes(mounts, allRouters)

	app := routerKey{file: "app.ts", varName: "app"}
	users := routerKey{file: "users.ts", varName: "users"}
	if len(prefixes[app]) != 1 || prefixes[app][0] != "" {
		t.Fatalf("app prefix: %v", prefixes[app])
	}
	if len(prefixes[users]) != 1 || prefixes[users][0] != "/api" {
		t.Fatalf("users prefix: %v", prefixes[users])
	}
}
