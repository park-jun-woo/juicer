//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what TestScanOneFilePass1_CollectsMounts 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestScanOneFilePass1_CollectsMounts(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "users.ts", `const r = express.Router();`)
	appPath := filepath.Join(dir, "app.ts")
	writeFile(t, dir, "app.ts", `import express from 'express';
import userRouter from './users';
const app = express();
const userSchema = z.object({ name: z.string() });
app.use('/api', userRouter);
`)
	parsed, allRouters, schemas, schemaSrc := newPass1Maps()
	entries := scanOneFilePass1(appPath, parsed, allRouters, dir, nil, schemas, schemaSrc)

	if _, ok := parsed[appPath]; !ok {
		t.Fatal("file not parsed")
	}
	if _, ok := schemas["userSchema"]; !ok {
		t.Fatalf("schema not collected: %v", schemas)
	}
	found := false
	for _, e := range entries {
		if e.prefix == "/api" && e.varName == "userRouter" {
			found = true
		}
	}
	if !found {
		t.Fatalf("mount entry not found: %+v", entries)
	}
}
