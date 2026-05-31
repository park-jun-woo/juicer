//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what appendIfNotEmpty 비빈 문자열만 추가 테스트
package prisma

import "testing"

func TestAppendIfNotEmpty(t *testing.T) {
	dst := []string{"a"}
	dst = appendIfNotEmpty(dst, "")
	if len(dst) != 1 {
		t.Fatalf("empty must not append: %v", dst)
	}
	dst = appendIfNotEmpty(dst, "b")
	if len(dst) != 2 || dst[1] != "b" {
		t.Fatalf("non-empty must append: %v", dst)
	}
}
