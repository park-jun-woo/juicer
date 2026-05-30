//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what mergeImportedAliases: cross-file 별칭 병합 / 미해결 import 스킵 / perFile 미존재
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
			{name: "External", module: "external_pkg"}, // unresolved -> skipped
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

func TestMergeImportedAliases_NoAliasInSource(t *testing.T) {
	dir := t.TempDir()
	mkFile(t, dir, "deps.py", "x = 1")
	fi := fileInfo{
		absPath: filepath.Join(dir, "main.py"),
		imports: []importInfo{{name: "SessionDep", module: ".deps"}},
	}
	// perFile has no alias entry for the source file
	global := map[string]string{}
	mergeImportedAliases(fi, map[string]map[string]string{}, global)
	if len(global) != 0 {
		t.Fatalf("expected no merge, got %v", global)
	}
}
