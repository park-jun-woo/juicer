//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what resolveEnumPathArg Enum.Member 경로를 값으로 해석(동일파일/비멤버/임포트) 테스트
package nestjs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveEnumPathArg(t *testing.T) {
	src := []byte(`enum RouteKey { Asset = 'assets' }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	// same-file resolution
	if v, ok := resolveEnumPathArg("RouteKey.Asset", root, src, "f.ts", nil, ""); !ok || v != "assets" {
		t.Errorf("same-file: (%q,%v)", v, ok)
	}
	// not a member expression
	if _, ok := resolveEnumPathArg("plain", root, src, "f.ts", nil, ""); ok {
		t.Error("plain string should be false")
	}
	// member with nested access -> false
	if _, ok := resolveEnumPathArg("A.B.C", root, src, "f.ts", nil, ""); ok {
		t.Error("nested member should be false")
	}

	// cross-file import resolution
	dir := t.TempDir()
	enumFile := filepath.Join(dir, "keys.ts")
	if err := os.WriteFile(enumFile, []byte(`export enum K { X = 'xx' }`), 0o644); err != nil {
		t.Fatal(err)
	}
	emptyRoot, _ := parseTypeScript([]byte(`const a = 1;`))
	imports := map[string]string{"K": "./keys"}
	if v, ok := resolveEnumPathArg("K.X", emptyRoot, []byte(`const a=1;`), filepath.Join(dir, "c.ts"), imports, dir); !ok || v != "xx" {
		t.Errorf("cross-file: (%q,%v)", v, ok)
	}
}
