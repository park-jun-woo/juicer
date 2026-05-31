//ff:func feature=scan type=test topic=flask control=iteration dimension=1
//ff:what isJSONSource request.json 속성 / request.get_json() 호출 판별 테스트
package flask

import "testing"

func TestIsJSONSource(t *testing.T) {
	// attribute request.json
	src := []byte("x = request.json\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	attrs := findAllByType(root, "attribute")
	found := false
	for _, a := range attrs {
		if isJSONSource(a, src) {
			found = true
		}
	}
	if !found {
		t.Error("request.json attribute should be json source")
	}

	// call request.get_json()
	src2 := []byte("x = request.get_json()\n")
	root2, _ := parsePython(src2)
	calls := findAllByType(root2, "call")
	if len(calls) == 0 || !isJSONSource(calls[0], src2) {
		t.Error("request.get_json() should be json source")
	}

	// nil -> false
	if isJSONSource(nil, src) {
		t.Error("nil should be false")
	}
}
