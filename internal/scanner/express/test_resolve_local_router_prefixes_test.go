//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 단위 테스트: resolveLocalRouterPrefixes가 인라인 라우터 prefix를 올바르게 합산하는지 검증한다
package express

import "testing"

func TestResolveLocalRouterPrefixes(t *testing.T) {
	tests := []struct {
		name     string
		entries  []mountEntry
		wantPath map[string]string // filePath -> expected prefix
	}{
		{
			name: "single inline router",
			entries: []mountEntry{
				{prefix: "/v1", varName: "v1Router", filePath: "", sourceRouter: "app"},
				{prefix: "/orgs", varName: "orgRouter", filePath: "routes/orgs.ts", sourceRouter: "v1Router"},
			},
			wantPath: map[string]string{
				"routes/orgs.ts": "/v1/orgs",
			},
		},
		{
			name: "multi-level inline",
			entries: []mountEntry{
				{prefix: "/api", varName: "apiRouter", filePath: "", sourceRouter: "app"},
				{prefix: "/v1", varName: "v1Router", filePath: "", sourceRouter: "apiRouter"},
				{prefix: "/users", varName: "usersRouter", filePath: "routes/users.ts", sourceRouter: "v1Router"},
			},
			wantPath: map[string]string{
				"routes/users.ts": "/api/v1/users",
			},
		},
		{
			name: "no inline router",
			entries: []mountEntry{
				{prefix: "/orgs", varName: "orgRouter", filePath: "routes/orgs.ts", sourceRouter: "app"},
			},
			wantPath: map[string]string{
				"routes/orgs.ts": "/orgs",
			},
		},
		{
			name: "mixed inline and direct",
			entries: []mountEntry{
				{prefix: "/v1", varName: "v1Router", filePath: "", sourceRouter: "app"},
				{prefix: "/orgs", varName: "orgRouter", filePath: "routes/orgs.ts", sourceRouter: "v1Router"},
				{prefix: "/health", varName: "healthRouter", filePath: "routes/health.ts", sourceRouter: "app"},
			},
			wantPath: map[string]string{
				"routes/orgs.ts":   "/v1/orgs",
				"routes/health.ts": "/health",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resolveLocalRouterPrefixes(tt.entries)
			assertPrefixMap(t, result, tt.wantPath)
		})
	}
}
