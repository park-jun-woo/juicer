//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what findTSFiles가 .js/.jsx/.mjs/.cjs 소스를 수집하고 .d.ts/테스트 파일을 제외하는지 검증
package express

import (
	"path/filepath"
	"testing"
)

func TestFindTSFiles_CollectsJS(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "a.js", "const x = 1;")
	writeFile(t, dir, "b.jsx", "const y = 2;")
	writeFile(t, dir, "c.mjs", "export const z = 3;")
	writeFile(t, dir, "d.cjs", "module.exports = {};")
	writeFile(t, dir, "e.ts", "const e: number = 5;")
	writeFile(t, dir, "f.tsx", "const f = 6;")
	// excluded
	writeFile(t, dir, "types.d.ts", "declare const g: number;")
	writeFile(t, dir, "a.test.js", "test('x', () => {});")
	writeFile(t, dir, "b.spec.js", "describe('y', () => {});")
	writeFile(t, dir, "node_modules/pkg/index.js", "module.exports = 0;")

	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatalf("findTSFiles error: %v", err)
	}

	got := map[string]bool{}
	for _, f := range files {
		got[filepath.Base(f)] = true
	}

	for _, want := range []string{"a.js", "b.jsx", "c.mjs", "d.cjs", "e.ts", "f.tsx"} {
		if !got[want] {
			t.Errorf("expected %s collected, got %v", want, got)
		}
	}
	for _, notWant := range []string{"types.d.ts", "a.test.js", "b.spec.js", "index.js"} {
		if got[notWant] {
			t.Errorf("expected %s excluded, got %v", notWant, got)
		}
	}
}
