//ff:func feature=scan type=test control=sequence topic=express
//ff:what scanOneFilePass1: 파싱·라우터·스키마·마운트 수집 / 파싱실패 nil
package express

import (
	"path/filepath"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func newPass1Maps() (map[string]*fileInfo, map[string]map[string]bool, map[string]*sitter.Node, map[string][]byte) {
	return map[string]*fileInfo{}, map[string]map[string]bool{}, map[string]*sitter.Node{}, map[string][]byte{}
}

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

func TestScanOneFilePass1_ParseError(t *testing.T) {
	parsed, allRouters, schemas, schemaSrc := newPass1Maps()
	entries := scanOneFilePass1("/no/such/file.ts", parsed, allRouters, "/no/such", nil, schemas, schemaSrc)
	if entries != nil {
		t.Fatalf("expected nil on parse error, got %+v", entries)
	}
}
