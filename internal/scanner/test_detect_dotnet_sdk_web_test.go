//ff:func feature=scan type=test control=sequence
//ff:what TestDetectDotnet_SdkWeb -- Microsoft.NET.Sdk.Web으로 감지
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectDotnet_SdkWeb(t *testing.T) {
	dir := t.TempDir()
	csproj := `<Project Sdk="Microsoft.NET.Sdk.Web">
  <PropertyGroup>
    <TargetFramework>net8.0</TargetFramework>
  </PropertyGroup>
</Project>`
	if err := os.WriteFile(filepath.Join(dir, "WebApi.csproj"), []byte(csproj), 0o644); err != nil {
		t.Fatal(err)
	}
	if !detectDotnet(dir) {
		t.Error("expected dotnet detected via SDK.Web")
	}
}
