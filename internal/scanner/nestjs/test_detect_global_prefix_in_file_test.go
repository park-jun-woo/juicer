//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefixInFile 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestDetectGlobalPrefixInFile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "main.ts", `app.setGlobalPrefix('api');`)
	prefix, found := detectGlobalPrefixInFile(filepath.Join(dir, "main.ts"))
	if !found || prefix != "api" {
		t.Fatalf("got %q %v", prefix, found)
	}
}
