//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefixInFile_NotFound 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestDetectGlobalPrefixInFile_NotFound(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "main.ts", `app.listen(3000);`)
	prefix, found := detectGlobalPrefixInFile(filepath.Join(dir, "main.ts"))
	if found || prefix != "" {
		t.Fatalf("got %q %v", prefix, found)
	}
}
