//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestFindTSFiles_TSX 테스트
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
