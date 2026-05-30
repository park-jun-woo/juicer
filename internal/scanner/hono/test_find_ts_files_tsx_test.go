//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what .tsx 파일은 수집하고 .d.ts 파일은 제외하는지 테스트
package hono

import "testing"

func TestFindTSFiles_TSX(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "a.ts", "export const a = 1")
	writeFile(t, dir, "b.tsx", "export const b = 2")
	writeFile(t, dir, "types.d.ts", "export type T = number")

	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatalf("findTSFiles error: %v", err)
	}
	found := map[string]bool{}
	for _, f := range files {
		found[f[len(dir)+1:]] = true
	}
	if !found["a.ts"] {
		t.Errorf("expected a.ts collected, got %v", found)
	}
	if !found["b.tsx"] {
		t.Errorf("expected b.tsx collected, got %v", found)
	}
	if found["types.d.ts"] {
		t.Errorf("expected types.d.ts excluded, got %v", found)
	}
}

func TestFindTSFiles_ExcludedDirs(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", "x")
	writeFile(t, dir, "node_modules/lib.ts", "x")
	writeFile(t, dir, "dist/out.ts", "x")
	writeFile(t, dir, "build/out.ts", "x")
	writeFile(t, dir, ".git/hook.ts", "x")
	writeFile(t, dir, "__tests__/case.ts", "x")

	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		rel := f[len(dir)+1:]
		if rel != "app.ts" {
			t.Errorf("unexpected collected file: %s", rel)
		}
	}
	if len(files) != 1 {
		t.Fatalf("expected only app.ts, got %v", files)
	}
}

func TestFindTSFiles_TestFilesExcluded(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", "x")
	writeFile(t, dir, "app.test.ts", "x")
	writeFile(t, dir, "app.spec.ts", "x")

	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		rel := f[len(dir)+1:]
		if rel != "app.ts" {
			t.Errorf("test file leaked: %s", rel)
		}
	}
}
