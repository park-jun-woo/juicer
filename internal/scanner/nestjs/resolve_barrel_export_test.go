//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveBarrelExport_Found barrel re-export 추적 성공 테스트
package nestjs

import "testing"

func TestResolveBarrelExport_Found(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dtos/create-category.dto.ts", `
export class CreateCategoryDto {
  name: string;
}
`)
	writeFile(t, dir, "dtos/index.ts", `
export * from './create-category.dto';
export * from './update-category.dto';
`)
	got := resolveBarrelExport(dir+"/dtos/index.ts", "CreateCategoryDto")
	if got == "" {
		t.Fatal("expected resolved path, got empty")
	}
}
