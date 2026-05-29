//ff:func feature=scan type=test control=sequence
//ff:what TestDetectDotnet -- ASP.NET Core 프로젝트 감지 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectDotnet(t *testing.T) {
	dir := t.TempDir()
	csproj := `<Project Sdk="Microsoft.NET.Sdk.Web">
  <PropertyGroup>
    <TargetFramework>net8.0</TargetFramework>
  </PropertyGroup>
  <ItemGroup>
    <PackageReference Include="Microsoft.AspNetCore.OpenApi" Version="8.0.0" />
  </ItemGroup>
</Project>`
	if err := os.WriteFile(filepath.Join(dir, "MyApp.csproj"), []byte(csproj), 0o644); err != nil {
		t.Fatal(err)
	}
	if !detectDotnet(dir) {
		t.Error("expected dotnet detected")
	}
}
