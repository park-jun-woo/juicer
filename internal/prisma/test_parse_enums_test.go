//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what parseEnums 주석 제거 후 enum 블록 파싱 테스트
package prisma

import "testing"

func TestParseEnums(t *testing.T) {
	src := `// c
enum Role {
  ADMIN
  USER
}
model User { id Int @id }`
	enums := parseEnums(src)
	if len(enums) != 1 || enums[0].Name != "Role" {
		t.Fatalf("got %+v", enums)
	}
	if len(enums[0].Values) != 2 {
		t.Errorf("values: %v", enums[0].Values)
	}
}
