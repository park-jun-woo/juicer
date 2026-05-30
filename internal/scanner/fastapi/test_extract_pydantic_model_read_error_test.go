//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractPydanticModel_ReadError 테스트
package fastapi

import (
	"path/filepath"
	"testing"
)

func TestExtractPydanticModel_ReadError(t *testing.T) {
	if _, err := extractPydanticModel(filepath.Join(t.TempDir(), "missing.py"), "User"); err == nil {
		t.Fatal("expected read error")
	}
}
