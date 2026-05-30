//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestMergeImportedAliases_NoAliasInSource 테스트
package fastapi

import (
	"path/filepath"
	"testing"
)

func TestMergeImportedAliases_NoAliasInSource(t *testing.T) {
	dir := t.TempDir()
	mkFile(t, dir, "deps.py", "x = 1")
	fi := fileInfo{
		absPath: filepath.Join(dir, "main.py"),
		imports: []importInfo{{name: "SessionDep", module: ".deps"}},
	}

	global := map[string]string{}
	mergeImportedAliases(fi, map[string]map[string]string{}, global)
	if len(global) != 0 {
		t.Fatalf("expected no merge, got %v", global)
	}
}
