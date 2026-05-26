//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractDTO_FileError 테스트
package nestjs

import "testing"

func TestExtractDTO_FileError(t *testing.T) {
	_, err := extractDTO("/nonexistent/file.ts", "Dto", nil, "", nil)
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
