//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what parseSource 주석 제거 후 모든 model 파싱 테스트
package prisma

import "testing"

func TestParseSource(t *testing.T) {
	src := `// a comment
model User {
  id Int @id
  email String
}`
	models := parseSource(src)
	if len(models) != 1 || models[0].name != "User" {
		t.Fatalf("got %+v", models)
	}
	if len(models[0].fields) != 2 {
		t.Errorf("fields: %v", models[0].fields)
	}
}
