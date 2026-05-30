//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestExtractFields_EmbeddedAndExcluded 테스트
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
