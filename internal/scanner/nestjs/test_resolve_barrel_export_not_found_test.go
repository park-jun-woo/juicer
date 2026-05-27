//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveBarrelExport_NotFound barrel re-export 미발견 테스트
package nestjs

import "testing"

func TestResolveBarrelExport_NotFound(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dtos/other.dto.ts", `
export class OtherDto { name: string; }
`)
	writeFile(t, dir, "dtos/index.ts", `
export * from './other.dto';
`)
	got := resolveBarrelExport(dir+"/dtos/index.ts", "MissingDto")
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
