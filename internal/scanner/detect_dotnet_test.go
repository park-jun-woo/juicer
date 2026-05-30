//ff:func feature=scan type=test control=sequence
//ff:what detectDotnet — *.csproj의 AspNetCore 의존 감지 분기를 검증
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

func TestDetectDotnet_AspNetCorePackage(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "api.csproj"),
		[]byte(`<PackageReference Include="Microsoft.AspNetCore.App" />`), 0o644)
	if !detectDotnet(dir) {
		t.Fatal("expected true for AspNetCore package")
	}
}

func TestDetectDotnet_Miss(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "lib.csproj"),
		[]byte(`<Project Sdk="Microsoft.NET.Sdk"></Project>`), 0o644)
	// also create a skipped dir to exercise SkipDir branch
	os.MkdirAll(filepath.Join(dir, "obj"), 0o755)
	os.WriteFile(filepath.Join(dir, "obj", "x.csproj"), []byte("Microsoft.AspNetCore"), 0o644)
	if detectDotnet(dir) {
		t.Fatal("expected false for non-web csproj (obj dir skipped)")
	}
}

func TestDetectDotnet_NoCsproj(t *testing.T) {
	if detectDotnet(t.TempDir()) {
		t.Fatal("expected false when no csproj")
	}
}
