//ff:func feature=scan type=test topic=hono control=sequence
//ff:what resolveOneImport 상대 import의 기본/명명 바인딩을 파일 경로로 해석 테스트
package hono

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveOneImport(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "routes.ts")
	if err := os.WriteFile(target, []byte("export const r = 1;"), 0o644); err != nil {
		t.Fatal(err)
	}

	// default import
	fi := mustParse(t, []byte(`import router from './routes';`))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	imports := map[string]string{}
	resolveOneImport(stmt, fi.Src, dir, imports, dir)
	if imports["router"] != target {
		t.Errorf("default import: %v", imports)
	}

	// non-relative import -> ignored
	fi2 := mustParse(t, []byte(`import x from 'hono';`))
	stmt2 := findAllByType(fi2.Root, "import_statement")[0]
	imports2 := map[string]string{}
	resolveOneImport(stmt2, fi2.Src, dir, imports2, dir)
	if len(imports2) != 0 {
		t.Errorf("non-relative should be ignored: %v", imports2)
	}
}
