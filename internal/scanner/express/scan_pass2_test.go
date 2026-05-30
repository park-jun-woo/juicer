//ff:func feature=scan type=test control=sequence topic=express
//ff:what scanPass2: 라우트→Endpoint 생성 / 라우터없는 파일 스킵
package express

import (
	"path/filepath"
	"testing"
)

func TestScanPass2_BuildsEndpoints(t *testing.T) {
	dir := t.TempDir()
	appPath := filepath.Join(dir, "app.ts")
	writeFile(t, dir, "app.ts", `import express from 'express';
const app = express();
app.get('/users', (req, res) => { res.json([]); });
`)
	// also a file with no routers -> exercises the continue branch
	noRouterPath := filepath.Join(dir, "util.ts")
	writeFile(t, dir, "util.ts", `export const x = 1;`)

	ctx := scanPass1([]string{appPath, noRouterPath}, dir)
	eps := scanPass2(ctx, dir)
	if len(eps) == 0 {
		t.Fatalf("expected endpoints, got %+v", eps)
	}
	found := false
	for _, e := range eps {
		if e.Path == "/users" && e.Method == "GET" {
			found = true
		}
	}
	if !found {
		t.Fatalf("GET /users not found: %+v", eps)
	}
}
