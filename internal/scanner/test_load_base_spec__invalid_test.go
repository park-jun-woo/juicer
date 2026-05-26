//ff:func feature=scan type=extract control=sequence
//ff:what TestLoadBaseSpec_Invalid 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadBaseSpec_Invalid(t *testing.T) {
	dir := t.TempDir()
	specPath := filepath.Join(dir, "openapi.yaml")
	// 빈 파일은 empty document 에러를 발생시킨다
	os.WriteFile(specPath, []byte(""), 0o644)

	_, err := LoadBaseSpec(specPath)
	if err == nil {
		t.Fatal("expected error for empty document")
	}
}
