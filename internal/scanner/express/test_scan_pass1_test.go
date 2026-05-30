//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestScanPass1 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestScanPass1(t *testing.T) {
	dir := t.TempDir()
	usersPath := filepath.Join(dir, "users.ts")
	appPath := filepath.Join(dir, "app.ts")
	writeFile(t, dir, "users.ts", `const users = express.Router();
users.get('/list', (req, res) => { res.json([]); });
`)
	writeFile(t, dir, "app.ts", `import express from 'express';
import users from './users';
const app = express();
app.use('/api', users);
`)
	ctx := scanPass1([]string{appPath, usersPath}, dir)
	if ctx == nil {
		t.Fatal("nil ctx")
	}
	if len(ctx.parsed) != 2 {
		t.Fatalf("expected 2 parsed, got %d", len(ctx.parsed))
	}
	usersKey := routerKey{file: usersPath, varName: "users"}
	if len(ctx.routerPrefixes[usersKey]) == 0 || ctx.routerPrefixes[usersKey][0] != "/api" {
		t.Fatalf("users prefix wrong: %v", ctx.routerPrefixes[usersKey])
	}
}
