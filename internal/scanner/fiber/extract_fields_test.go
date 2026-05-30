//ff:func feature=scan type=test control=iteration dimension=1
//ff:what extractFields — 구조체 필드 추출(임베딩/제외) 테스트
package fiber

import "testing"

func TestExtractFields_EmbeddedAndExcluded(t *testing.T) {
	src := `package m
type Base struct {
	ID int ` + "`json:\"id\"`" + `
}
type User struct {
	Base
	Name   string ` + "`json:\"name\"`" + `
	Secret string ` + "`json:\"-\"`" + `
}
`
	st, _ := structFields(t, src, "User")
	fields := extractFields(st, map[string]bool{})

	names := map[string]bool{}
	for _, f := range fields {
		names[f.Name] = true
	}
	if !names["ID"] {
		t.Errorf("embedded ID not expanded: %v", names)
	}
	if !names["Name"] {
		t.Errorf("Name missing: %v", names)
	}
	if names["Secret"] {
		t.Errorf("json:\"-\" Secret should be excluded: %v", names)
	}
}

func TestExtractFields_Simple(t *testing.T) {
	src := `package m
type T struct {
	A string
	B int
}
`
	st, _ := structFields(t, src, "T")
	fields := extractFields(st, map[string]bool{})
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
}
