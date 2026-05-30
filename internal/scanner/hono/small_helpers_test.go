//ff:func feature=scan type=test control=sequence topic=hono
//ff:what prefixKey / readObjectStringProp / resolveRelativePath / unquoteTS 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestPrefixKey(t *testing.T) {
	if got := prefixKey("file.ts", "app"); got != "file.ts\x00app" {
		t.Fatalf("got %q", got)
	}
}

func TestReadObjectStringProp_Found(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { method: "post", count: 3 };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	if got := readObjectStringProp(obj, "method", fi.Src); got != "post" {
		t.Fatalf("got %q", got)
	}
}

func TestReadObjectStringProp_NotString(t *testing.T) {
	// value present but not a string -> ""
	fi := mustParse(t, []byte(`const o = { count: 3 };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	if got := readObjectStringProp(obj, "count", fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestReadObjectStringProp_Missing(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { a: "x" };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	if got := readObjectStringProp(obj, "missing", fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestResolveRelativePath_Direct(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "mod.ts", "x\n")
	// import path includes extension -> direct stat hit
	if got := resolveRelativePath(dir, "./mod.ts"); got != filepath.Join(dir, "mod.ts") {
		t.Fatalf("got %q", got)
	}
}

func TestResolveRelativePath_WithExt(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "mod.ts", "x\n")
	if got := resolveRelativePath(dir, "./mod"); got != filepath.Join(dir, "mod.ts") {
		t.Fatalf("got %q", got)
	}
}

func TestResolveRelativePath_TsxExt(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "comp.tsx", "x\n")
	if got := resolveRelativePath(dir, "./comp"); got != filepath.Join(dir, "comp.tsx") {
		t.Fatalf("got %q", got)
	}
}

func TestResolveRelativePath_NotFound(t *testing.T) {
	dir := t.TempDir()
	if got := resolveRelativePath(dir, "./missing"); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestUnquoteTS(t *testing.T) {
	cases := []struct{ in, want string }{
		{`"hi"`, "hi"},
		{`'hi'`, "hi"},
		{"`hi`", "hi"},
		{"x", "x"},          // too short
		{"", ""},            // empty
		{`"mismatch'`, `"mismatch'`}, // mismatched quotes
		{"noquote", "noquote"},
	}
	for _, c := range cases {
		if got := unquoteTS(c.in); got != c.want {
			t.Errorf("unquoteTS(%q)=%q want %q", c.in, got, c.want)
		}
	}
}
