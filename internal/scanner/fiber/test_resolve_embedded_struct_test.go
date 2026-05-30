//ff:func feature=scan type=test control=sequence
//ff:what TestResolveEmbedded_Struct 테스트
package fiber

import "testing"

func TestResolveEmbedded_Struct(t *testing.T) {
	src := `package m
type Base struct {
	ID int ` + "`json:\"id\"`" + `
}
type User struct {
	Base
}
`
	st, _ := structFields(t, src, "User")

	embType := st.Field(0).Type()
	fields := resolveEmbedded(embType, map[string]bool{})
	if len(fields) != 1 || fields[0].Name != "ID" {
		t.Fatalf("expected ID field, got %v", fields)
	}
}
