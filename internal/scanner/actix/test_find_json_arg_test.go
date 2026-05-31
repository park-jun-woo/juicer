//ff:func feature=scan type=test topic=actix control=iteration dimension=1
//ff:what findJSONArg .json(arg) 호출의 첫 인자 노드 반환 테스트
package actix

import "testing"

func TestFindJSONArg(t *testing.T) {
	src := []byte(`fn h() { HttpResponse::Ok().json(UserResponse { id: 1 }) }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	scoped := findAllByType(root, "scoped_identifier")
	found := false
	for _, sid := range scoped {
		a := findJSONArg(sid, src)
		if a != nil && a.Type() == "struct_expression" {
			found = true
		}
	}
	if !found {
		t.Error("json struct arg not found")
	}
}
