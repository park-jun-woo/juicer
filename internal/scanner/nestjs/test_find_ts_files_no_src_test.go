//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindTSFiles_NoSrc 테스트
package nestjs

import "testing"

func TestFindTSFiles_NoSrc(t *testing.T) {
	files, err := findTSFiles(t.TempDir())
	if err != nil || files != nil {
		t.Fatalf("expected nil, got %v %v", files, err)
	}
}
