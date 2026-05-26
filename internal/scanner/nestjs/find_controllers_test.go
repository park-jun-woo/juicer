//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindTSFiles_Basic 테스트
package nestjs

import "testing"

func TestFindTSFiles_Basic(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.ts", "console.log('hello');")
	writeFile(t, dir, "src/app.controller.ts", "class AppCtrl {}")
	writeFile(t, dir, "src/types.d.ts", "declare type X = string;")
	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 ts files (excluding .d.ts), got %d", len(files))
	}
}
