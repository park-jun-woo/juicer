//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveBarrelExport_FileError 존재하지 않는 파일 에러 확인 테스트
package nestjs

import "testing"

func TestResolveBarrelExport_FileError(t *testing.T) {
	got := resolveBarrelExport("/nonexistent/index.ts", "Dto")
	if got != "" {
		t.Fatalf("expected empty for missing file, got %q", got)
	}
}
