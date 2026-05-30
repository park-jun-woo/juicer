//ff:func feature=scan type=test control=sequence
//ff:what TestDetectDotnet_AspNetCorePackage 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectDotnet_AspNetCorePackage(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "api.csproj"),
		[]byte(`<PackageReference Include="Microsoft.AspNetCore.App" />`), 0o644)
	if !detectDotnet(dir) {
		t.Fatal("expected true for AspNetCore package")
	}
}
