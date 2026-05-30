//ff:func feature=scan type=test control=sequence
//ff:what TestDetectDotnet_Hit 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectDotnet_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "app.csproj"),
		[]byte(`<Project Sdk="Microsoft.NET.Sdk.Web"></Project>`), 0o644)
	if !detectDotnet(dir) {
		t.Fatal("expected true")
	}
}
