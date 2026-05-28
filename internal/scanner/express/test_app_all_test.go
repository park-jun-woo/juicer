//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what app.all("/health", handler) → 5개 HTTP 메서드 엔드포인트 테스트
package express

import "testing"

func TestAppAll(t *testing.T) {
	dir := t.TempDir()
	src := `
const express = require("express");
const app = express();
app.all("/health", healthCheck);
`
	writeFile(t, dir, "app.ts", src)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 5 {
		t.Fatalf("expected 5 endpoints for app.all, got %d", len(result.Endpoints))
	}
	methods := map[string]bool{}
	for _, ep := range result.Endpoints {
		methods[ep.Method] = true
		if ep.Path != "/health" {
			t.Errorf("path: want /health, got %s", ep.Path)
		}
	}
	for _, m := range []string{"GET", "POST", "PUT", "PATCH", "DELETE"} {
		if !methods[m] {
			t.Errorf("missing method %s", m)
		}
	}
}
