//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestMergeImportedAliases_Merges 테스트
package fastapi

import (
	"path/filepath"
	"testing"
)

func TestMergeImportedAliases_Merges(t *testing.T) {
	dir := t.TempDir()
	depsPath := mkFile(t, dir, "deps.py", "x = 1")
	appPath := filepath.Join(dir, "main.py")

	fi := fileInfo{
		absPath: appPath,
		imports: []importInfo{
			{name: "SessionDep", module: ".deps"},
			{name: "External", module: "external_pkg"},
		},
	}
	perFile := map[string]map[string]string{
		depsPath: {"SessionDep": "get_session"},
	}
	global := map[string]string{}
	mergeImportedAliases(fi, perFile, global)

	if global["SessionDep"] != "get_session" {
		t.Fatalf("expected SessionDep merged, got %v", global)
	}
	if _, ok := global["External"]; ok {
		t.Fatalf("unresolved import should be skipped: %v", global)
	}
}
