//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestSingleIntArg 테스트
package quarkus

import "testing"

func TestSingleIntArg(t *testing.T) {
	root, _ := parseJava([]byte(`class C { @Min(5) int x; }`))
	src := []byte(`class C { @Min(5) int x; }`)
	ann := findAllByType(root, "annotation")[0]
	v, ok := singleIntArg(ann, src)
	if !ok || v != 5 {
		t.Fatalf("got %d %v", v, ok)
	}
}
