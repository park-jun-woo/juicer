//ff:func feature=scan type=test control=sequence
//ff:what TestDetectDotnet_Miss 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectDotnet_Miss(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "lib.csproj"),
		[]byte(`<Project Sdk="Microsoft.NET.Sdk"></Project>`), 0o644)

	os.MkdirAll(filepath.Join(dir, "obj"), 0o755)
	os.WriteFile(filepath.Join(dir, "obj", "x.csproj"), []byte("Microsoft.AspNetCore"), 0o644)
	if detectDotnet(dir) {
		t.Fatal("expected false for non-web csproj (obj dir skipped)")
	}
}
