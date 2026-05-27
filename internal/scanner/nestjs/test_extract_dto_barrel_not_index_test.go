//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractDTO_BarrelNotIndex 비 index.ts barrel 파일 DTO 추출 무시 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestExtractDTO_BarrelNotIndex(t *testing.T) {
	dir := t.TempDir()
	// file is not index.ts — barrel tracking should not trigger
	writeFile(t, dir, "dtos/barrel.ts", `
export * from './some.dto';
`)
	cache := make(map[string][]scanner.Field)
	fields, err := extractDTO(dir+"/dtos/barrel.ts", "SomeDto", nil, "", cache)
	if err != nil {
		t.Fatal(err)
	}
	if fields != nil {
		t.Fatalf("expected nil for non-index barrel file, got %v", fields)
	}
}
