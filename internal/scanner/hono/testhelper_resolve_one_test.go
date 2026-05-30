//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what resolveOne 테스트 헬퍼
package hono

import (
	"path/filepath"
	"testing"
)

func resolveOne(t *testing.T, dir, appSrc string) map[string]string {
	t.Helper()
	writeFile(t, dir, "app.ts", appSrc+"\n")
	fi, err := parseFile(filepath.Join(dir, "app.ts"))
	if err != nil {
		t.Fatal(err)
	}
	imports := map[string]string{}
	for _, stmt := range findAllByType(fi.Root, "import_statement") {
		resolveOneImport(stmt, fi.Src, dir, imports, dir)
	}
	return imports
}
