//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindTSFiles_RootFallback 테스트 (root/src 없으면 root 자체 순회)
package nestjs

import "testing"

func TestFindTSFiles_RootFallback(t *testing.T) {
	dir := t.TempDir()
	// No src/ subdirectory: controllers live directly under root (e.g. the
	// caller already passed a .../src path).
	writeFile(t, dir, "app.controller.ts", "class AppCtrl {}")
	writeFile(t, dir, "users/users.controller.ts", "class UsersCtrl {}")
	writeFile(t, dir, "types.d.ts", "declare type X = string;")
	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 ts files via root fallback (excluding .d.ts), got %d", len(files))
	}
}
