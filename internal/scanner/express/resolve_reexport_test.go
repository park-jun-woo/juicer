//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveReexport가 배럴 re-export specifier 바인딩명을 source 실파일로 해석하는지 검증한다
package express

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveReexportFollowsBarrel(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "user.validation.ts")
	if err := os.WriteFile(target, []byte("const createUser = { body: 1 };\nexport default { createUser };\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	barrel := filepath.Join(dir, "index.ts")
	src := []byte("export { default as userValidation } from './user.validation';\nexport { foo } from './bar';\n")
	if err := os.WriteFile(barrel, src, 0o644); err != nil {
		t.Fatal(err)
	}
	fi, err := parseFile(barrel)
	if err != nil {
		t.Fatal(err)
	}

	got := resolveReexport(fi, "userValidation", dir, nil)
	if got != target {
		t.Fatalf("userValidation: got %q, want %q", got, target)
	}

	if got := resolveReexport(fi, "missingName", dir, nil); got != "" {
		t.Fatalf("missingName: got %q, want empty", got)
	}
}
